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
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ResponseGenerator("token not provided", false))
			return
		}
		refreshToken, err := ctx.Cookie("refreshToken")
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ResponseGenerator("token not provided", false))
			return
		}
		uid, valid, err := utils.DecodeJWT(token)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ResponseGenerator("token broken", false))
			log.Print(err)
			return
		}
		if !valid {
			uid, refreshTokenValid, err := utils.DecodeJWT(refreshToken)
			if err != nil {
				ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ResponseGenerator("token broken", false))
				log.Print(err)
				return
			}
			if !refreshTokenValid {
				ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ResponseGenerator("user logged out", false))
				log.Print(err)
				return
			}
			jwt, _, err := utils.GenerateJWT(uid)
			if err != nil {
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, utils.ResponseGenerator("Internal error", false))
				log.Print(err)
				return
			}
			tokenCookie := &http.Cookie{
				Name:     "token",
				Path:     "/",
				Value:    jwt,
				Domain:   ctx.Request.Host,
				Expires:  time.Now().Add(1 * time.Hour),
				Secure:   true,
				HttpOnly: true,
				SameSite: http.SameSiteNoneMode,
			}
			http.SetCookie(ctx.Writer, tokenCookie)
		}
		ctx.Set("uid", uid)
		ctx.Next()
	}
}
