package quotes

import (
	"api-go/pkg/helpers/auth"
	"context"
	"encoding/json"
	"fmt"
	"database/sql"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func LovesQuote(db *sql.DB) func (c *gin.Context) {
	return func(c *gin.Context) {
		idParams := c.Param("id")
		ctx := context.Background()
		var (
			API = "https://programming-quotes-api.herokuapp.com/quotes/id/" + idParams
			quote Quote
			getToken = c.GetHeader("token")
			userID float64
			quoteID string
		)

		tx, err := db.BeginTx(ctx, nil)
		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		res, err := http.Get(API); if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		response, _ := ioutil.ReadAll(res.Body)
		_ = res.Body.Close()
		err = json.Unmarshal(response, &quote); if err != nil {
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

		_ = db.QueryRow("SELECT quotes_id FROM users_quotes WHERE quotes_id = $1", idParams).Scan(&quoteID)
		if quoteID == idParams {
			c.AbortWithStatusJSON(400, gin.H{
				"error": "quote already exist on your favorite",
			})
			return
		}
		fmt.Printf("\n %v \n", quoteID)

		_, err = db.ExecContext(ctx, "INSERT INTO users_quotes(user_id, quotes_id, quotes, author) VALUES($1, $2, $3, $4)", userID, idParams, quote.En, quote.Author)
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