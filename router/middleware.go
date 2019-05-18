package router

import (
	"encoding/json"
	"fmt"
	_ "github.com/dgrijalva/jwt-go"
	"github.com/frayeralex/go-api/config"
	"github.com/frayeralex/go-api/middleware"
	"log"
	"net/http"
	"strings"
)

var hmacSampleSecret = []byte(config.JWT_SECRET)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		token := r.Header.Get("Authorization")
		if token == "" {
			w.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(w).Encode(map[string]string{
				"msg": "Authorization header required",
			})
			return
		}

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

func jsonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
