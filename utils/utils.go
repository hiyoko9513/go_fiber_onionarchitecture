package utils

import "strings"

// Contains Check if the slice contains a specific string
func Contains(lines []string, envName string) bool {
	for _, line := range lines {
		if strings.HasPrefix(line, envName) {
			return true
		}
	}
	return false
}
