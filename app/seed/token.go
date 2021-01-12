package seed

import (
	"fmt"
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/model/orm"
	"gorm.io/gorm"
)

func TokenRun(db *gorm.DB)  {
	// Constant token
	// Used for testing convenience
	userID := uint(1)
	tokenUUID := "66f0ac33-f031-4ae5-8cae-cb5eef3536e6"
	// eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJPbmJvYXJkaW5nIERlYW4iLCJ1c2VyX2lkIjoxLCJ0b2tlbl91dWlkIjoiNjZmMGFjMzMtZjAzMS00YWU1LThjYWUtY2I1ZWVmMzUzNmU2In0.XH597NUrRydchWDpiw_ax104ymtldISH9riwNiQL7Rc

	// Create record
	err := db.Create(&orm.Token{UserID: userID, TokenUUID: tokenUUID}).Error
	if err != nil {
		fmt.Println(err)
	}
}
