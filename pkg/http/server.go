package http

import (
	"api-go/pkg/http/quotes"
	"api-go/pkg/http/users"
	"api-go/pkg/middlewares"
	"database/sql"
	"github.com/gin-gonic/gin"
)

func SetupServer(db *sql.DB, env string) *gin.Engine {
	var router *gin.Engine
	if env == "test" {
		router = gin.New()
		router.Use(gin.Recovery())
	} else {
		router = gin.Default()
	}
	r := router.Group("/v1")
	{
		r.GET("/quotes", quotes.GetQuote())
		r.GET("/userquotes", quotes.GetAllUserQuotes(db))
		r.GET("/user/quotes", middlewares.CheckHeader(), middlewares.ReadTokenHeader(db), quotes.GetAllFavorites(db))
		r.POST("/user/login", users.Login(db))
		r.POST("/user/create", users.CreateOneUser(db))		
		r.POST("/favoritequotes/:id", middlewares.CheckHeader(), middlewares.ReadTokenHeader(db), quotes.LovesQuote(db))
		r.DELETE("/deletequote/:id", middlewares.CheckHeader(), middlewares.ReadTokenHeader(db), quotes.DeleteOne(db))
	}
	return router
}