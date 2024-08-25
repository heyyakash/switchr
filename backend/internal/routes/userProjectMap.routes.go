package routes

import (
	"gihtub.com/heyyakash/switchr/internal/handler"
	"gihtub.com/heyyakash/switchr/internal/middleware"
	"github.com/gin-gonic/gin"
)

func UserProjectMapRoutes(c *gin.Engine) {
	c.GET("/userprojectmap", middleware.Authenticated(), handler.GetUserProjects())
	c.POST("/userprojectmap/:pid", middleware.Authenticated(), handler.DeleteMembers())
	c.GET("/userprojectmap/members/:pid", middleware.Authenticated(), handler.ProjectMembers())
	c.POST("/userprojectmap/members/delete/:pid", middleware.Authenticated(), handler.DeleteMembers())
}
