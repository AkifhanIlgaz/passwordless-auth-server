package handler

import (
	"net/http"

	"github.com/AkifhanIlgaz/passworless-auth-server/internal/model"
	"github.com/AkifhanIlgaz/passworless-auth-server/internal/service"
	"github.com/AkifhanIlgaz/passworless-auth-server/pkg/message"
	"github.com/AkifhanIlgaz/passworless-auth-server/pkg/response"
	"github.com/AkifhanIlgaz/passworless-auth-server/pkg/validator"
	"github.com/labstack/echo/v4"
)

// TODO: Create a generic function to bind the request and validate it

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) LoginWithEmail(c echo.Context) error {
	var loginRequest model.LoginWithEmailRequest

	if err := c.Bind(&loginRequest); err != nil {
		c.Logger().Error(err)
		return response.WithError(c, http.StatusBadRequest, message.InvalidRequest)
	}

	if err := c.Validate(&loginRequest); err != nil {
		c.Logger().Error(err)
		return response.WithError(c, http.StatusBadRequest, message.EmailRequired)
	}

	return response.WithSuccess(c, http.StatusOK, echo.Map{})
}

func (h *AuthHandler) LoginWithPhone(c echo.Context) error {
	var loginRequest model.LoginWithPhoneRequest

	if err := c.Bind(&loginRequest); err != nil {
		c.Logger().Error(err)
		return response.WithError(c, http.StatusBadRequest, message.InvalidRequest)
	}

	if err := c.Validate(&loginRequest); err != nil {
		c.Logger().Error(err)
		return response.WithError(c, http.StatusBadRequest, message.PhoneRequired)
	}

	formattedPhoneNumber, err := validator.IsValidPhoneNumber(loginRequest.Phone)
	if err != nil {
		c.Logger().Error(err)
		return response.WithError(c, http.StatusBadRequest, message.InvalidPhoneNumber)
	}

	// TODO: Use the formatted phone number to send the OTP
	_ = formattedPhoneNumber

	return response.WithSuccess(c, http.StatusOK, echo.Map{})
}
