package handler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

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

func CreateProject() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req CreateProjectRequest
		uid := ctx.MustGet("uid").(string)
		if err := ctx.BindJSON(&req); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "Bad json",
				"success": false,
			})
			return
		}
		project := &modals.Projects{
			Name:      req.Name,
			CreatedBy: uid,
		}
		err := db.Store.CreateProject(project)
		if err != nil {
			log.Print(err)
			ctx.AbortWithStatusJSON(500, utils.ResponseGenerator("Some Error Occuered", false))
			return
		}
		userprojectmap := &modals.UserProjectMap{
			Uid:  uid,
			Pid:  project.Pid,
			Role: constants.Owner,
		}
		if err := db.Store.CreateUserProjectMap(userprojectmap); err != nil {
			log.Print(err)
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, utils.ResponseGenerator("Some Error Occured", false))
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
		intid, err := strconv.Atoi(ctx.Param("id"))
		id := uint(intid)
		if err != nil {
			log.Print(err)
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, utils.ResponseGenerator("Some Error Occurred", false))
			return
		}
		if err := db.Store.DeleteProjectById(id); err != nil {
			log.Print(err)
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, utils.ResponseGenerator("Some Error occurred", false))
			return
		}

		ctx.JSON(http.StatusOK, utils.ResponseGenerator("Successfully deleted", true))
	}
}
