package middleware

import (
	"net/http"

	"gihtub.com/heyyakash/switchr/internal/utils"
	"github.com/gin-gonic/gin"
)

func IsAPIAuthenticated() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := string(ctx.GetHeader("token"))
		decoded, valid, err := utils.DecodeJWT(token)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ResponseGenerator("Broken token", false))
			return
		}
		if !valid {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ResponseGenerator("Token Expired", false))
			return
		}
		if decoded.Type != "api-token" {
			ctx.AbortWithStatusJSON(http.StatusForbidden, utils.ResponseGenerator("Bad Token", false))
			return
		}
		ctx.Set("pid", decoded.Pid)
		ctx.Next()
	}
}
