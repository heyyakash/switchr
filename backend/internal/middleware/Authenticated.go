package middleware

import (
	"log"
	"net/http"
	"time"

	"gihtub.com/heyyakash/switchr/internal/utils"
	"github.com/gin-gonic/gin"
)

// func Authenticated() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		token, err := ctx.Cookie("token")
// 		if err != nil {
// 			ctx.Status(http.StatusFound)
// 			return
// 		}
// 		log.Print("Token : ", token)
// 		refreshToken, err := ctx.Cookie("refreshtoken")
// 		if err != nil {
// 			ctx.Status(http.StatusFound)
// 			return
// 		}
// 		claims, valid, err := utils.DecodeJWT(token)
// 		if err != nil {
// 			// ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ResponseGenerator("token broken", false))
// 			ctx.Status(http.StatusFound)
// 			log.Print(err)
// 			return
// 		}
// 		if !valid {
// 			reftokenclaims, refreshTokenValid, err := utils.DecodeJWT(refreshToken)
// 			if reftokenclaims.Uid != claims.Uid {
// 				ctx.AbortWithStatus(http.StatusFound)
// 				return
// 			}
// 			if err != nil {
// 				// ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ResponseGenerator("token broken", false))
// 				ctx.Redirect(http.StatusFound, fmt.Sprintf("%s/login", utils.GetString("CLIENT_ORIGIN")))
// 				log.Print(err)
// 				return
// 			}
// 			if !refreshTokenValid {
// 				// ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ResponseGenerator("user logged out", false))
// 				ctx.Redirect(http.StatusFound, fmt.Sprintf("%s/login", utils.GetString("CLIENT_ORIGIN")))
// 				log.Print(err)
// 				return
// 			}
// 			jwt, _, err := utils.GenerateJWT(reftokenclaims.Uid)
// 			if err != nil {
// 				// ctx.AbortWithStatusJSON(http.StatusInternalServerError, utils.ResponseGenerator("Internal error", false))
// 				ctx.Redirect(http.StatusFound, fmt.Sprintf("%s/login", utils.GetString("CLIENT_ORIGIN")))
// 				log.Print(err)
// 				return
// 			}
// 			tokenCookie := &http.Cookie{
// 				Name:     "token",
// 				Path:     "/",
// 				Value:    jwt,
// 				Domain:   "localhost",
// 				Expires:  time.Now().Add(1 * time.Hour),
// 				Secure:   false,
// 				HttpOnly: true,
// 				SameSite: http.SameSiteLaxMode,
// 			}
// 			http.SetCookie(ctx.Writer, tokenCookie)
// 		}
// 		if err != nil {

// 		}
// 		ctx.Set("uid", claims.Uid)
// 		ctx.Next()
// 	}
// }

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
				ctx.AbortWithStatusJSON(http.StatusFound, utils.ResponseGenerator("Broken Token", false))
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

		}
		ctx.Set("uid", claims.Uid)
		ctx.Next()
	}
}
