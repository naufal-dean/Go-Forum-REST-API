package seed

import (
	"fmt"
	"github.com/bxcodec/faker/v3"
	"gorm.io/gorm"
)

func Run(db *gorm.DB) {
	fmt.Println("[+] Start seeding table...")

	faker.SetGenerateUniqueValues(true)

	UserRun(db)
	ThreadRun(db)
	PostRun(db)

	fmt.Println("[+] Seeding table completed...")
	fmt.Println()
}
