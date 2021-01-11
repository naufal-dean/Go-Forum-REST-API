package orm

type User struct {
	BaseModel
	Email    string   `gorm:"size:255;not null;unique" json:"email"`
	Password string   `gorm:"size:100;not null" json:"-"`
	Name     string   `gorm:"size:255;not null;unique" json:"name"`
	Threads  []Thread `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	Posts    []Post   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	Tokens   []Token  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
}
