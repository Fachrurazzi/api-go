package main

import (
	postgres "api-go/pkg/db"
	"api-go/pkg/http"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func init() {
	if err := godotenv.Load("../.env"); err != nil {
		panic(err.Error())
	}
}

func main() {
	db, err := postgres.Connect(""); if err != nil {
		panic(err)
	}
	defer db.Close()
	_ = http.SetupServer(db, "").Run(":3000")
}