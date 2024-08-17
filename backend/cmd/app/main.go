package main

import (
	"log"

	"gihtub.com/heyyakash/switchr/internal/cache"
	"gihtub.com/heyyakash/switchr/internal/db"
	"gihtub.com/heyyakash/switchr/internal/routes"
	"github.com/gin-gonic/gin"
)

func Init() {
	db.Init()
}

func InitRedis() {
	cache.Redisdb.ConnectRedis()
}

func InitRoutes(r *gin.Engine) {
	routes.AccountRouter(r)
}

func main() {
	Init()
	InitRedis()
	r := gin.Default()
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})
	log.Print("Server Started on port 8080")

	//initializing routes
	InitRoutes(r)

	//starting server
	r.Run(":8020")
}
