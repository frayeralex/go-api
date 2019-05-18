package config

import (
	"os"
)

var PORT = os.Getenv("PORT")
var JWT_SECRET = os.Getenv("JWT_SECRET")
