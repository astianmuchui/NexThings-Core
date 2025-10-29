package api

import (
    "github.com/gofiber/fiber/v2"
    "github.com/google/uuid"

    "github.com/astianmuchui/nexthings-core/internal/models"
    "github.com/astianmuchui/nexthings-core/internal/schemas"
    "github.com/astianmuchui/nexthings-core/internal/services/mail"
    "github.com/astianmuchui/nexthings-core/internal/utils"
    "github.com/astianmuchui/nexthings-core/internal/db"
    "time"
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

    pwd := []byte(payload.Password)
    user.Password = utils.PasswordHash(pwd)

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

    token := utils.JwtGenerateToken(&user)
    return c.Status(fiber.StatusCreated).JSON(fiber.Map{
        "token": token,
        "user":  user.ToResponse(),
    })
}

func UserApiLoginHandler(c *fiber.Ctx) error {
    var payload *schemas.UserLoginRequest

    if err := c.BodyParser(&payload); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error":   "Invalid Request Format",
            "message": err,
        })
    }
    var user models.User
    q := db.DB.Model(models.User{}).Where("email = ? OR username = ? OR phone_number = ?", payload.Identifier, payload.Identifier, payload.Identifier).First(&user)
    if q.Error != nil {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "error": "Invalid Credentials",
        })
    }

    verified := utils.PasswordVerify(user.Password, []byte(payload.Password))
    if !verified {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "error": "Invalid Credentials",
        })
    }

    // Verify user account if not already verified
    if user.EmailVerifiedAt.String() == "0001-01-01 00:00:00 +0000 UTC" || user.EmailVerifiedAt.String() == "" {
        url := utils.GetURL(c)
        var email mail.Email

        go func(url string, usr *models.User) {
            email.SendUserVerificationEmail(url, usr)
        }(url, &user)

        return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
            "error": "Email Not Verified",
        })
    }

    token := utils.JwtGenerateToken(&user)

    return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "token": token,
        "user":  user.ToResponse(),
    })

}

func UserApiVerifyAccountHandler(c *fiber.Ctx) error {

    uid := c.Params("uid")
    token := c.Params("token")

    var user = &models.User{Uuid: uuid.MustParse(uid), EmailVerifyToken: uuid.MustParse(token)}
    err := user.Retreive()

    if err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error": "User Not Found",
        })
    }

    if user.EmailVerifyToken.String() != token {
        return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
            "error": "invalid token",
        })
    }

    user.EmailVerifiedAt = time.Now().UTC()
    user.EmailVerifyToken = uuid.New()
    err = user.Update()

    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Unable to Verify User",
        })
    }

    return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "message": "User Updated Successfully",
    })
}

func UserApiResetPasswordHandler(c *fiber.Ctx) error {

    var req *schemas.UserPasswordResetRequest

    if err := c.BodyParser(req); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid Request Format",
        })
    }

    var user = models.User{Uuid: req.Uid}

    err := user.Retreive()
    if err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error": "User Not Found",
        })
    }

    if user.PasswordResetToken != req.Token {
        return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
            "error": "Invalid Token",
        })
    }

    pwd := utils.PasswordHash([]byte(req.NewPassword))
    user.Password = pwd

    user.PasswordResetToken = uuid.New()
    err = user.Update()

    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Unable to update user",
        })
    }

    return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "message": "Password Updated Successfully",
    })
}
