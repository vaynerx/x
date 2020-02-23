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

// MustMapEnv maps the provided environment varible to a key in the provided struct.
func MustMapEnv(target *string, def string) {
	var value string
	if value = os.Getenv(def); value == "" {
		fmt.Print(fmt.Sprintf("environment variable %q not set", def))
	}
	*target = value
}
