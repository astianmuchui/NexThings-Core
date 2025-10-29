package middleware

import "github.com/gofiber/fiber/v2"

func JwtAuthMiddleware(c *fiber.Ctx) fiber.Handler {

    return func(c *fiber.Ctx) error {

        return c.Next()
    }
}
