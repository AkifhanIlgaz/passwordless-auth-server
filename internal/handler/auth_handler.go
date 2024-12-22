package handler

import (
	"fmt"
	"net/http"

	"github.com/AkifhanIlgaz/passworless-auth-server/internal/model"
	"github.com/AkifhanIlgaz/passworless-auth-server/internal/service"
	"github.com/AkifhanIlgaz/passworless-auth-server/pkg/message"
	"github.com/AkifhanIlgaz/passworless-auth-server/pkg/response"
	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) Login(c echo.Context) error {
	var loginRequest model.LoginRequest

	if err := c.Bind(&loginRequest); err != nil {
		c.Logger().Error(err)
		return response.WithError(c, http.StatusBadRequest, message.InvalidRequest)
	}

	if err := c.Validate(&loginRequest); err != nil {
		c.Logger().Error(err)
		return response.WithError(c, http.StatusBadRequest, message.EmailRequired)

	}

	fmt.Println("Login request: ", loginRequest)

	return response.WithSuccess(c, http.StatusOK, echo.Map{})
}
