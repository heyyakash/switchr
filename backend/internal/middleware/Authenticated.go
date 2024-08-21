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
		refreshToken, err := ctx.Cookie("refreshtoken")
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ResponseGenerator("refresh token not provided", false))
			return
		}
		claims, valid, err := utils.DecodeJWT(token)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ResponseGenerator("token broken", false))
			log.Print(err)
			return
		}
		if !valid {
			reftokenclaims, refreshTokenValid, err := utils.DecodeJWT(refreshToken)
			if reftokenclaims.Uid != claims.Uid {
				ctx.AbortWithStatus(http.StatusBadRequest)
				return
			}
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
			jwt, _, err := utils.GenerateJWT(reftokenclaims.Uid)
			if err != nil {
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, utils.ResponseGenerator("Internal error", false))
				log.Print(err)
				return
			}
			tokenCookie := &http.Cookie{
				Name:     "token",
				Path:     "/",
				Value:    jwt,
				Domain:   "localhost",
				Expires:  time.Now().Add(1 * time.Hour),
				Secure:   false,
				HttpOnly: true,
				SameSite: http.SameSiteLaxMode,
			}
			http.SetCookie(ctx.Writer, tokenCookie)
		}
		ctx.Set("uid", claims.Uid)
		ctx.Next()
	}
}
