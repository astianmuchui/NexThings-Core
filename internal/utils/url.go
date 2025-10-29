package utils

import (
	"github.com/gofiber/fiber/v2"

)

func GetURL(c *fiber.Ctx) string {
	return c.Protocol() + "://" + c.Hostname()
}