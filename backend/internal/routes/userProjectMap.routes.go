package routes

import (
	"gihtub.com/heyyakash/switchr/internal/handler"
	"gihtub.com/heyyakash/switchr/internal/middleware"
	"github.com/gin-gonic/gin"
)

func UserProjectMapRoutes(c *gin.Engine) {
	c.GET("/userprojectmap", middleware.Authenticated(), handler.GetUserProjects())
}
