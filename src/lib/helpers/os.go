package helpers

import "os"

// GetEnvOrDefault will get the environment variable from key.
// If it cannot find the variable, it will fallback to a
// provided default value.
func GetEnvOrDefault(key string, fallback string) string {
	value := os.Getenv(key)

	if len(value) == 0 {
		return fallback
	}

	return value
}
