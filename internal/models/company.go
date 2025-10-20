package models


type Company struct {
	BaseModel

	Name string `json:"name" gorm:"uniqueIndex;not null"`
}