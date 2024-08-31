package handler

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"gihtub.com/heyyakash/switchr/internal/cache"
	"gihtub.com/heyyakash/switchr/internal/constants"
	"gihtub.com/heyyakash/switchr/internal/db"
	"gihtub.com/heyyakash/switchr/internal/modals"
	"gihtub.com/heyyakash/switchr/internal/utils"
	"github.com/gin-gonic/gin"
)

type CreateProjectRequest struct {
	Name string `json:"name"`
}

type ShareProjectStruct struct {
	Role  int    `json:"role"`
	Email string `json:"email"`
	Pid   string `json:"pid"`
}

func CreateProject() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req CreateProjectRequest
		uid := ctx.MustGet("uid").(string)
		if err := ctx.BindJSON(&req); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ResponseGenerator("Invalid fields", false))
			return
		}
		project := &modals.Projects{
			Name:      req.Name,
			CreatedBy: uid,
		}
		err := db.Store.CreateProject(project)
		if err != nil {
			log.Print(err)
			ctx.AbortWithStatusJSON(500, utils.ResponseGenerator("Some Error Occurred", false))
			return
		}
		userprojectmap := &modals.UserProjectMap{
			Uid:  uid,
			Pid:  project.Pid,
			Role: constants.Role["owner"],
		}
		if err := db.Store.CreateUserProjectMap(userprojectmap); err != nil {
			log.Print(err)
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, utils.ResponseGenerator("Some Error Occurred", false))
			return
		}

		newProjectList, err := db.Store.GetAllProjectsByUid(uid)
		if err == nil {
			cache.Redisdb.Set(fmt.Sprintf("USER-%s-PROJECTS", uid), newProjectList)
			cache.Redisdb.Set(fmt.Sprintf("PROJECT-%s", project.Pid), req)
		} else {
			log.Print(err)
		}

		ctx.JSON(http.StatusOK, utils.ResponseGenerator("Project Created", true))

	}
}

func GetProject() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uid := ctx.MustGet("uid").(string)
		pid := string(ctx.Param("pid"))
		_, err := db.Store.GetUserProjectMapByUidAndPid(uid, pid)
		if err != nil {
			log.Print(err)
			ctx.AbortWithStatusJSON(http.StatusNotFound, utils.ResponseGenerator("Not Found", false))
			return
		}
		val, err := cache.Redisdb.Get(fmt.Sprintf("PROJECT-%s", pid))
		if err == nil {
			ctx.JSON(http.StatusOK, utils.ResponseGenerator(val, true))
		}
		project, err := db.Store.GetProjectByPid(pid)
		if err != nil {
			log.Print(err)
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, utils.ResponseGenerator("Some Error Occured", false))
			return
		}
		ctx.JSON(http.StatusOK, utils.ResponseGenerator(project, true))

	}
}

func DeleteProject() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pid := string(ctx.Param("pid"))
		uid := ctx.MustGet("uid").(string)
		userprojetmap, err := db.Store.GetUserProjectMapByUidAndPid(uid, pid)
		if err != nil {
			log.Print(err)
			ctx.AbortWithStatusJSON(http.StatusNotFound, utils.ResponseGenerator("Not found", false))
			return
		}
		if userprojetmap.Role == constants.Role["reader"] || userprojetmap.Role == constants.Role["editor"] {
			log.Print(err)
			ctx.AbortWithStatusJSON(http.StatusForbidden, utils.ResponseGenerator("You don't have authority to delete the project", false))
			return
		}
		if err := db.Store.DeleteProjectByPid(pid); err != nil {
			log.Print(err)
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, utils.ResponseGenerator("Some Error occurred", false))
			return
		}

		ctx.JSON(http.StatusOK, utils.ResponseGenerator("Successfully deleted", true))
	}
}

