package input

import (
	"github.com/go-playground/validator/v10"
)

type SignupInput struct {
	OriginalID string `json:"originalID"`
	Email      string `json:"email" validate:"required"`
	Password   string `json:"password" validate:"required"`
}

func (i SignupInput) Validate() error {
	return validator.New().Struct(i)
}

type SigninInput struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (i SigninInput) Validate() error {
	return validator.New().Struct(i)
}
