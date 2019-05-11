package router

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

var Handler = mux.NewRouter()
var BooksRouter *mux.Router
var ActivityRouter *mux.Router

func InitRouting() {
	Handler.Use(loggingMiddleware)

	api := Handler.PathPrefix("/api").Subrouter()
	BooksRouter = api.PathPrefix("/books").Subrouter()
	ActivityRouter = api.PathPrefix("/activities").Subrouter()

	BooksRoutesInit(BooksRouter)
	ActivitiesRotesInit(ActivityRouter)
}
