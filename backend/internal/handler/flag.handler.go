package handler

import (
	"fmt"
	"log"
	"net/http"

	"gihtub.com/heyyakash/switchr/internal/cache"
	"gihtub.com/heyyakash/switchr/internal/constants"
	"gihtub.com/heyyakash/switchr/internal/db"
	"gihtub.com/heyyakash/switchr/internal/modals"
	"gihtub.com/heyyakash/switchr/internal/utils"
	"github.com/gin-gonic/gin"
)

type CreateFlagRequest struct {
	Flag  string `json:"flag"`
	Value string `json:"value"`
	Pid   string `json:"pid"`
}

type UpdateFlagRequest struct {
	Value string `json:"value"`
}

func CreateFlag() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req CreateFlagRequest
		if err := ctx.BindJSON(&req); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ResponseGenerator("Bad data provided", false))
			return
		}
		uid := ctx.MustGet("uid").(string)
		userprojectmap, err := db.Store.GetUserProjectMapByUidAndPid(uid, req.Pid)
		if err != nil {
			log.Print(err)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ResponseGenerator("Bad request", false))
			return
		}
		if userprojectmap.Role == constants.Role["reader"] {
			ctx.AbortWithStatusJSON(http.StatusForbidden, utils.ResponseGenerator("Not permitted", false))
			return
		}
		_, err = db.Store.GetFlagByNameAndPid(req.Flag, req.Pid)
		if err == nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ResponseGenerator("Flag with the same name already exists", false))
			return
		}

		flag := modals.Featureflag{
			CreatedBy: uid,
			Flag:      req.Flag,
			Value:     req.Value,
			Pid:       req.Pid,
			UpdatedBy: uid,
		}
		if err := db.Store.CreateFlag(&flag); err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, utils.ResponseGenerator("Some Error Occurred", false))
			return
		}

		// creating new cache
		if err := cache.Redisdb.Set(fmt.Sprintf("PID-%s-FLAG-%s", req.Pid, flag.Flag), flag.Value); err != nil {
			log.Print(err)
		}

		ctx.JSON(http.StatusOK, utils.ResponseGenerator("Flag Created Successfully", true))

	}
}

func GetFlagByFid() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fid := string(ctx.Param("fid"))
		// val, err := cache.Redisdb.Get(fmt.Sprintf("FLAG-%s", fid))
		// if err == nil {
		// 	ctx.JSON(http.StatusOK, utils.ResponseGenerator(val, true))
		// 	return
		// }
		flagval, err := db.Store.GetFlagByFid(fid)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ResponseGenerator("Bad request", false))
			return
		}
		ctx.JSON(http.StatusOK, utils.ResponseGenerator(flagval, true))
	}
}

func DeleteFlag() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uid := ctx.MustGet("uid").(string)
		fid := string(ctx.Param("fid"))
		flag, err := db.Store.GetFlagByFid(fid)
		if err != nil {
			log.Print(err)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ResponseGenerator("Bad request", false))
			return
		}
		userprojectmap, err := db.Store.GetUserProjectMapByUidAndPid(uid, flag.Pid)
		if err != nil {
			log.Print(err)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ResponseGenerator("Bad request", false))
			return
		}
		if userprojectmap.Role == constants.Role["owner"] || userprojectmap.Role == constants.Role["editor"] {
			if err := db.Store.DeleteFlagByFid(fid); err != nil {
				log.Print(err)
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, utils.ResponseGenerator("Some Error Occurred", false))
				return
			}
			cache.Redisdb.Del(fmt.Sprintf("PID-%s-FLAG-%s", userprojectmap.Pid, flag.Flag))
			ctx.JSON(http.StatusOK, utils.ResponseGenerator("Flag Deleted Successfully", true))
			return
		}

		ctx.AbortWithStatusJSON(http.StatusForbidden, utils.ResponseGenerator("Forbidden", false))

	}
}

func GetFlagByPid() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pid := string(ctx.Param("pid"))
		res, err := db.Store.GetFlagByPid(pid)
		if err != nil {
			log.Print(err)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ResponseGenerator("Bad request", false))
			return
		}
		ctx.JSON(http.StatusOK, utils.ResponseGenerator(res, true))
	}
}

func UpdateFlag() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req UpdateFlagRequest
		if err := ctx.BindJSON(&req); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ResponseGenerator("Bad data provided", false))
			return
		}
		uid := ctx.MustGet("uid").(string)
		fid := string(ctx.Param("fid"))
		flag, err := db.Store.GetFlagByFid(fid)
		if err != nil {
			log.Print(err)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ResponseGenerator("Bad request", false))
			return
		}
		userprojectmap, err := db.Store.GetUserProjectMapByUidAndPid(uid, flag.Pid)
		if err != nil {
			log.Print(err)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ResponseGenerator("Bad request", false))
			return
		}

		if userprojectmap.Role == constants.Role["owner"] || userprojectmap.Role == constants.Role["editor"] {
			flag.Value = req.Value
			if err := db.Store.UpdateFlag(&flag); err != nil {
				log.Print(err)
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, utils.ResponseGenerator("Some Error Occurred", false))
				return
			}
			cache.Redisdb.Set(fmt.Sprintf("PID-%s-FLAG-%s", userprojectmap.Pid, flag.Flag), flag.Value)
			ctx.JSON(http.StatusOK, utils.ResponseGenerator("Flag Upated Successfully", true))
			return
		}

		ctx.AbortWithStatusJSON(http.StatusForbidden, utils.ResponseGenerator("Forbidden", false))
	}
}
