package utils

import (
    "github.com/astianmuchui/nexthings-core/internal/models"
    "time"
    "github.com/golang-jwt/jwt/v5"
    "github.com/google/uuid"
)

func JwtGenerateToken(user *models.User) string {
    token := jwt.New(jwt.SigningMethodHS256)

    claims := token.Claims.(jwt.MapClaims)
    claims["uuid"] = user.Uuid
    claims["email"] = user.Email
    claims["username"] = user.Username
    claims["exp"] = time.Now().Add(time.Hour * 72).Unix() // Token expires in 72 hours (3 days)

    t, err := token.SignedString([]byte(uuid.New().String()))
    if err != nil {
        return ""
    }

    return t
}
