package main

import (
	"fmt"
	"log"

	"github.com/AkifhanIlgaz/passworless-auth-server/internal/config"
	"github.com/AkifhanIlgaz/passworless-auth-server/internal/handler"
	"github.com/AkifhanIlgaz/passworless-auth-server/internal/route"
	"github.com/AkifhanIlgaz/passworless-auth-server/internal/service"
	"github.com/AkifhanIlgaz/passworless-auth-server/pkg/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	prettyLogger "github.com/rdbell/echo-pretty-logger"
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

	e.Validator = validator.NewCustomValidator()

	e.Use(prettyLogger.Logger)
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(devCorsConfig))

	rg := e.Group("/api")

	authService := service.NewAuthService()

	authHandler := handler.NewAuthHandler(authService)

	authRoute := route.NewAuthRoute(authHandler)

	authRoute.Register(rg)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.Port)))
}
