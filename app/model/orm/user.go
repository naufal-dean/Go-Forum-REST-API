package orm

import (
	"gitlab.com/pinvest/internships/hydra/onboarding-dean/app/lib/hash"
	"gorm.io/gorm"
)

type User struct {
	BaseModel
	Email    string   `gorm:"size:255;not null;unique" json:"email"`
	Password string   `gorm:"size:100;not null" json:"-"`
	Name     string   `gorm:"size:255;not null;unique" json:"name"`
	Threads  []Thread `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	Posts    []Post   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	Tokens   []Token  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.Password = hash.MakePasswordHash(u.Password)
	return
}

func (u User) PasswordValid(password string) bool {
	return hash.CheckPasswordHash(password, u.Password)
}
