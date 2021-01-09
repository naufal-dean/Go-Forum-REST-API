package orm

type Token struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	UserID    uint   `gorm:"not null" json:"user_id"`
	TokenUUID string `gorm:"size:255;not null;unique" json:"name"`
}
