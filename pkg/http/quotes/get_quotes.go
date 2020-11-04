package quotes

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

type Quote struct {
	En string `json:"en"`
	Author string `json:"author"`
	Id string `json:"id"`
}

func GetQuote() func(c *gin.Context) {
	return func(c *gin.Context) {
		const API = "https://programming-quotes-api.herokuapp.com/quotes/random"
		var quote Quote

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
			return
		}
		c.JSON(200, quote)
	}
}