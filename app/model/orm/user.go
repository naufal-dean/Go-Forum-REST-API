package orm

import (
	"github.com/pkg/errors"
	"github.com/naufal-dean/go-forum-rest-api/app/lib/hash"
	"gorm.io/gorm"
)

type User struct {
	BaseModel
	Email    string   `gorm:"size:255;not null;unique" json:"email"`
	Password string   `gorm:"size:100;not null" json:"-"`
	Name     string   `gorm:"size:255;not null" json:"name"`
	Threads  []Thread `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	Posts    []Post   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	Tokens   []Token  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.Password != "" {
		u.Password = hash.MakePasswordHash(u.Password)
	} else {
		return errors.New("password can not be empty")
	}
	return
}

func (u User) PasswordValid(password string) bool {
	return hash.CheckPasswordHash(password, u.Password)
}
