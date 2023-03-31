package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type BaseModel struct {
	Token     uuid.UUID `gorm:"primary_key;type:char(36);column:token"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (b *BaseModel) BeforeCreate(db *gorm.DB) error {
	b.Token = uuid.NewV4()
	return nil
}
