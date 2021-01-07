package orm

import "gorm.io/gorm"

type Thread struct {
	gorm.Model
	Name string `gorm:"not null" json:"name"`
	UserID  int    `json:"user_id"`
}
