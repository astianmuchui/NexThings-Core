package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/astianmuchui/nexthings-core/internal/models"
	"github.com/astianmuchui/nexthings-core/internal/schemas"
	"github.com/astianmuchui/nexthings-core/internal/services/mail"
	"github.com/astianmuchui/nexthings-core/internal/utils"

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
	user.EmailVerifyToken = uuid.New()

	err = user.Retreive()

	if err == nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": "Email, Phone or Username Already Exists",
		})
	}

	err = user.Create()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Unable to create user",
		})
	}

	var email mail.Email
	url := utils.GetURL(c)
	go func(url string, usr *models.User) {
		email.SendUserVerificationEmail(url, usr)
	}(url, &user)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"user": user,
	})
}
