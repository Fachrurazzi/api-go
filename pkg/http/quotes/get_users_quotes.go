package quotes

import (
	"context"
	"database/sql"
	"github.com/gin-gonic/gin"
	"time"
)

type UserQuotes struct {
	IdOfQuote int `json:"id_of_quote"`
	Name string `json:"user_name"`
	Quote string `json:"quote"`
	QuoteAuthor string `json:"author_quote"`
	QuoteCreated time.Time `json:"quote_created"`
}

func GetAllUserQuotes(db *sql.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var (
			ctx = context.Background()
		)

		const query = `SELECT
							uq.id as id_of_quote, name as user_name,
							uq.quotes as quote, uq.author as author_quote,
							uq.created_at as quote_created
						FROM users
						INNER JOIN users_quotes uq
						ON users.id = uq.user_id
					  `

		rows, err := db.QueryContext(ctx, query)
		if err != nil {
			c.AbortWithStatusJSON(400, gin.H{
				"error": err.Error(),
			})
		}
		defer rows.Close()

		quotesUsers := make([]UserQuotes, 0)
		for rows.Next() {
			var (
				quotes UserQuotes
			)
			if err := rows.Scan(&quotes.IdOfQuote, &quotes.Name, &quotes.Quote, &quotes.QuoteAuthor, &quotes.QuoteCreated); err != nil {
				c.AbortWithStatusJSON(400, gin.H{
					"error": err.Error(),
				})
				return
			}
			quotesUsers = append(quotesUsers, quotes)
		}

		if err:= rows.Err(); err != nil {
			c.AbortWithStatusJSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, quotesUsers)
	}
}