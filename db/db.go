package db

import (
	"github.com/go-bongo/bongo"
	"log"
)

var Connection bongo.Connection

func Connect() {
	connection, err := bongo.Connect(MongoConfig)
	if err != nil {
		log.Fatal(err)
	}
	Connection = *connection
}
