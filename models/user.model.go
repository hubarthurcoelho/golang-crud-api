package models

import (
	"github.com/hubarthurcoelho/golang-crud-api/database"
	"gorm.io/gorm"
)

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
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
