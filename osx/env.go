package osx

import (
	"fmt"
	"os"
)

// GetenvDefault returns an environment variable or the default value if it is empty.
func GetenvDefault(key string, def string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return def
}

// MustMapEnv maps the provided environment varible to given target pointer.
func MustMapEnv(target *string, key string) {
	var value string
	if value = os.Getenv(key); value == "" {
		fmt.Fprintf(os.Stderr, "environment variable %q not set", key)
	}
	*target = value
}
