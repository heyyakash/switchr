package handler

import (
	"log"
	"net/http"

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
		token, err := utils.GenerateJWT(user.Email)
		if err != nil {
			log.Print(err)
			ctx.JSON(http.StatusInternalServerError, utils.ResponseGenerator("Couldn't Generate Token", false))
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"message": token,
			"success": true,
		})
	}
}

func LoginUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req modals.Users
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, utils.ResponseGenerator("Bad Fields Provided", false))
			return
		}
		log.Print(req.Email)
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
		token, err := utils.GenerateJWT(req.Email)
		if err != nil {
			log.Print(err)
			ctx.JSON(http.StatusInternalServerError, utils.ResponseGenerator("Couldn't Generate Token", false))
			return
		}
		ctx.JSON(http.StatusOK, utils.ResponseGenerator(token, true))
	}
}

func GetUserByToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func DeleteUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}

func UpdateUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}
