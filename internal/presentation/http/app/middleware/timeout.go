package middleware

import (
	"time"

	"hiyoko-fiber/internal/shared"

	"github.com/gofiber/fiber/v2"
)

func timeoutMiddleware(timeout time.Duration) fiber.Handler {
	return func(c *fiber.Ctx) error {
		done := make(chan bool)
		go func() {
			defer close(done)
			c.Next()
		}()

		select {
		case <-done:
			return nil
		case <-time.After(timeout):
			return shared.ResponseRequestTimeout(c, shared.NoneCode)
		}
	}
}
