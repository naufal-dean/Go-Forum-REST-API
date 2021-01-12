package core

import (
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type App struct {
	DB       *gorm.DB
	Router   *mux.Router
	Validate *validator.Validate
}
