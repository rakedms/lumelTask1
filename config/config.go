package config

import "os"

func GetMongoURI() string {
	return os.Getenv("MONGO_URI")
}
