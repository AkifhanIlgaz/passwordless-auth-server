package main

func main() {

}package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var devCorsConfig = middleware.CORSConfig{
	AllowOrigins: []string{"*"},
	AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	AllowHeaders: []string{"*"},
}

func main() {
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
