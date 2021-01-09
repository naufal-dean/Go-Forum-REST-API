package orm

type Post struct {
	BaseModel
	Title    string `gorm:"not null" json:"title"`
	Content  string `gorm:"not null" json:"content"`
	UserID   uint   `json:"user_id"`
	ThreadID uint   `json:"thread_id"`
}
