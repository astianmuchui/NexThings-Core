package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"github.com/astianmuchui/nexthings-core/internal/env"

)

func init() {
	env.Load()
}

func main() {
	var app *fiber.App = fiber.New(fiber.Config{
		Prefork: true,
		ServerHeader: "NexThings",
		AppName: "NexThings IoT Core",
	})

	app.Use(logger.New())
	app.Use(recover.New())

	var listenPort int
	var address string
	listenPort, err := env.GetHttpListenPort()

	if err != nil {
		log.Error("Could not read HTTP Port from .env, Using default port: %v", env.DEFAULT_PORT)
	}


	address = fmt.Sprintf(":%d", listenPort)

	app.Listen(address)
}
