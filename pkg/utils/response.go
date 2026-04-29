package utils

import "github.com/gin-gonic/gin"

func Ok(c *gin.Context, status int, data any, message string) {
	c.JSON(status, gin.H{
		"success": true,
		"data":    data,
		"message": message,
	})
}

func Success(c *gin.Context, status int, key string, data any, message string) {
	c.JSON(status, gin.H{
		"success": true,
		key:       data,
		"message": message,
	})
}

func Fail(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{
		"success": false,
		"error": gin.H{
			"code":    status,
			"message": message,
		},
	})
}
