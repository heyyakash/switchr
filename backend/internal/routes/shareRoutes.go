package routes

import (
	"gihtub.com/heyyakash/switchr/internal/handler"
	"gihtub.com/heyyakash/switchr/internal/middleware"
	"github.com/gin-gonic/gin"
)

func ShareRoutes(c *gin.Engine) {
	c.POST("/share", middleware.Authenticated(), handler.ShareProject())
	c.GET("/share/confirm/:token", handler.ConfirmShareProject())
}
