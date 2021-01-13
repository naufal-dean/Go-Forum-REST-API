// Package app wraps the app functionalities.
package app

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/core"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/model/orm"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/route"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/seed"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"reflect"
	"strings"
)

func InitApp(a *core.App, flags map[string]bool) {
	initAppDB(a, flags["db-seed"], flags["db-fresh"])
	initAppRouter(a)
	initAppValidate(a)
}

func initAppDB(a *core.App, seedFlag bool, freshFlag bool) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=%s",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_DBNAME"), os.Getenv("DB_TIMEZONE"))
	// Open connection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Database connection failed")
	}
	// Save to app.DB
	a.DB = db
	// Reset database
	if freshFlag {
		a.DB.Migrator().DropTable(orm.Models...)
	}
	// Setup auto migrate
	err = a.DB.AutoMigrate(orm.Models...)
	if err != nil {
		log.Fatal("Database setup failed")
	}
	// Seed database
	if seedFlag {
		seed.Run(a.DB)
	}
}

func initAppRouter(a *core.App) {
	a.Router = mux.NewRouter()
	route.Setup(a)
}

func initAppValidate(a *core.App) {
	a.Validate = validator.New()
	a.Validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		// Get JSON name for vErr.Field()
		name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
}
