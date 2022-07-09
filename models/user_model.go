package models

type User struct {
	BaseModel
	Name     string `json:"name" gorm:"not null"`
	Surname  string `json:"surname"`
	Username string `json:"username" gorm:"not null;unique"`
	Password string `json:"-" gorm:"not null"`
}
