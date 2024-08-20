package handler

import (
	"log"
	"net/http"

	"gihtub.com/heyyakash/switchr/internal/db"
	"gihtub.com/heyyakash/switchr/internal/utils"
	"github.com/gin-gonic/gin"
)

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
