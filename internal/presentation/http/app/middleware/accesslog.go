package middleware

import (
	"time"

	"hiyoko-fiber/pkg/logging/file"

	"github.com/gofiber/fiber/v2"
)

func accessLogger() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()

		err := c.Next()

		latency := time.Now().Sub(start)
		requestID := c.Locals(RequestIDContextKey)
		logger.Access(
			"request id", requestID,
			"ip", c.IP(),
			"method", c.Method(),
			"path", c.Path(),
			"status", c.Response().StatusCode(),
			"latency", latency,
			"error", err,
		)

		return err
	}
}
