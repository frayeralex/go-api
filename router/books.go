package router

import (
	"github.com/frayeralex/go-api/api/books"
	"github.com/gorilla/mux"
)

func BooksRoutesInit(r *mux.Router) {
	r.HandleFunc("", books.GetAll).Methods("GET")
	r.HandleFunc("", books.Create).Methods("POST")
	r.HandleFunc("/{id}", books.GetOne).Methods("GET")
	r.HandleFunc("/{id}", books.Update).Methods("PUT")
	r.HandleFunc("/{id}", books.Patch).Methods("PATCH")
	r.HandleFunc("/{id}", books.Delete).Methods("DELETE")
}