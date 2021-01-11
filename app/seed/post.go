package seed

import (
	"fmt"
	"github.com/bxcodec/faker/v3"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/model/orm"
	"gorm.io/gorm"
	"math/rand"
)

func PostRun(db *gorm.DB) {
	fmt.Println("[+] Posts table seeder started...")

	// Pluck user id
	var userIDs []uint
	db.Model(&orm.User{}).Pluck("id", &userIDs)

	// Pluck thread id
	var threadIDs []uint
	db.Model(&orm.Thread{}).Pluck("id", &threadIDs)

	// Generate data
	var threads []orm.Post
	for i := 0; i < 100; i++ {
		// Create empty object
		threads = append(threads, orm.Post{
			Title:    faker.Sentence(),
			Content:  faker.Paragraph(),
			UserID:   userIDs[rand.Intn(len(userIDs))],
			ThreadID: threadIDs[rand.Intn(len(threadIDs))],
		})
	}

	// Create record
	err := db.Create(&threads).Error
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("[+] Posts table seeder finished...")
}
