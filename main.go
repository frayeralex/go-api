package main

import (
	"fmt"
	"github.com/frayeralex/go-api/config"
	"github.com/frayeralex/go-api/db"
	"github.com/frayeralex/go-api/router"
	"log"
	"net/http"
)

func main() {
	db.Connect()
	fmt.Println("CONNECTION SUCCESS")
	router.InitRouting()

	log.Fatal(http.ListenAndServe(":" + config.PORT, router.Handler))
}
