package router

import (
	"fmt"
	_ "github.com/dgrijalva/jwt-go"
	"github.com/frayeralex/go-api/config"
	"github.com/frayeralex/go-api/middleware"
	"log"
	"net/http"
	"strings"
)

var hmacSampleSecret = []byte(config.JWT_SECRET)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		tokenString := strings.Replace(r.Header.Get("Authorization"), "Bearer ", "", 1)

		claims, err := middleware.ParseJwt(tokenString)
		if err == nil {
			fmt.Println(claims)
			next.ServeHTTP(w, r)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
		}
	})
}
