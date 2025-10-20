package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"

)

type BaseModel struct {
	gorm.Model
	Uuid uuid.UUID `json:"uuid" gorm:"uniqueIndex;not null"`
}
