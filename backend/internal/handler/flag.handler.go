package handler

import (
	"fmt"
	"log"
	"net/http"

	"gihtub.com/heyyakash/switchr/internal/cache"
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

func CreateFlag() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req CreateFlagRequest
		if err := ctx.BindJSON(&req); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ResponseGenerator("Bad data provided", false))
			return
		}
		uid := ctx.MustGet("uid").(string)
		pid := string(ctx.Param("pid"))
		flag := modals.Featureflag{
			CreatedBy: uid,
			Flag:      req.Flag,
			Value:     req.Value,
			Pid:       pid,
		}
		if err := db.Store.CreateFlag(&flag); err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, utils.ResponseGenerator("Some Error Occurred", false))
			return
		}

		// creating new cache
		if err := cache.Redisdb.Set(fmt.Sprintf("FLAG-%s", flag.Fid), flag); err != nil {
			log.Print(err)
		}

		ctx.JSON(http.StatusOK, utils.ResponseGenerator("Flag Created Successfully", true))

	}
}

func GetFlagByFid() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fid := string(ctx.Param("fid"))
		val, err := cache.Redisdb.Get(fid)
		if err == nil {
			ctx.JSON(http.StatusOK, utils.ResponseGenerator(val, true))
			return
		}
		flagval, err := db.Store.GetFlagByFid(fid)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ResponseGenerator("Bad request", false))
			return
		}
		ctx.JSON(http.StatusOK, utils.ResponseGenerator(flagval, true))
	}
}
