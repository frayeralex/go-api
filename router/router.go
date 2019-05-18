package router

import (
	"github.com/gorilla/mux"
)

var Handler = mux.NewRouter()
var BooksRouter *mux.Router
var ActivityRouter *mux.Router
var AuthRouter *mux.Router

func InitRouting() {
	Handler.Use(loggingMiddleware)
	Handler.Use(jsonMiddleware)

	api := Handler.PathPrefix("/api").Subrouter()
	BooksRouter = api.PathPrefix("/books").Subrouter()
	ActivityRouter = api.PathPrefix("/activities").Subrouter()
	ActivityRouter.Use(authMiddleware)

	AuthRouter = api.PathPrefix("/auth").Subrouter()

	BooksRoutesInit(BooksRouter)
	ActivitiesRoutesInit(ActivityRouter)
	AuthRoutesInit(AuthRouter)
}
