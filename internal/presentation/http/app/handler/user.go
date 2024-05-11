package handler

import (
	"hiyoko-fiber/internal/application/usecase"
	"hiyoko-fiber/internal/pkg/auth/v1"
	"hiyoko-fiber/internal/pkg/ent/util"
	"hiyoko-fiber/internal/presentation/http/app/oapi"
	"hiyoko-fiber/internal/shared"

	"github.com/gofiber/fiber/v2"
)

type UserHandler interface {
	GetMe(c *fiber.Ctx) error
}

type userHandler struct {
	UserUseCase usecase.UserUseCase
}

func NewUserHandler(u usecase.UserUseCase) UserHandler {
	return &userHandler{u}
}

func (h *userHandler) GetMe(c *fiber.Ctx) error {
	ctx := c.Context()
	claims, err := auth.GetClaimsFromCtx(c)
	if err != nil {
		return shared.ResponseUnauthorized(c)
	}
	user, err := h.UserUseCase.GetUser(ctx, util.ULID(claims.ID))
	if err != nil {
		return shared.ResponseNotFound(c, shared.NoneCode)
	}

	return shared.ResponseOK(c, oapi.MeResponse{
		ID:         user.ID,
		OriginalID: &user.OriginalID,
		Email:      user.Email,
	})
}
