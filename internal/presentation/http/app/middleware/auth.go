package middleware

import (
	"hiyoko-fiber/internal/shared"
	"hiyoko-fiber/pkg/auth/v1"
	"hiyoko-fiber/pkg/logging/file"

	"github.com/gofiber/fiber/v2"
)

func CheckAuth() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		_, err := auth.GetClaimsFromCtx(c)
		if err != nil {
			logger.Error("invalid token", "error", err)
			return shared.ResponseUnauthorized(c)
		}
		return c.Next()
	}
}
