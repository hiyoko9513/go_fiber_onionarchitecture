package handler

import (
	"hiyoko-fiber/internal/application/usecase"
	"hiyoko-fiber/internal/presentation/http/app/input"
	"hiyoko-fiber/internal/presentation/http/app/oapi"
	"hiyoko-fiber/internal/shared"
	"hiyoko-fiber/pkg/logging/file"

	"github.com/gofiber/fiber/v3"
)

type AuthHandler interface {
	Signup(c fiber.Ctx) error

	//GetUser(c fiber.Ctx) error
	//UpdateUser(c fiber.Ctx) error
}

type authHandler struct {
	UserUseCase usecase.UserUseCase
}

func NewAuthHandler(u usecase.UserUseCase) AuthHandler {
	return &authHandler{u}
}
func (h *authHandler) Signup(c fiber.Ctx) error {
	ctx := c.Context()
	userInput := &input.SignupInput{
		ID:       c.FormValue("id"),
		Email:    c.FormValue("email"),
		Password: c.FormValue("password"),
	}
	if err := userInput.Validate(); err != nil {
		return shared.ResponseBadRequest(c, shared.NoneCode)
	}

	user, err := h.UserUseCase.Signup(ctx, userInput)
	if err != nil {
		logger.Error("Error create user", "userInput", userInput, "error", err)
		return shared.ResponseBadRequest(c, shared.NoneCode)
	}

	return shared.ResponseCreate(c, oapi.SignupResponse{
		Token: "test",
		User: oapi.UserSchema{
			Id:    user.ID,
			Email: user.Email,
			Sub:   user.Sub,
		},
	})
}
