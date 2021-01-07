package orm

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"size:255;not null;unique" json:"email"`
	Password string `gorm:"size:100;not null" json:"-"`
	Name     string `gorm:"size:255;not null;unique" json:"name"`
	Threads  []Thread
	Posts    []Post
}
