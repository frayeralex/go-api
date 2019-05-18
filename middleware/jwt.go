package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/frayeralex/go-api/config"
	"time"
)

var hmacSampleSecret = []byte(config.JWT_SECRET)

func GenerateJwt(id, name string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   id,
		"name": name,
		"nbf":  time.Now(),
	})

	return token.SignedString(hmacSampleSecret)
}

func ParseJwt(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return hmacSampleSecret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
