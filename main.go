package main

import (
	"github.com/frayeralex/go-api/config"
	"github.com/frayeralex/go-api/db"
	"github.com/frayeralex/go-api/router"
	"log"
	"net/http"
)

func main() {
	db.Connect()
	router.InitRouting()

	log.Fatal(http.ListenAndServe(":" + config.PORT, router.Handler))
}
