package routes

import (
	"gihtub.com/heyyakash/switchr/internal/handler"
	"gihtub.com/heyyakash/switchr/internal/middleware"
	"github.com/gin-gonic/gin"
)

func ApiRoutes(c *gin.Engine) {
	c.GET("/api/create/:pid", middleware.Authenticated(), handler.CreateToken())
	c.GET("/api/get/:key", middleware.IsAPIAuthenticated(), handler.GetFlagFromAPI())
}
