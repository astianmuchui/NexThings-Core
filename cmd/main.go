package main

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/idempotency"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"github.com/astianmuchui/nexthings-core/internal/db"
	"github.com/astianmuchui/nexthings-core/internal/env"
	"github.com/astianmuchui/nexthings-core/internal/routes"
	"github.com/astianmuchui/nexthings-core/internal/utils"

)

func init() {
	env.Load()
	db.Connect()
	utils.RunMigrations()
}

func main() {
	var app *fiber.App = fiber.New(fiber.Config{
		Prefork:      true,
		ServerHeader: "NexThings",
		AppName:      "NexThings IoT Core",
	})

	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(helmet.New())
	app.Use(cors.New())

	/* Minimal Rate Limiting */
	app.Use(limiter.New(limiter.Config{
		Max:               100,
		Expiration:        10 * time.Second,
		LimiterMiddleware: limiter.SlidingWindow{},
	}))

	/* Idempotency */

	app.Use(idempotency.New(idempotency.Config{
		Lifetime: 5 * time.Minute,
	}))


	routes.GetRoutes(app)

	var listenPort int
	var address string
	listenPort, err := env.GetHttpListenPort()

	if err != nil {
		log.Error("Could not read HTTP Port from .env, Using default port: %v", env.DEFAULT_PORT)
	}

	address = fmt.Sprintf(":%d", listenPort)

	app.Listen(address)
}
