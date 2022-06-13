package main

import (
	"crypto/subtle"
	"errors"
	"github.com/labstack/echo"
	"net/http"
	"os"
)

func basicAuthMiddleware(username, password string, context echo.Context) (bool, error) {
	user := os.Getenv("AUTH_USER")
	pass := os.Getenv("AUTH_SECRET")

	validateUsername := subtle.ConstantTimeCompare([]byte(username), []byte(user))
	validatePassword := subtle.ConstantTimeCompare([]byte(password), []byte(pass))

	if validateUsername == 1 && validatePassword == 1 {
		return true, nil
	}

	return false, newHTTPError(http.StatusUnauthorized, "Unauthorized", errors.New("unauthorized").Error())
}