func ShareProject() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uid := ctx.MustGet("uid").(string)
		var req ShareProjectStruct
		if err := ctx.BindJSON(&req); err != nil {
			log.Print(err)
			ctx.AbortWithStatusJSON(http.StatusBadGateway, utils.ResponseGenerator("Bad Request", false))
			return
		}
		userProjectMap, err := db.Store.GetUserProjectMapByUidAndPid(uid, req.Pid)
		if err != nil {
			log.Print(err)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ResponseGenerator("Bad Request", false))
			return
		}
		if userProjectMap.Role == constants.Role["reader"] {
			log.Print(err)
			ctx.AbortWithStatusJSON(http.StatusForbidden, utils.ResponseGenerator("Forbidden", false))
			return
		}
		res, err := db.Store.GetUserByEmail(req.Email)
		if err != nil {
			log.Print(err)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ResponseGenerator("Since the requested user doesn't have a switchr account we could'nt send them the invitation", false))
			return
		}
		_, err = db.Store.GetUserProjectMapByUidAndPid(res.Uid, req.Pid)
		if err == nil {
			log.Print(err)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ResponseGenerator("Requested user is already a part of the project", false))
			return
		}
		jwt, err := utils.GenerateJWTWithTypeUidAndPid(res.Uid, req.Pid, req.Role, "share", time.Now().Add(5*time.Minute).Unix())
		if err != nil {
			log.Print(err)
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, utils.ResponseGenerator("Some Error Occured", false))
			return
		}
		email := modals.Email{
			To:      req.Email,
			Subject: "Invitation to join Project on Switchr",
			Content: fmt.Sprintf("Hello %s\nYou have been invited to join %s by %s\nYour link is %s/share/confirm/%s", res.FullName, userProjectMap.Project.Name, userProjectMap.User.FullName, utils.GetString("HOST"), jwt),
		}

		if err := utils.SendEmail(&email); err != nil {
			log.Print(err)
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, utils.ResponseGenerator("Some Error Occured", false))
			return
		}

		ctx.JSON(http.StatusOK, utils.ResponseGenerator("Invitation Sent successfully", true))

	}
}

func ConfirmShareProject() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := string(ctx.Param("token"))
		val, valid, err := utils.DecodeJWT(token)
		if err != nil || !valid {
			ctx.JSON(http.StatusBadRequest, utils.ResponseGenerator("Invitation has expired", false))
			return
		}
		user, err := db.Store.GetUserByUid(val.Uid)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, utils.ResponseGenerator("User does not exist", false))
			return
		}
		project, err := db.Store.GetProjectByPid(val.Pid)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, utils.ResponseGenerator("Project does not exist", false))
			return
		}
		userProjectMap := modals.UserProjectMap{
			Pid:  project.Pid,
			Uid:  user.Uid,
			Role: val.Role,
		}
		err = db.Store.CreateUserProjectMap(&userProjectMap)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, utils.ResponseGenerator("Some error occured", false))
			return
		}
		ctx.JSON(http.StatusOK, utils.ResponseGenerator("User Added", true))
	}
}

func UpdateProject() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req modals.Projects
		uid := ctx.MustGet("uid").(string)
		if err := ctx.BindJSON(&req); err != nil {
			log.Print(err)
			ctx.JSON(http.StatusBadRequest, utils.ResponseGenerator("Bad Request", false))
			return
		}
		userprojectmap, err := db.Store.GetUserProjectMapByUidAndPid(uid, req.Pid)
		if err != nil {
			log.Print(err)
			ctx.AbortWithStatusJSON(http.StatusNotFound, utils.ResponseGenerator("Not found", false))
			return
		}
		if userprojectmap.Role != constants.Role["owner"] {
			ctx.AbortWithStatusJSON(http.StatusForbidden, utils.ResponseGenerator("You don't have the authority to update the project", false))
			return
		}
		if err := db.Store.UpdateProjectWithPid(&req, userprojectmap.Pid); err != nil {
			log.Print(err)
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, utils.ResponseGenerator("Some Error Occurred", false))
			return
		}
		ctx.JSON(http.StatusOK, utils.ResponseGenerator("Details updated successfully", true))
	}
}
