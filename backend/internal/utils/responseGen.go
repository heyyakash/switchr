package utils

import "github.com/gin-gonic/gin"

func ResponseGenerator(message interface{}, success bool) *gin.H {
	return &gin.H{
		"message": message,
		"success": success,
	}
}
