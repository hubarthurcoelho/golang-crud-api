package models

import (
	"time"

	"github.com/hubarthurcoelho/golang-crud-api/database"
	"gorm.io/gorm"
)

type User struct {
	ID        uint      `gorm:"AUTO_INCREMENT" json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateUserData struct {
	Username string `validate:"min=8"`
	Email    string `validate:"min=8"`
	Password string `validate:"min=8"`
}

func CreateUser(user *User) (tx *gorm.DB) {
	return database.DB.Create(user)
}
