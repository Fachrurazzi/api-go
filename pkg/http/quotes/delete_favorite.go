package quotes

import (
	"api-go/pkg/helpers/auth"
	"context"
	"database/sql"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"strings"
)

func DeleteOne(db *sql.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var (
			idParams = c.Param("id")
			getToken = c.GetHeader("token")
			userID float64
			id int
			ctx = context.Background()
		)
		idQuote := strings.Split(idParams, ",")

		tx, err := db.BeginTx(ctx, nil)
		if err != nil {
			c.AbortWithStatusJSON(400, gin.H{
				"error": err.Error(),
			})
		}

		token, _ := auth.ExtractToken(getToken)
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			iAreaID := claims["id"].(float64)
			iAreaID, _ = claims["id"].(float64)
			userID = iAreaID
		}

		err = db.QueryRow("SELECT id FROM users_quotes WHERE user_id = $1 AND id = any($2)", userID, pq.Array(idQuote)).Scan(&id)
		if err == sql.ErrNoRows || err != nil {
			c.AbortWithStatusJSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}

		_, err = db.ExecContext(ctx, "DELETE FROM users_quotes WHERE user_id = $1 AND id = any($2)", userID, pq.Array(idQuote))
		if err != nil {
			_ = tx.Rollback()
			c.AbortWithStatusJSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		err = tx.Commit(); if err != nil {
			c.AbortWithStatusJSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, true)
	}
}