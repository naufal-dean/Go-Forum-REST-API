// Package app wraps the app functionalities.
package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/naufal-dean/onboarding-dean-local/app/core"
	"github.com/naufal-dean/onboarding-dean-local/app/model/orm"
	"github.com/naufal-dean/onboarding-dean-local/app/route"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

var app *core.App

func initApp() {
	app = &core.App{}
	initAppDB()
	initAppRouter()
}

func initAppDB() {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=%s",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_DBNAME"), os.Getenv("DB_TIMEZONE"))
	// Open connection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Database connection failed")
	}
	// Save to app.DB
	app.DB = db
	// Setup auto migrate
	err = app.DB.AutoMigrate(orm.Models...)
	if err != nil {
		log.Fatal("Database setup failed")
	}
}

func initAppRouter() {
	app.Router = mux.NewRouter()
	route.Setup(app)
}

func Exec() {
	initApp()

	addr := ":8080"
	fmt.Printf("Server started at %s...\n", addr)
	http.ListenAndServe(addr, app.Router)
}
