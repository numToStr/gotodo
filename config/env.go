package config

import (
	"fmt"
	"os"
)

var (
	// PORT returns the server listening port
	PORT = getEnv("PORT")
)

func getEnv(name string) string {
	if value, exists := os.LookupEnv(name); exists {
		return value
	}

	panic(fmt.Sprintf(`Environment variable not found :: %v`, name))
}
