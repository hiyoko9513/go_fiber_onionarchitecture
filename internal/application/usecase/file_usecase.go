package usecase

import (
	"path/filepath"

	"hiyoko-fiber/internal/pkg/auth/v1"
	"hiyoko-fiber/utils"
)

type FileUseCase interface {
	GenJWTSecretKey(envPath string) error
}

type fileUseCase struct {
}

func NewFileUseCase() FileUseCase {
	return &fileUseCase{}
}

func (h *fileUseCase) GenJWTSecretKey(envPath string) error {
	secretKey, err := auth.GenerateRandomBase64String(255)
	if err != nil {
		return err
	}
	return utils.EnvFile(filepath.Join(envPath, ".env")).RegisterVariable("JWT_SECRET_KEY", secretKey)
}
