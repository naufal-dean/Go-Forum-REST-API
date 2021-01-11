// @Version 1.0.0
// @Title Onboarding Dean API
// @Description API created for onboarding process in Pinhome.
// @ContactName Naufal Dean A
// @ContactEmail naufal.dean@gmail.com
// @ContactURL https://github.com/naufal-dean
// @LicenseName MIT
// @LicenseURL https://en.wikipedia.org/wiki/MIT_License
// @Security AuthorizationHeader read write
// @SecurityScheme AuthorizationHeader apiKey header Authorization
package main

import (
	"flag"
	"github.com/joho/godotenv"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app"
	"log"
)

func init() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load environment variable")
	}
}

func main() {
	// Get command line flags
	seedFlagPtr := flag.Bool("db-seed", false, "Use database seeder")
	freshFlagPtr := flag.Bool("db-fresh", false, "Reset the database")
	// Parse flags
	flag.Parse()
	flags := map[string]bool{
		"db-seed":  *seedFlagPtr,
		"db-fresh": *freshFlagPtr,
	}
	// Execute app
	app.Exec(flags)
}
