package db

import (
	"github.com/go-bongo/bongo"
	"os"
)

var MongoConfig = &bongo.Config{
	ConnectionString: os.Getenv("MONGODB_URI"),
	Database: os.Getenv("MONGODB_NAME"),
}
