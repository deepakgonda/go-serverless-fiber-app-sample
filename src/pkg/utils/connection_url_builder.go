package utils

import (
	"fmt"
	"os"

	"github.com/deepakgonda/go-serverless-test/src/pkg/logger"
)

// ConnectionURLBuilder func for building URL connection.
func ConnectionURLBuilder(n string) (string, error) {
	// Define URL to connection.
	var url string

	// Switch given names.
	switch n {
	case "postgres":
		// URL for PostgreSQL connection.
		url = fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASS"),
			os.Getenv("DB_NAME"),
		)
	case "redis":
		// URL for Redis connection.
		url = fmt.Sprintf(
			"%s:%s",
			os.Getenv("REDIS_HOST"),
			os.Getenv("REDIS_PORT"),
		)
	default:
		// Return error message.
		return "", fmt.Errorf("connection name '%v' is not supported", n)
	}

	logger.Info.Printf("Connection String Generated for %v connection: %v", n, url)
	// Return connection URL.
	return url, nil
}
