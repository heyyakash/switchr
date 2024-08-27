package handler

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"gihtub.com/heyyakash/switchr/internal/cache"
	"gihtub.com/heyyakash/switchr/internal/constants"
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
			Content: fmt.Sprintf("Heyy! Your login link is as follows and is only valid for 5 minutes. \n%s/user/magic/verify/%s", utils.GetString("HOST"), token),
		}
		if err := utils.SendEmail(email); err != nil {
			log.Print(err)
			return
		}
		ctx.JSON(http.StatusOK, utils.ResponseGenerator("Email sent successfully", true))
	}
}
func LoginViaMagicLink() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := string(ctx.Param("token"))
		claims, valid, err := utils.DecodeJWT(token)
		if err != nil {
			log.Print(err)
			ctx.JSON(http.StatusBadRequest, utils.ResponseGenerator("Expired", false))
			return
		}

		if valid {
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
				ctx.Redirect(http.StatusFound, fmt.Sprintf("%s/dashboard", utils.GetString("CLIENT_ORIGIN")))
				return
			}
		}
		log.Print("Invalid")
		ctx.JSON(http.StatusBadRequest, utils.ResponseGenerator("Expired", false))
	}
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
		if err := db.Store.UpdateUser(&user, user.Uid); err != nil {
			log.Print(err)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ResponseGenerator("Some Error Occured", false))
			return
		}
		ctx.AbortWithStatusJSON(http.StatusOK, utils.ResponseGenerator("User Verified Successfullu", true))

	}
}

func Logout() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenCookie := &http.Cookie{
			Name:     "token",
			Path:     "/",
			Value:    "",
			Domain:   "localhost",
			Expires:  time.Unix(0, 0),
			MaxAge:   -1,
			Secure:   false,
			HttpOnly: true,
			SameSite: http.SameSiteLaxMode,
		}
		refreshTokenCookie := &http.Cookie{
			Name:     "refreshtoken",
			Path:     "/",
			Value:    "",
			Domain:   "localhost",
			Expires:  time.Unix(0, 0),
			MaxAge:   -1,
			Secure:   false,
			HttpOnly: true,
			SameSite: http.SameSiteLaxMode,
		}
		http.SetCookie(ctx.Writer, tokenCookie)
		http.SetCookie(ctx.Writer, refreshTokenCookie)
		ctx.JSON(http.StatusFound, utils.ResponseGenerator("Logged out", true))
	}
}

func GetRolesList() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		val, err := cache.Redisdb.Get("roles")
		if err != nil {
			_ = cache.Redisdb.Set("roles", constants.Role)
			ctx.JSON(http.StatusOK, utils.ResponseGenerator(constants.Role, true))
			return
		}
		ctx.JSON(http.StatusOK, utils.ResponseGenerator(val, true))

	}
}

func DeleteUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}

func UpdateUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		type updateUserStruct struct {
			FullName string `json:"fullname"`
			Uid      string `json:"uid"`
		}
		var req updateUserStruct
		if err := ctx.BindJSON(&req); err != nil {
			log.Print(err)
			ctx.JSON(http.StatusBadRequest, utils.ResponseGenerator("Bad Request", false))
			return
		}
		uid := ctx.MustGet("uid").(string)
		if uid != req.Uid {
			ctx.AbortWithStatusJSON(http.StatusForbidden, utils.ResponseGenerator("Forbidden", false))
			return
		}
		if err := db.Store.UpdateUser(&modals.Users{FullName: req.FullName, Uid: req.Uid}, req.Uid); err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, utils.ResponseGenerator("Some Error Occured", false))
			return
		}
		ctx.JSON(http.StatusOK, utils.ResponseGenerator("Updated Successfully", true))

	}
}
func ChangePassword() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		type ChangePasswordStruct struct {
			Current string `json:"current"`
			New     string `json:"new"`
		}
		var req ChangePasswordStruct
		if err := ctx.BindJSON(&req); err != nil {
			log.Print(err)
			ctx.JSON(http.StatusBadRequest, utils.ResponseGenerator("Bad Request", false))
			return
		}
		uid := ctx.MustGet("uid").(string)
		user, err := db.Store.GetUserByUidWithPassword(uid)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, utils.ResponseGenerator("User Not found", false))
			return
		}

		if checkPass := utils.CheckPassword(user.Password, req.Current); !checkPass {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ResponseGenerator("Entered current password is wrong", false))
			return
		}

		if checkNewHash := utils.CheckPassword(user.Password, req.New); checkNewHash {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ResponseGenerator("New and current password cannot be same", false))
			return
		}
		user.Password = utils.Hash(req.New)

		if err := db.Store.UpdateUser(&user, user.Uid); err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, utils.ResponseGenerator("Some Error Occured", false))
			return
		}
		ctx.JSON(http.StatusOK, utils.ResponseGenerator("Updated Successfully", true))

	}
}

