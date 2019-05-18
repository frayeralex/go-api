package router

import (
	"github.com/gorilla/mux"
)

var Handler = mux.NewRouter()
var BooksRouter *mux.Router
var ActivityRouter *mux.Router
var AuthRouter *mux.Router

func InitRouting() {
	Handler.Use(LoggingMiddleware)

	api := Handler.PathPrefix("/api").Subrouter()
	BooksRouter = api.PathPrefix("/books").Subrouter()
	ActivityRouter = api.PathPrefix("/activities").Subrouter()
	ActivityRouter.Use(AuthMiddleware)

	AuthRouter = api.PathPrefix("/auth").Subrouter()

	BooksRoutesInit(BooksRouter)
	ActivitiesRoutesInit(ActivityRouter)
	AuthRoutesInit(AuthRouter)
}
