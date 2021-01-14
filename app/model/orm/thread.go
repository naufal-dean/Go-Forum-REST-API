package orm

type Thread struct {
	BaseModel
	Name        string `gorm:"not null" json:"name"`
	Description string `gorm:"not null" json:"description"`
	UserID      uint   `json:"user_id"`
	Posts       []Post `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
}
