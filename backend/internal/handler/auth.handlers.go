package handler

import (
	"fmt"
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
		_, err := db.Store.GetUserByEmail(user.Email)
		if err != nil {

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
				Domain:   "localhost",
				Expires:  time.Now().Add(1 * time.Hour),
				Secure:   false,
				HttpOnly: true,
				SameSite: http.SameSiteLaxMode,
			}
			refreshTokenCookie := &http.Cookie{
				Name:     "refreshtoken",
				Path:     "/",
				Value:    refreshToken,
				Domain:   "localhost",
				Expires:  time.Now().Add(7 * 24 * time.Hour),
				Secure:   false,
				HttpOnly: true,
				SameSite: http.SameSiteLaxMode,
			}
			http.SetCookie(ctx.Writer, tokenCookie)
			http.SetCookie(ctx.Writer, refreshTokenCookie)

			ctx.JSON(http.StatusOK, utils.ResponseGenerator("Success", true))
			return
		}
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ResponseGenerator("User already exists", false))

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
			Domain:   "localhost",
			Expires:  time.Now().Add(1 * time.Hour),
			Secure:   false,
			HttpOnly: true,
			SameSite: http.SameSiteLaxMode,
		}
		refreshTokenCookie := &http.Cookie{
			Name:     "refreshtoken",
			Path:     "/",
			Value:    refreshToken,
			Domain:   "localhost",
			Expires:  time.Now().Add(7 * 24 * time.Hour),
			Secure:   false,
			HttpOnly: true,
			SameSite: http.SameSiteLaxMode,
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

func SendMagicLink() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req modals.Users
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, utils.ResponseGenerator("Bad Fields Provided", false))
			return
		}
		user, err := db.Store.GetUserByEmail(req.Email)
		if err != nil {
			return
		}
		token, err := utils.GenerateJWTWithType(user.Email, "magic_link", time.Now().Add(5*time.Minute).Unix())
		if err != nil {
			log.Print(err)
			return
		}
		email := &modals.Email{
			To:      req.Email,
			Subject: "Magic Link",
			Content: fmt.Sprintf("Heyy! Your login link is as follows and is only valid for 5 minutes. \n%s", token),
		}
		if err := utils.SendEmail(email); err != nil {
			log.Print(err)
			return
		}
	}
}
func LoginViaMagicLink() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := string(ctx.Param("token"))
		claims, valid, err := utils.DecodeJWT(token)
		if err != nil {
			log.Print(err)
			return
		}

		if valid {
			log.Print("Valid")
			if claims.Type == "magic_link" {
				user, err := db.Store.GetUserByEmail(claims.Email)
				if err != nil {
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
					Domain:   "localhost",
					Expires:  time.Now().Add(1 * time.Hour),
					Secure:   false,
					HttpOnly: true,
					SameSite: http.SameSiteLaxMode,
				}
				refreshTokenCookie := &http.Cookie{
					Name:     "refreshtoken",
					Path:     "/",
					Value:    refreshToken,
					Domain:   "localhost",
					Expires:  time.Now().Add(7 * 24 * time.Hour),
					Secure:   false,
					HttpOnly: true,
					SameSite: http.SameSiteLaxMode,
				}
				http.SetCookie(ctx.Writer, tokenCookie)
				http.SetCookie(ctx.Writer, refreshTokenCookie)
				ctx.JSON(200, utils.ResponseGenerator("hei", true))
			}
		}
	}
}

func DeleteUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}

func UpdateUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}

func SendVerificationMail() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uid := ctx.MustGet("uid").(string)
		host := utils.GetString("HOST")
		user, err := db.Store.GetUserByUid(uid)
		if err != nil {
			log.Print(err)
			ctx.JSON(http.StatusInternalServerError, utils.ResponseGenerator("Error sending email!", false))
			return
		}
		token, err := utils.GenerateJWTWithType(user.Email, "verification", time.Now().Add(5*time.Minute).Unix())
		if err != nil {
			log.Print("Error : Could not generate JWT : ", err)
			ctx.JSON(http.StatusInternalServerError, utils.ResponseGenerator("Error sending email!", false))
			return
		}
		mail := &modals.Email{
			To:      user.Email,
			Subject: "Verification Email",
			Content: fmt.Sprintf("Heyy!! Your verification link is below \n%s/user/verify/%s", host, token),
		}
		if err := utils.SendEmail(mail); err != nil {
			log.Print(err)
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, utils.ResponseGenerator("Error sending email!", false))
			return
		}
		ctx.JSON(http.StatusOK, utils.ResponseGenerator("Email has been sent!", true))
	}
}

func VerifyUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Param("token")
		if token == "" {
			return
		}
		claims, valid, err := utils.DecodeJWT(token)
		if err != nil {
			log.Print(err)
			return
		}
		if !valid {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ResponseGenerator("Link has expired", false))
		}
		if claims.Type != "verification" {
			ctx.JSON(http.StatusBadRequest, utils.ResponseGenerator("Bad Request", false))
			return
		}
		user, err := db.Store.GetUserByEmail(claims.Email)
		if err != nil {
			log.Print(err)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ResponseGenerator("Some Error Occured", false))
			return
		}
		user.Verified = true
		if err := db.Store.UpdateUser(&user); err != nil {
			log.Print(err)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ResponseGenerator("Some Error Occured", false))
			return
		}
		ctx.AbortWithStatusJSON(http.StatusOK, utils.ResponseGenerator("User Verified Successfullu", true))

	}
}
