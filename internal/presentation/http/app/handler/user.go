package handler

import (
	"hiyoko-fiber/internal/application/usecase"
	"hiyoko-fiber/internal/presentation/http/app/input"
	"hiyoko-fiber/internal/shared"
	"hiyoko-fiber/pkg/logging/file"

	"github.com/gofiber/fiber/v3"
)

type UserHandler interface {
	Signup(c fiber.Ctx) error

	//GetUser(c fiber.Ctx) error
	//UpdateUser(c fiber.Ctx) error
}

type userHandler struct {
	UserUseCase usecase.UserUseCase
}

func NewUserHandler(u usecase.UserUseCase) UserHandler {
	return &userHandler{u}
}

//func (h *userHandler) GetUser(c fiber.Ctx) error {
//	ctx := c.Context()
//	id := c.Params("id")
//
//	user, err := h.UserUseCase.GetUser(ctx, id)
//	if err != nil {
//		return shared.ResponseNotFound(c, shared.NoneCode)
//	}
//
//	return shared.ResponseOK(c, user)
//}

func (h *userHandler) Signup(c fiber.Ctx) error {
	ctx := c.Context()
	userInput := &input.UserCreateInput{
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

	return shared.ResponseCreate(c, user)
}

//func (h *userHandler) UpdateUser(c fiber.Ctx) error {
//	ctx := c.Context()
//	id := c.FormValue("id")
//
//	user, err := h.UserUseCase.UpdateUser(ctx, id)
//	if err != nil {
//		return shared.ResponseBadRequest(c, shared.NoneCode)
//	}
//
//	return shared.ResponseOK(c, user)
//}
