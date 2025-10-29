//@formatter:on

package utils

import (
    "golang.org/x/crypto/bcrypt"
    "github.com/gofiber/fiber/v2/log"
)

func PasswordHash(pwd []byte) string {
    hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
    if err != nil {
        log.Info(err)
    }
    return string(hash)
}

func PasswordVerify(hashed string, raw_pwd []byte) bool {
    bytehash := []byte(hashed)
    err := bcrypt.CompareHashAndPassword(bytehash, raw_pwd)
    if err != nil {
        log.Info(err)
        return false
    }
    return true
}
