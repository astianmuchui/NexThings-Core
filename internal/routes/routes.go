package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/astianmuchui/nexthings-core/internal/handlers"

)

func GetRoutes(app *fiber.App) {
	app.Get("/", handlers.HomeHandler)
}