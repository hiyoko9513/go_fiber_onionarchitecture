package util

import (
	"fmt"
	"os"
	"strconv"

	"hiyoko-fiver/pkg/logging/file"

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

func (e Env) GetString(defaultVal string) string {
	value := os.Getenv(string(e))
	if value == "" {
		return defaultVal
	}
	return value
}

func (e Env) GetInt(defaultVal int) int {
	valString := e.GetString("")
	if valString == "" {
		return defaultVal
	}
	val, err := strconv.Atoi(valString)
	if err != nil {
		return defaultVal
	}
	return val
}
