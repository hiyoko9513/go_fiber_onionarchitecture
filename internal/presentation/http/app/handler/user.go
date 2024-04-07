package handler

import (
	"hiyoko-fiber/internal/application/usecase"
	"hiyoko-fiber/internal/domain/entities/users"
	"hiyoko-fiber/internal/shared"

	"github.com/gofiber/fiber/v3"
)

type UserHandler interface {
	GetMe(c fiber.Ctx) error
}

type userHandler struct {
	UserUseCase usecase.UserUseCase
}

func NewUserHandler(u usecase.UserUseCase) UserHandler {
	return &userHandler{u}
}

func (h *userHandler) GetMe(c fiber.Ctx) error {
	return shared.ResponseCreate(c, &users.UserEntity{})
}
