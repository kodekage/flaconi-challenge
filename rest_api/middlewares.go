package main

import (
	"crypto/subtle"
	"errors"
	"github.com/labstack/echo"
	"net/http"
)

func basicAuthMiddleware(username, password string, context echo.Context) (bool, error) {
	validateUsername := subtle.ConstantTimeCompare([]byte(username), []byte("flaconi"))
	validatePassword := subtle.ConstantTimeCompare([]byte(password), []byte("secret"))

	if validateUsername == 1 && validatePassword == 1 {
		return true, nil
	}

	return false, newHTTPError(http.StatusUnauthorized, "Unauthorized", errors.New("unauthorized").Error())
}
