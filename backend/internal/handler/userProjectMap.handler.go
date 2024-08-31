package handler

import (
	"log"
	"net/http"

	"gihtub.com/heyyakash/switchr/internal/constants"
	"gihtub.com/heyyakash/switchr/internal/db"
	"gihtub.com/heyyakash/switchr/internal/utils"
	"github.com/gin-gonic/gin"
)

type DeleteMemberReq struct {
	MemId string `json:"memid"`
}

func GetUserProjects() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uid := ctx.MustGet("uid").(string)
		res, err := db.Store.GetUserProjectMapByUid(uid)
		if err != nil {
			log.Print(res)
			ctx.AbortWithStatusJSON(http.StatusNotFound, utils.ResponseGenerator("No Projects Found", false))
			return
		}
		ctx.JSON(http.StatusOK, utils.ResponseGenerator(res, true))
	}
}

func ProjectMembers() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pid := string(ctx.Param("pid"))
		uid := ctx.MustGet("uid").(string)
		_, err := db.Store.GetUserProjectMapByUidAndPid(uid, pid)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusForbidden, utils.ResponseGenerator("Forbidden", false))
			return
		}
		res, err := db.Store.GetMembersByPid(pid)
		if err != nil {
			log.Print(res)
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, utils.ResponseGenerator("Some error occuered", false))
			return
		}
		ctx.JSON(http.StatusOK, utils.ResponseGenerator(res, true))
	}
}

func DeleteMembers() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pid := string(ctx.Param("pid"))
		uid := ctx.MustGet("uid").(string)
		var req DeleteMemberReq
		if err := ctx.BindJSON(&req); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ResponseGenerator("Bad JSON", false))
			return
		}
		userprojectmap, err := db.Store.GetUserProjectMapByUidAndPid(uid, pid)
		if err != nil || userprojectmap.Role == constants.Role["reader"] {
			ctx.AbortWithStatusJSON(http.StatusForbidden, utils.ResponseGenerator("Forbidden", false))
			return
		}
		if userprojectmap.Role == constants.Role["reader"] {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.ResponseGenerator("Readers don't have authority to remove members", false))
			return
		}

		memberProjectMap, err := db.Store.GetUserProjectMapByUidAndPid(req.MemId, pid)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ResponseGenerator("Member does not exists", false))
			return
		}

		if userprojectmap.Role == constants.Role["editor"] && memberProjectMap.Role == constants.Role["owner"] {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.ResponseGenerator("Editor cannot remove owner", false))
			return
		}

		if memberProjectMap.Role == constants.Role["owner"] {
			allOwners, err := db.Store.FetchAllOwnersOfAProject(pid)
			if err != nil {
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, utils.ResponseGenerator("Some Error occuered", false))
				return
			}
			if len(allOwners) == 1 && allOwners[0].Uid == memberProjectMap.Uid {
				ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ResponseGenerator("Cannot remove the only remaining owner of the project", false))
				return
			}
		}

		if err := db.Store.DeleteUserProjectMapByUidPid(req.MemId, pid); err != nil {
			log.Print(err)
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, utils.ResponseGenerator("Some error occured", false))
			return
		}

		ctx.JSON(http.StatusOK, utils.ResponseGenerator("Member successfully removed from the project", true))

	}
}
