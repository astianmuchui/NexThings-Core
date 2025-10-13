package db

import (
	"github.com/gofiber/fiber/v2/log"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/astianmuchui/nexthings-core/internal/env"

)

var DatabaseName string
var WorkingEnv string

func init() {
	env.Load()
	WorkingEnv = env.GetWorkingEnvironment()
}

func Connect() *gorm.DB  {
	var db *gorm.DB
	if WorkingEnv == env.LOCAL {
		// Use sqlite

		db, connErr := gorm.Open(sqlite.Open("nexthings.sqlite"), &gorm.Config{})

		if connErr != nil {
			log.Error("Failed to connect to database: ", connErr)
		}
		return db
	} else if WorkingEnv == env.PRODUCTION {
		// Use Postgres Connection

		dsn, dsnErr := env.GetDatabaseDSN()

		if dsnErr != nil {
			log.Error("Unable to get DSN")
		}

		db, connError := gorm.Open(postgres.Open(dsn), &gorm.Config{})

		if connError!= nil {
			log.Error("Unable to get DSN")
		}

		return db
	}
	return db
}
