package db

import "os"

const (
	DATABASE        = "Kun"
	USER_COLLECTION = "users"
)

var MONGO_URL = os.Getenv("MONGO_URL")
