package routes

import (
	"gihtub.com/heyyakash/switchr/internal/handler"
	"gihtub.com/heyyakash/switchr/internal/middleware"
	"github.com/gin-gonic/gin"
)

func FlagRoutes(c *gin.Engine) {
	c.GET("/flags/pid/:pid", middleware.Authenticated(), handler.GetFlagByPid())
	c.POST("/flags/create", middleware.Authenticated(), handler.CreateFlag())
	// c.GET("/flags/pid/:pid", middleware.Authenticated(), handler.GetFlagByPid())
	// c.GET("/flags/pid/:pid", middleware.Authenticated(), handler.GetFlagByPid())
}
