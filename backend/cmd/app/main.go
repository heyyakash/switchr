package main

import (
	"log"
	"time"

	"github.com/gin-contrib/cors"

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
	routes.UserProjectMapRoutes(r)
	routes.ProjectRoutes(r)
	routes.FlagRoutes(r)
	routes.ApiRoutes(r)
}

func main() {
	Init()
	InitRedis()
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{"PUT", "PATCH", "POST", "OPTIONS", "GET", "DELETE"},
		// AllowHeaders:     []string{"Origin", "auth-token", "content-type", "token"},
		AllowHeaders: []string{"*"},
		// ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:3000"
		},
		MaxAge: 12 * time.Hour,
	}))

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
