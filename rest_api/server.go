package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"os"
)

func main() {
	server := echo.New()
	port := os.Getenv("APP_PORT")
	server.HTTPErrorHandler = customHttpErrorHandler

	server.Use(middleware.BasicAuth(basicAuthMiddleware))

	server.POST("/api/v1/process", requestHandler)

	server.Logger.Fatal(server.Start(":" + port))
}
