package seed

import (
	"fmt"
	"github.com/bxcodec/faker/v3"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/model/orm"
	"gorm.io/gorm"
	"math/rand"
)

func ThreadRun(db *gorm.DB) {
	fmt.Println("[+] Threads table seeder started...")

	// Pluck user id
	var userIDs []uint
	db.Model(&orm.User{}).Pluck("id", &userIDs)

	// Generate data
	var threads []orm.Thread
	for i := 0; i < 25; i++ {
		// Create empty object
		threads = append(threads, orm.Thread{
			Name:   faker.Word(),
			UserID: userIDs[rand.Intn(len(userIDs))],
		})
	}

	// Create record
	err := db.Create(&threads).Error
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("[+] Threads table seeder finished...")
}