func SendForgotPasswordLink() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		type ForgotPasswordReq struct {
			Email string `json:"email"`
		}
		var req ForgotPasswordReq
		if err := ctx.BindJSON(&req); err != nil {
			log.Print(err)
			ctx.JSON(http.StatusBadRequest, utils.ResponseGenerator("Bad Request", false))
			return
		}

		user, err := db.Store.GetUserByEmail(req.Email)
		if err != nil {
			log.Print(err)
			ctx.JSON(http.StatusBadRequest, utils.ResponseGenerator("Account does not exists", false))
			return
		}
		jwt, err := utils.GenerateJWTWithTypeAndUID(user.Uid, "forgot-password", time.Now().Add(5*time.Minute).Unix())
		if err != nil {
			log.Print(err)
			ctx.JSON(http.StatusInternalServerError, utils.ResponseGenerator("Some Error occured", false))
			return
		}
		log.Print("Token : ", jwt)
		email := modals.Email{
			To:      req.Email,
			Content: fmt.Sprintf("Heyy %s, Your link to reset your switchr account password is \n%s/changepass/%s \nThe link is active for 5 minutes only", user.FullName, utils.GetString("HOST"), jwt),
			Subject: "Change Switchr Account Password",
		}
		if err := utils.SendEmail(&email); err != nil {
			log.Print(err)
			ctx.JSON(http.StatusInternalServerError, utils.ResponseGenerator("Some Error occured", false))
			return
		}
		ctx.JSON(http.StatusOK, utils.ResponseGenerator("Kindly check your email for the reset password link", true))
	}
}

func RedirectToChangePassword() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := string(ctx.Param("token"))
		claims, valid, err := utils.DecodeJWT(token)
		if err != nil {
			log.Print(err)
			ctx.JSON(http.StatusBadRequest, utils.ResponseGenerator("Bad Request", false))
			return
		}
		if !valid {
			ctx.JSON(http.StatusBadRequest, utils.ResponseGenerator("Sorry, the link has expired", false))
			return
		}
		if claims.Type != "forgot-password" {
			ctx.JSON(http.StatusBadRequest, utils.ResponseGenerator("Bad request", false))
			return
		}
		tokenCookie := &http.Cookie{
			Name:     "token",
			Path:     "/",
			Value:    token,
			Domain:   "localhost",
			Expires:  time.Now().Add(5 * time.Minute),
			Secure:   false,
			HttpOnly: true,
			SameSite: http.SameSiteLaxMode,
		}
		http.SetCookie(ctx.Writer, tokenCookie)
		ctx.Redirect(http.StatusFound, fmt.Sprintf("%s/changepassword", utils.GetString("CLIENT_ORIGIN")))
	}
}

func ChangeForgotPassword() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		type ChangePasswordReq struct {
			New string `json:"new"`
		}
		var req ChangePasswordReq
		if err := ctx.BindJSON(&req); err != nil {
			log.Print(err)
			ctx.JSON(http.StatusBadRequest, utils.ResponseGenerator("Bad Request", false))
			return
		}

		uid := ctx.MustGet("uid").(string)
		user, err := db.Store.GetUserByUidWithPassword(uid)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, utils.ResponseGenerator("User Not found", false))
			return
		}

		if checkNewHash := utils.CheckPassword(user.Password, req.New); checkNewHash {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ResponseGenerator("New and current password cannot be same", false))
			return
		}
		user.Password = utils.Hash(req.New)

		if err := db.Store.UpdateUser(&user, user.Uid); err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, utils.ResponseGenerator("Some Error Occured", false))
			return
		}
		tokenCookie := &http.Cookie{
			Name:     "token",
			Path:     "/",
			Value:    "",
			Domain:   "localhost",
			Expires:  time.Unix(0, 0),
			MaxAge:   -1,
			Secure:   false,
			HttpOnly: true,
			SameSite: http.SameSiteLaxMode,
		}
		http.SetCookie(ctx.Writer, tokenCookie)
		ctx.JSON(http.StatusOK, utils.ResponseGenerator("Updated Successfully", true))
	}
}
