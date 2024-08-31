package middleware

import (
	"log"
	"net/http"
	"time"

	"gihtub.com/heyyakash/switchr/internal/utils"
	"github.com/gin-gonic/gin"
)

func Authenticated() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := ctx.Cookie("token")
		if err != nil || len(token) == 0 {
			ctx.AbortWithStatusJSON(http.StatusFound, utils.ResponseGenerator("Broken Token", false))
			return
		}
		claims, valid, err := utils.DecodeJWT(token)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusFound, utils.ResponseGenerator("token broken", false))
			log.Print(err)
			return
		}
		if claims.Type != "auth" {
			ctx.AbortWithStatusJSON(http.StatusFound, utils.ResponseGenerator("Wrong token", false))
			return
		}
		if !valid {
			refreshtoken, err := ctx.Cookie("refreshtoken")
			if err != nil || len(refreshtoken) == 0 {
				ctx.AbortWithStatusJSON(http.StatusFound, utils.ResponseGenerator("User Session over", false))
				return
			}
			refclaims, valid, err := utils.DecodeJWT(token)
			if err != nil || !valid {
				ctx.AbortWithStatusJSON(http.StatusFound, utils.ResponseGenerator("refresh token broken", false))
				log.Print(err)
				return
			}
			if refclaims.Uid != claims.Uid {
				ctx.AbortWithStatusJSON(http.StatusFound, utils.ResponseGenerator("refresh token broken", false))
				return
			}
			jwt, _, err := utils.GenerateJWT(refclaims.Uid)
			if err != nil {
				ctx.AbortWithStatusJSON(http.StatusFound, utils.ResponseGenerator("Internal error", false))
				log.Print(err)
				return
			}
			cookie := utils.CreateCookie("token", jwt, time.Now().Add(1*time.Hour))
			http.SetCookie(ctx.Writer, cookie)
			ctx.Set("uid", refclaims.Uid)
			ctx.Next()
		} else {
			ctx.Set("uid", claims.Uid)
			ctx.Next()
		}
	}
}
