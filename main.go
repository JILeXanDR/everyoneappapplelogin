package main

import (
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	app := echo.New()
	app.Debug = true

	app.Use(middleware.Recover())
	app.Use(middleware.Logger())

	app.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, echo.Map{"message": "hello world"})
	})

	app.POST("/sign_in_with_apple", func(c echo.Context) error {
		return c.JSON(http.StatusOK, echo.Map{"message": "hello world"})
	})

	log.Fatal(app.Start(":" + os.Getenv("PORT")))
}
