package orm

type Thread struct {
	BaseModel
	Name   string `gorm:"not null" json:"name"`
	UserID uint   `json:"user_id"`
	Posts  []Post
}
