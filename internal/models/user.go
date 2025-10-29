package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/astianmuchui/nexthings-core/internal/db"

)

type User struct {
	BaseModel

	FirstName                      string    `json:"firstname"`
	LastName                       string    `json:"lastname"`
	Username                       string    `json:"username" gorm:"uniqueIndex;not null"`

	Email                          string    `json:"email" gorm:"uniqueIndex;not null"`
	PhoneNumber                    string    `json:"phone" gorm:"uniqueIndex;not null"`
	City                           string    `json:"city"`
	Country                        string    `json:"country"`

	EmailVerifiedAt                time.Time `json:"EmailVerifiedAt"`
	PhoneVerifiedAt                time.Time `json:"PhoneVerifiedAt"`

	EmailVerifyToken               uuid.UUID `json:"EmailVerifyToken"`
	EmailVerificationCode          int32     `json:"EmailVerificationCode"`
	EmailVerificationCodeExpiresAt time.Time `json:"EmailVerificationCodeExpiresAt"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.Uuid = uuid.New()
	return nil
}

func (u *User) Create() (err error) {
	result := db.DB.Create(u)

	return result.Error
}

func (u *User) Update() (err error) {
	result := db.DB.Save(u)

	return result.Error
}

func (u *User) Retreive() error {
	result := db.DB.Model(User{}).Where("email = ? or uuid = ? or username = ? or phone_number = ?", u.Email, u.Uuid, u.Username, u.PhoneNumber).First(&u)

	return result.Error
}

func (u *User) Delete() error {
	result := db.DB.Delete(u)

	return result.Error
}
