package models

import (
	"github.com/hubarthurcoelho/golang-crud-api/database"
	"gorm.io/gorm"
)

type User struct {
	Username string `json:"username" gorm:"column:username;type:varchar(255)"`
	Xd       string `json:"xd" gorm:"column:xd;type:varchar(255)"`
	Email    string `json:"email" gorm:"column:email;type:varchar(255)"`
	Password string `json:"password" gorm:"column:password;type:varchar(255)"`
	BaseModel
}

type CreateUserData struct {
	Username string `validate:"min=8"`
	Email    string `validate:"min=8"`
	Password string `validate:"min=8"`
}

func CreateUser(user *User) (tx *gorm.DB) {
	return database.DB.Create(user)
}
