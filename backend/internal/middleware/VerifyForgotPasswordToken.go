package middleware

import (
	"log"
	"net/http"

	"gihtub.com/heyyakash/switchr/internal/utils"
	"github.com/gin-gonic/gin"
)

func VerifyForgotPasswordToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := ctx.Cookie("token")
		if err != nil || len(token) == 0 {
			ctx.AbortWithStatusJSON(http.StatusFound, utils.ResponseGenerator("Broken Token", false))
			return
		}
		claims, valid, err := utils.DecodeJWT(token)

		if err != nil {
			log.Print(err)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ResponseGenerator("Bad Request", false))
			return
		}
		if !valid {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ResponseGenerator("Sorry, the link has expired", false))
			return
		}
		if claims.Type != "forgot-password" {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ResponseGenerator("Bad request", false))
			return
		}
		ctx.Set("uid", claims.Uid)
		ctx.Next()
	}
}
