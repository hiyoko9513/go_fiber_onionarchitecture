package handler

import (
	"hiyoko-fiber/internal/application/usecase"
)

type CliHandler interface {
	GenJWTSecretKeyForApp() error
}

type cliHandler struct {
	FileUseCase usecase.FileUseCase
}

func NewCliHandler(f usecase.FileUseCase) CliHandler {
	return &cliHandler{f}
}

func (h *cliHandler) GenJWTSecretKeyForApp() error {
	err := h.FileUseCase.GenJWTSecretKey("cmd/app")
	if err != nil {
		return err
	}
	return nil
}
