package handler

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"gihtub.com/heyyakash/switchr/internal/cache"
	"gihtub.com/heyyakash/switchr/internal/db"
	"gihtub.com/heyyakash/switchr/internal/utils"
	"github.com/gin-gonic/gin"
)

func CreateToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pid := string(ctx.Param("pid"))
		// uid := ctx.MustGet("uid").(string)
		uid := "6bb9c676-0e49-4d28-a5f9-21f162849c4d"
		_, err := db.Store.GetUserProjectMapByUidAndPid(uid, pid)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusForbidden, utils.ResponseGenerator("Forbidden", false))
			return
		}
		jwt, err := utils.GenerateApiJWTWithType(pid, "api-token", time.Now().Add(120*24*60*time.Minute).Unix())
		if err != nil {
			log.Print(err)
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, utils.ResponseGenerator("Some Error Occured", false))
		}
		ctx.JSON(http.StatusOK, utils.ResponseGenerator(jwt, true))
	}
}

func GetFlagFromAPI() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pid := ctx.MustGet("pid").(string)
		key := string(ctx.Param("key"))
		res, err := cache.Redisdb.Get(fmt.Sprintf("PID-%s-FLAG-%s", pid, key))
		if err != nil {
			val, err := db.Store.GetFlagByNameAndPid(key, pid)
			if err != nil {
				log.Print(err)
				ctx.AbortWithStatusJSON(http.StatusNotFound, utils.ResponseGenerator("Record not found", false))
				return
			}
			if err := cache.Redisdb.Set(fmt.Sprintf("PID-%s-FLAG-%s", pid, key), val); err != nil {
				log.Print("Redis err : ", err)
			}
			ctx.JSON(http.StatusOK, gin.H{"flag": val})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"flag": res})
	}
}
