package main

import (
	"github.com/joho/godotenv"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app"
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
