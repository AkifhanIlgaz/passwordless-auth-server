package main

import (
	"log"
	"net/http"

	"github.com/AkifhanIlgaz/passworless-auth-server/internal/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var devCorsConfig = middleware.CORSConfig{
	AllowOrigins: []string{"*"},
	AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	AllowHeaders: []string{"*"},
}

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	// TODO: Use the config
	_ = config

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(devCorsConfig))

	e.GET("/", func(c echo.Context) error {

		return c.JSON(http.StatusOK, map[string]any{
			"status":  "success",
			"message": "Hello World",
		})
	})

	e.Logger.Fatal(e.Start(":8000"))
}
