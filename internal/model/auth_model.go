package model

type LoginWithEmailRequest struct {
	Email string `json:"email" validate:"email,required"`
}

type LoginWithPhoneRequest struct {
	Phone string `json:"phone" validate:"required"`
}
