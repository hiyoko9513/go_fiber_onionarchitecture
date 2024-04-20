package utils

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"hiyoko-fiber/pkg/logging/file"

	"github.com/joho/godotenv"
)

func LoadEnv(rootPath string) {
	envPath := fmt.Sprintf("%s/.env", rootPath)
	err := godotenv.Load(envPath)
	if err != nil {
		logger.Fatal("failed to load environment", "error", err)
	}
}

type Env string

func (e Env) GetString(defaultVals ...string) string {
	value := os.Getenv(string(e))
	if value == "" {
		if len(defaultVals) > 0 {
			return defaultVals[0]
		}
		return ""
	}
	return value
}

func (e Env) GetInt(defaultVals ...int) int {
	valString := os.Getenv(string(e))
	value, err := strconv.Atoi(valString)

	if err != nil {
		if len(defaultVals) > 0 {
			return defaultVals[0]
		}
		return 0
	}
	return value
}

func (e Env) GetDuration(defaultVals ...time.Duration) time.Duration {
	valString := os.Getenv(string(e))
	value, err := time.ParseDuration(valString)

	if err != nil {
		if len(defaultVals) > 0 {
			return defaultVals[0]
		}
		return 0
	}
	return value
}
