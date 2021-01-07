package core

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type App struct {
	DB     *gorm.DB
	Router *mux.Router
}
