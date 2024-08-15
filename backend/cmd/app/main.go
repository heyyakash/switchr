package main

import (
	"log"

	"gihtub.com/heyyakash/switchr/internal/db"
	"github.com/gin-gonic/gin"
)

func Init() {
	db.Init()
}

func main() {
	Init()
	r := gin.Default()
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})
	log.Print("Server Started on port 8080")
	r.Run(":8020")

}
