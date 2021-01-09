package orm

type User struct {
	BaseModel
	Email    string `gorm:"size:255;not null;unique" json:"email"`
	Password string `gorm:"size:100;not null" json:"-"`
	Name     string `gorm:"size:255;not null;unique" json:"name"`
	Threads  []Thread
	Posts    []Post
	Tokens   []Token
}
