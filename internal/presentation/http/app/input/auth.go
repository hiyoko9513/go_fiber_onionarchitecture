package input

import "github.com/go-playground/validator/v10"

type SignupInput struct {
	ID       string `json:"id"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (i SignupInput) Validate() error {
	return validator.New().Struct(i)
}