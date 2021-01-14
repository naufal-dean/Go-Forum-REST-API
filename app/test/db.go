package test

import (
	"github.com/pkg/errors"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/core"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/model/orm"
	"gorm.io/gorm"
)

func DatabaseUp(a *core.App) error {
	// Migrate models
	err := a.DB.AutoMigrate(orm.Models...)
	if err != nil {
		panic(err)
	}
	// Seed database
	err = seedDB(a.DB)
	if err != nil {
		return err
	}
	// Succeed
	return nil
}

func seedDB(db *gorm.DB) error {
	// Create users record
	for _, user := range UsersData {
		err := db.Create(&user).Error
		if err != nil {
			return errors.New("failed to seed users table")
		}
	}
	// Create threads record
	for _, thread := range ThreadsData {
		err := db.Create(&thread).Error
		if err != nil {
			return errors.New("failed to seed threads table")
		}
	}
	// Create posts record
	for _, post := range PostsData {
		err := db.Create(&post).Error
		if err != nil {
			return errors.New("failed to seed posts table")
		}
	}
	// Create tokens record
	for _, token := range TokensData {
		err := db.Create(&token).Error
		if err != nil {
			return errors.New("failed to seed tokens table")
		}
	}
	// Succeed
	return nil
}

func DatabaseDown(a *core.App) {
	a.DB.Migrator().DropTable(orm.Models...)
}
