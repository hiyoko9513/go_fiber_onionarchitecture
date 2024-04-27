package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

type EnvFile string

// LoadEnv load env variables
func (e EnvFile) LoadEnv() error {
	err := godotenv.Load(string(e))
	if err != nil {
		return err
	}
	return nil
}

// RegisterVariable registers variables in the env file
func (e EnvFile) RegisterVariable(key string, value string) error {
	file, err := os.Open(string(e))
	if err != nil {
		return err
	}
	defer file.Close()

	lines, err := readAndUpdateLines(file, key, value)
	if err != nil {
		return err
	}

	if !Contains(lines, key) {
		lines = append(lines, fmt.Sprintf("%s=%s", key, value))
	}

	return e.writeLines(lines)
}

// writeLines writes each line to the .env file
func (e EnvFile) writeLines(lines []string) error {
	file, err := os.Create(string(e))
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, line := range lines {
		_, err = fmt.Fprintln(writer, line)
		if err != nil {
			return err
		}
	}

	return writer.Flush()
}

// readAndUpdateLines reads .env file, updates existing keys if necessary, and returns all lines
func readAndUpdateLines(file *os.File, key string, value string) ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, key) {
			parts := strings.SplitN(line, "=", 2)
			if len(parts) > 1 && parts[1] == "" {
				line = fmt.Sprintf("%s=%s", key, value)
			}
		}

		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
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

func (e Env) GetBool(defaultVals ...bool) bool {
	valString := os.Getenv(string(e))
	value, err := strconv.ParseBool(valString)

	if err != nil {
		if len(defaultVals) > 0 {
			return defaultVals[0]
		}
		return false
	}
	return value
}
