package middlewares

import (
	"github.com/gin-gonic/gin"
)

func CheckHeader() gin.HandlerFunc {
	return func(c *gin.Context) {
		checkToken := c.GetHeader("token")
		if checkToken == "" {
			c.AbortWithStatusJSON(407, gin.H{
				"error": "required header",
			})
			return
		} else {
			c.Next()
		}
	}
}