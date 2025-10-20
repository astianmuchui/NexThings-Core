package utils

import (
	"github.com/astianmuchui/nexthings-core/internal/db"
	"github.com/astianmuchui/nexthings-core/internal/models"

)

func RunMigrations() {
	db.DB.AutoMigrate(
		&models.User{},
	)
}
