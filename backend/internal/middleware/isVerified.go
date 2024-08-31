package middleware

import (
	"log"
	"net/http"

	"gihtub.com/heyyakash/switchr/internal/db"
	"gihtub.com/heyyakash/switchr/internal/utils"
	"github.com/gin-gonic/gin"
)

func IsVerified() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uid := ctx.MustGet("uid").(string)
		res, err := db.Store.GetUserByUid(uid)
		if err != nil {
			log.Print(err)
			ctx.AbortWithStatusJSON(http.StatusForbidden, utils.ResponseGenerator("Forbidden", false))
			return
		}
		if !res.Verified {
			ctx.AbortWithStatusJSON(http.StatusForbidden, utils.ResponseGenerator("Resource cannot be accessed without verification. Kindly navigate to settings to verify your email", false))
			return
		}
		ctx.Next()
	}
}
