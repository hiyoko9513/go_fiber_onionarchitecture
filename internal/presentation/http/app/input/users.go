package input

import "github.com/go-playground/validator/v10"

// todo 独自バリデーションを作成する
type UserCreateInput struct {
	ID       string `json:"id"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (i UserCreateInput) Validate() error {
	return validator.New().Struct(i)
}
