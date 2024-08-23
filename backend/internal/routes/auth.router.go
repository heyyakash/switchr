package routes

import (
	"gihtub.com/heyyakash/switchr/internal/handler"
	"gihtub.com/heyyakash/switchr/internal/middleware"
	"github.com/gin-gonic/gin"
)

func AccountRouter(c *gin.Engine) {
	c.GET("/roles/list", middleware.Authenticated(), handler.GetRolesList())
	c.POST("/user/create", handler.CreateNewAccount())
	c.POST("/user/login", handler.LoginUser())
	c.POST("/user/magic", handler.SendMagicLink())
	c.GET("/user/magic/verify/:token", handler.LoginViaMagicLink())
	c.GET("/user/verify/:token", handler.VerifyUser())
	c.POST("/user/verify", middleware.Authenticated(), handler.SendVerificationMail())
	c.GET("/user", middleware.Authenticated(), handler.GetUserByToken())
	c.POST("/user/logout", middleware.Authenticated(), handler.Logout())
}
