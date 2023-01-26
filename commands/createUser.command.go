package commands

import (
	"github.com/hubarthurcoelho/golang-crud-api/models"
	"golang.org/x/crypto/bcrypt"
)

func CreateUserCommand(user *models.User) error {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	if err := models.CreateUser(user).Error; err != nil {
		return err
	}

	return nil
}
