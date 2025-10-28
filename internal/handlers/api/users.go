package api

import (
	"github.com/gofiber/fiber/v2"

	"github.com/astianmuchui/nexthings-core/internal/models"
	"github.com/astianmuchui/nexthings-core/internal/schemas"

)

func UserApiRegisterHandler(c *fiber.Ctx) error {

	var payload schemas.UserRegisterRequest
	err := c.BodyParser(&payload)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Invalid Request Format",
			"message": err,
		})
	}

	var user models.User
	user.FirstName = payload.FirstName
	user.LastName = payload.LastName
	user.Username = payload.Username
	user.Email = payload.Email
	user.PhoneNumber = payload.PhoneNumber
	user.Country = payload.Country
	user.City = payload.City

	// Check if the user exists
	err = user.Retreive()

	if err == nil {
		// User exists
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Email, Phone or Username Already Exists",
		})
	}

	err = user.Create()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Unable to create user",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"user": user,
	})
}
