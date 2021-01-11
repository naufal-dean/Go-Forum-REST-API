package seed

import (
	"fmt"
	"github.com/bxcodec/faker/v3"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/model/orm"
	"gorm.io/gorm"
)

func UserRun(db *gorm.DB) {
	fmt.Println("[+] Users table seeder started...")

	// Init users
	var users []orm.User
	users = append(users, orm.User{
		Name:     "Mr. First",
		Email:    "user@user.com",
		Password: "password",
	})

	// Generate fake data
	for i := 0; i < 50; i++ {
		// Create empty object
		users = append(users, orm.User{
			Name:     faker.Name(),
			Email:    faker.Email(),
			Password: "password",
		})
	}

	// Create record
	err := db.Create(&users).Error
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("[+] Users table seeder finished...")
}
