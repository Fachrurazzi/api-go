package auth

import (
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"os"
)

func ExtractToken(jwtToken string) (token *jwt.Token, err error) {
	getSecret := os.Getenv("JWT_SECRET")
	token, err = jwt.Parse(jwtToken, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
		}
		return []byte(getSecret), nil
	})
	return
}