package db

import (
	"database/sql"
	"os"
)

type Person struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}

func Connect(env string) (*sql.DB, error) {
	var pgInfo string
	if env == "test" {
		pgInfo = os.Getenv("PG_URL_TEST")
	} else  {
		pgInfo = os.Getenv("PG_URL")
	}
	db, err := sql.Open("postgres", pgInfo)
	if err != nil {
		return nil, err
	}
	return db, nil
}