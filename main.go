package main

import (
	"github.com/joho/godotenv"
	"github.com/naufal-dean/onboarding-dean-local/app"
	"log"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load environment variable")
	}
}

func main() {
	app.Exec()
}
