package middlewares

import (
	"api-go/pkg/helpers/auth"
	"database/sql"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func ReadTokenHeader(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		getToken := c.GetHeader("token")
		token, _ := auth.ExtractToken(getToken)
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			var ID float64
			var userID int

			switch theID := claims["id"].(type) {
			case float64:
				ID = theID
			}
			_ = db.QueryRow("SELECT id FROM users WHERE id = $1", ID).Scan(&userID)
				if userID < 1 {
					c.AbortWithStatusJSON(404, gin.H{
						"error": "Not Authorize",
					})
					return
				}
				c.Next()
		} else {
			c.AbortWithStatusJSON(400, gin.H{
				"error": "token expired, please login again",
			})
		return
		}
	}
}