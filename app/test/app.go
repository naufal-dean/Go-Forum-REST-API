package test

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/core"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func NewTestApp() (*core.App, error) {
	// Create new app
	a := &core.App{}

	// Router (actually not used the test)
	a.Router = mux.NewRouter()

	// Validate object
	a.Validate = validator.New()

	// Database object
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=%s",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"), os.Getenv("DB_DBNAME"), os.Getenv("DB_TIMEZONE"))
	// Open connection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, errors.New("database connection failed")
	}
	// Save to app.DB
	a.DB = db

	// Return
	return a, nil
}
