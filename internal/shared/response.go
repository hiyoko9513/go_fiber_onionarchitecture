package shared

import (
	"net/http"

	"hiyoko-fiber/pkg/logging/file"

	"github.com/gofiber/fiber/v3"
)

func ResponseOK(c fiber.Ctx, data interface{}) error {
	if err := c.SendStatus(http.StatusOK); err != nil {
		logger.Error("OK response Set status failed", "error", err, "data", data)
		return err
	}
	if err := c.JSON(data); err != nil {
		logger.Error("OK response JSON conversion failed", "error", err, "data", data)
		return err
	}
	return nil
}

func ResponseCreate(c fiber.Ctx, data interface{}) error {
	if err := c.SendStatus(http.StatusCreated); err != nil {
		logger.Error("Created response JSON conversion failed", "error", err, "data", data)
		return err
	}
	if err := c.JSON(data); err != nil {
		logger.Error("Created response JSON conversion failed", "error", err, "data", data)
		return err
	}
	return nil
}

func ResponseNoContent(c fiber.Ctx) error {
	if err := c.SendStatus(http.StatusNoContent); err != nil {
		logger.Error("NoContent response JSON conversion failed", "error", err)
		return err
	}
	return nil
}

func ResponseBadRequest(c fiber.Ctx, code string) error {
	if err := c.SendStatus(http.StatusBadRequest); err != nil {
		logger.Error("BadRequest response Set status failed", "error", err, "code", code)
		return err
	}
	if err := c.JSON(ErrorResponse{
		Code:    code,
		Message: GetErrorMessage(http.StatusBadRequest),
	}); err != nil {
		logger.Error("BadRequest response JSON conversion failed", "error", err, "code", code)
		return err
	}
	return nil
}

func ResponseNotFound(c fiber.Ctx, code string) error {
	if err := c.SendStatus(http.StatusNotFound); err != nil {
		logger.Error("NotFound response Set status failed", "error", err, "code", code)
		return err
	}
	if err := c.JSON(ErrorResponse{
		Code:    code,
		Message: GetErrorMessage(http.StatusNotFound),
	}); err != nil {
		logger.Error("NotFound response JSON conversion failed", "error", err, "code", code)
		return err
	}
	return nil
}
