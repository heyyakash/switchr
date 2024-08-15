package handler

import (
	"log"
	"net/http"
	"time"

	"gihtub.com/heyyakash/switchr/internal/db"
	"gihtub.com/heyyakash/switchr/internal/modals"
	"gihtub.com/heyyakash/switchr/internal/utils"
	"github.com/gin-gonic/gin"
)

func CreateNewAccount() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user modals.Users
		if err := ctx.BindJSON(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Bad json",
				"success": false,
			})
			return
		}
		hashedPass := utils.Hash(user.Password)
		user.Password = hashedPass
		if err := db.Store.CreateAccount(&user); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": err,
				"success": false,
			})
			return
		}
		token, refreshToken, err := utils.GenerateJWT(user.Uid)
		if err != nil {
			log.Print(err)
			ctx.JSON(http.StatusInternalServerError, utils.ResponseGenerator("Couldn't Generate Token", false))
			return
		}

		tokenCookie := &http.Cookie{
			Name:     "token",
			Path:     "/",
			Value:    token,
			Domain:   ctx.Request.Host,
			Expires:  time.Now().Add(1 * time.Hour),
			Secure:   true,
			HttpOnly: true,
			SameSite: http.SameSiteNoneMode,
		}
		refreshTokenCookie := &http.Cookie{
			Name:     "refreshtoken",
			Path:     "/",
			Value:    refreshToken,
			Domain:   ctx.Request.Host,
			Expires:  time.Now().Add(1 * time.Hour),
			Secure:   true,
			HttpOnly: true,
			SameSite: http.SameSiteNoneMode,
		}
		http.SetCookie(ctx.Writer, tokenCookie)
		http.SetCookie(ctx.Writer, refreshTokenCookie)
		ctx.JSON(http.StatusOK, utils.ResponseGenerator("Success", true))
	}
}

func LoginUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req modals.Users
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, utils.ResponseGenerator("Bad Fields Provided", false))
			return
		}
		user, err := db.Store.GetUserByEmail(req.Email)
		if err != nil {
			log.Print(err)
			ctx.JSON(http.StatusBadRequest, utils.ResponseGenerator("Wrong Credentials", false))
			return
		}
		check := utils.CheckPassword(user.Password, req.Password)
		if !check {
			ctx.JSON(http.StatusBadRequest, utils.ResponseGenerator("Wrong Credentials", false))
			return
		}
		token, refreshToken, err := utils.GenerateJWT(req.Email)
		if err != nil {
			log.Print(err)
			ctx.JSON(http.StatusInternalServerError, utils.ResponseGenerator("Couldn't Generate Token", false))
			return
		}
		tokenCookie := &http.Cookie{
			Name:     "token",
			Path:     "/",
			Value:    token,
			Domain:   ctx.Request.Host,
			Expires:  time.Now().Add(1 * time.Hour),
			Secure:   true,
			HttpOnly: true,
			SameSite: http.SameSiteNoneMode,
		}
		refreshTokenCookie := &http.Cookie{
			Name:     "refreshtoken",
			Path:     "/",
			Value:    refreshToken,
			Domain:   ctx.Request.Host,
			Expires:  time.Now().Add(7 * 24 * time.Hour),
			Secure:   true,
			HttpOnly: true,
			SameSite: http.SameSiteNoneMode,
		}
		http.SetCookie(ctx.Writer, tokenCookie)
		http.SetCookie(ctx.Writer, refreshTokenCookie)
		ctx.JSON(http.StatusOK, utils.ResponseGenerator("Success", true))
	}
}

func GetUserByToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uid := ctx.MustGet("uid").(string)

		res, err := db.Store.GetUserByUid(uid)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ResponseGenerator("error", false))
		}

		ctx.JSON(http.StatusOK, utils.ResponseGenerator(res, true))

	}
}

func DeleteUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}

func UpdateUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}
