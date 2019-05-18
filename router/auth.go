package router

import (
	"github.com/frayeralex/go-api/api/auth"
	"github.com/gorilla/mux"
)

func AuthRoutesInit(r *mux.Router) {
	r.HandleFunc("/login", auth.Login).Methods("POST")
	r.HandleFunc("/register", auth.Register).Methods("POST")
}
