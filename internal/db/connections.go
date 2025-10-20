package db

import (
	"github.com/gofiber/fiber/v2/log"
	"gorm.io/driver/postgres"
	sqlite "github.com/glebarez/sqlite"
	"gorm.io/gorm"

	"github.com/astianmuchui/nexthings-core/internal/env"

)

var DatabaseName string
var WorkingEnv string

var DB *gorm.DB

func init() {
	env.Load()
	WorkingEnv = env.GetWorkingEnvironment()
}

func Connect() {
	if WorkingEnv == env.PRODUCTION {
		// Use Postgres Connection
		dsn, dsnErr := env.GetDatabaseDSN()

		if dsnErr != nil {
			log.Error("Unable to get DSN")
		}

		var err error
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

		if err != nil {
			log.Error("Unable to connect to Postgres database: ", err)
		}

	} else {
		var err error
		DB, err = gorm.Open(sqlite.Open("nexthings.sqlite"), &gorm.Config{})

		if err != nil {
			log.Error("Failed to connect to SQLite database: ", err)
		}
	}
}
