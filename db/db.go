package db

import (
	"github.com/frayeralex/go-api/config"
	"github.com/go-bongo/bongo"
	"log"
)

var Connection bongo.Connection

func Connect() {
	connection, err := bongo.Connect(config.MongoConfig)
	if err != nil {
		log.Fatal(err)
	}
	Connection = *connection
}
