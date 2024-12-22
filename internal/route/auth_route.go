package route

import (
	"github.com/AkifhanIlgaz/passworless-auth-server/internal/handler"
	"github.com/labstack/echo/v4"
)

type AuthRoute struct {
	authHandler *handler.AuthHandler
}

func NewAuthRoute(authHandler *handler.AuthHandler) *AuthRoute {
	return &AuthRoute{authHandler: authHandler}
}

func (r *AuthRoute) Register(rg *echo.Group) {
	auth := rg.Group("/auth")

	auth.POST("/login/email", r.authHandler.LoginWithEmail)
	auth.POST("/login/phone", r.authHandler.LoginWithPhone)
}
