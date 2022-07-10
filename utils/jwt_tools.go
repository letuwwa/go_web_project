package utils

import (
	"github.com/golang-jwt/jwt"
	"time"
)

func GenerateJWT(username string) (string, error) {
	var mySigningKey = []byte(JWTSalt)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 3).Unix()
	return token.SignedString(mySigningKey)
}
