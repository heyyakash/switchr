package routes

import (
	"gihtub.com/heyyakash/switchr/internal/handler"
	"gihtub.com/heyyakash/switchr/internal/middleware"
	"github.com/gin-gonic/gin"
)

func AccountRouter(c *gin.Engine) {
	c.POST("/user/create", handler.CreateNewAccount())
	c.POST("/user/login", handler.LoginUser())
	c.GET("/user", middleware.Authenticated(), handler.LoginUser())
	c.POST("/user/magic")
	// c.POST("/user/email", auth.EmailExists())
}
