package quotes

import (
	"api-go/pkg/helpers/auth"
	"context"
	"database/sql"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

type UserQuote struct {
	ID float64 `json:"id"`
	Quote string `json:"quote"`
	Author string `json:"author"`
	CreatedAt time.Time `json:"created_at"`
}

func GetAllFavorites(db *sql.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		ctx := context.Background()
		var (
			getToken = c.GetHeader("token")
			userID float64
		)

		token, _ := auth.ExtractToken(getToken)
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			iAreaID := claims["id"].(float64)
			iAreaID, _ = claims["id"].(float64)
			userID = iAreaID
		}

		rows, err := db.QueryContext(ctx, "SELECT id, quotes, author, created_at FROM users_quotes WHERE user_id = $1", userID)
		if err != nil {
			c.AbortWithStatusJSON(400, gin.H{
				"error": err.Error(),
			})
		}
		defer rows.Close()

		quotes := make([]UserQuote, 0)
		for rows.Next() {
			var (
				quote UserQuote
			)
			if err := rows.Scan(&quote.ID, &quote.Quote, &quote.Author, &quote.CreatedAt); err != nil {
				c.AbortWithStatusJSON(400, gin.H{
					"error": err.Error(),
				})
				return
			}
			quotes = append(quotes, quote)
		}

		if err := rows.Err(); err != nil {
			c.AbortWithStatusJSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, quotes)
	}
}