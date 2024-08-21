package routes

import (
	"gihtub.com/heyyakash/switchr/internal/handler"
	"gihtub.com/heyyakash/switchr/internal/middleware"
	"github.com/gin-gonic/gin"
)

func FlagRoutes(c *gin.Engine) {
	c.GET("/flags/pid/:pid", middleware.Authenticated(), handler.GetFlagByPid())
	c.POST("/flags/create", middleware.Authenticated(), middleware.IsVerified(), handler.CreateFlag())
	c.PATCH("/flags/:fid", middleware.Authenticated(), middleware.IsVerified(), handler.UpdateFlag())
	c.DELETE("/flags/:fid", middleware.Authenticated(), middleware.IsVerified(), handler.DeleteFlag())
	// c.GET("/flags/pid/:pid", middleware.Authenticated(), handler.GetFlagByPid())
}
