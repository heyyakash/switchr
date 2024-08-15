package routes

import (
	"gihtub.com/heyyakash/switchr/internal/handler"
	"github.com/gin-gonic/gin"
)

func AccountRouter(c *gin.Engine) {
	c.POST("/user/create", handler.CreateNewAccount())
	c.POST("/user/login", handler.LoginUser())
	// c.POST("/user/email", auth.EmailExists())
}
