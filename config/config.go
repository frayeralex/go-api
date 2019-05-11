package config

import (
	"github.com/go-bongo/bongo"
	"os"
)

var MongoConfig = &bongo.Config{
	ConnectionString: os.Getenv("MONGO_URI"),
	Database: "test",
}

const (
	Activities = "activities"
)
