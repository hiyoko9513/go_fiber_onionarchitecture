package router

import (
	"hiyoko-fiber/internal/presentation/http/app/handler"

	"github.com/gofiber/fiber/v3"
)

func NewRouter(e *fiber.App, h handler.AppHandler) {
	api := e.Group("")
	v1 := api.Group("/v1")
	v1.Post("/signup", h.Signup)

	//v1.Get("/users/:id", h.GetUser)
	//v1Guard := v1.Use(middleware.auth())
	//v1Guard.GET("/users/me", h.GetMe)
	//v1Guard.PUT("/users/:id", h.UpdateUser)
	//v1Guard.DELETE("/users/:id", h.DeleteUser)
}
