package main

import (
	"log"
	"time"

	"github.com/gin-contrib/cors"

	"gihtub.com/heyyakash/switchr/internal/cache"
	"gihtub.com/heyyakash/switchr/internal/constants"
	"gihtub.com/heyyakash/switchr/internal/db"
	"gihtub.com/heyyakash/switchr/internal/routes"
	"gihtub.com/heyyakash/switchr/internal/utils"
	"github.com/gin-gonic/gin"
)

func Init() {
	db.Init()
	constants.LoadRoleConstants()
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
	routes.ShareRoutes(r)
}

func main() {
	Init()
	InitRedis()
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{utils.GetString("CLIENT_ORIGIN")},
		AllowMethods: []string{"PUT", "PATCH", "POST", "OPTIONS", "GET", "DELETE"},
		AllowHeaders: []string{"Origin", "auth-token", "content-type", "token"},
		// AllowHeaders: []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == utils.GetString("CLIENT_ORIGIN")
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
	if utils.GetString("ENV") == "prod" {
		r.RunTLS(":8020", utils.GetString("CERTIFICATE"), utils.GetString("KEY"))
	} else {
		r.Run(":8020")
	}
}
