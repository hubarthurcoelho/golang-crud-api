package migrations

import (
	"github.com/hubarthurcoelho/golang-crud-api/database"
	"github.com/hubarthurcoelho/golang-crud-api/models"
	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) error {
	if err := database.DB.AutoMigrate(&models.User{}); err != nil {
		return err
	}

	return nil
}
