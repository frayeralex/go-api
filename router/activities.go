package router

import (
	"github.com/frayeralex/go-api/api/activities"
	"github.com/gorilla/mux"
)

func ActivitiesRotesInit (r *mux.Router) {
	r.HandleFunc("", activities.GetAll).Methods("GET").Queries()
	r.HandleFunc("", activities.Create).Methods("POST")
	r.HandleFunc("/{id}", activities.GetOne).Methods("GET")
	r.HandleFunc("/{id}", activities.Update).Methods("PUT")
	r.HandleFunc("/{id}", activities.Delete).Methods("DELETE")
}

