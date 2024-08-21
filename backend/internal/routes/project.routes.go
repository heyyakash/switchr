package routes

import (
	"gihtub.com/heyyakash/switchr/internal/handler"
	"gihtub.com/heyyakash/switchr/internal/middleware"
	"github.com/gin-gonic/gin"
)

func ProjectRoutes(c *gin.Engine) {
	c.POST("/project/create", middleware.Authenticated(), handler.CreateProject())
	c.GET("/project/:pid", middleware.Authenticated(), handler.GetProject())
}
