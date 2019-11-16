package lib

import "os"

// GetEnv an environment variable. If no such envar is found, `fallback` is used.
func GetEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
