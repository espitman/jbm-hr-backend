package database

import (
	"fmt"
	"log"
	"os"

	"github.com/espitman/jbm-hr-backend/ent"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// NewClient creates a new Ent client
func NewClient() (*ent.Client, error) {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found: %v", err)
	}

	// Get database configuration from environment variables
	dbHost := getEnvOrDefault("DB_HOST", "localhost")
	dbPort := getEnvOrDefault("DB_PORT", "5432")
	dbUser := getEnvOrDefault("DB_USER", "postgres")
	dbPassword := getEnvOrDefault("DB_PASSWORD", "postgres")
	dbName := getEnvOrDefault("DB_NAME", "jbm_hr")
	dbSSLMode := getEnvOrDefault("DB_SSL_MODE", "disable")

	// Construct connection string
	connStr := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s",
		dbUser, dbPassword, dbHost, dbPort, dbName, dbSSLMode)

	client, err := ent.Open("postgres", connStr)
	if err != nil {
		log.Printf("failed opening connection to postgres: %v", err)
		return nil, err
	}
	return client, nil
}

// getEnvOrDefault returns the value of the environment variable or a default value if not set
func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
