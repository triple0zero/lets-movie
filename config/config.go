package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func GetEnvVariable(key string) (val string) {

	if err := godotenv.Load(); err != nil {

		log.Print("No .env file found")
	}

	val, ok := os.LookupEnv(key)
	if !ok {
		log.Printf("%s is not set", key)
		return
	} else {
		//os.Setenv(key, val)
		return os.Getenv(key)
	}
}
