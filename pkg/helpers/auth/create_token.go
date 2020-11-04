package auth

import (
	jwt "github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

func CreateToken(id int, email string) (str string, timeExp time.Time, err error) {
	week, _ := time.ParseDuration("7d")
	timeExp = time.Now().Add(week)
	claims := jwt.MapClaims{"authorized": true, "id": id, "email": email, "exp": timeExp.Unix()}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	getSecret := os.Getenv("JWT_SECRET")
	if str, err = token.SignedString([]byte(getSecret)); err != nil {
		return
	} else {
		return
	}
}