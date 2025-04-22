package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"gin-project/ent"

	_ "github.com/lib/pq"
)

var Client *ent.Client

// InitDB initializes the database connection using Ent
func InitDB() (*ent.Client, error) {
	// Get database connection details from environment variables
	host := getEnv("DB_HOST", "localhost")
	port := getEnv("DB_PORT", "5432")
	user := getEnv("DB_USER", "postgres")
	password := getEnv("DB_PASSWORD", "postgres")
	dbname := getEnv("DB_NAME", "hr_db")

	// Create connection string
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Create Ent client
	client, err := ent.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed opening connection to postgres: %v", err)
	}

	// Run the auto migration tool
	if err := client.Schema.Create(context.Background()); err != nil {
		return nil, fmt.Errorf("failed creating schema resources: %v", err)
	}

	Client = client
	log.Println("Successfully connected to database using Ent")
	return client, nil
}

// CloseDB closes the database connection
func CloseDB() {
	if Client != nil {
		if err := Client.Close(); err != nil {
			log.Printf("Error closing database connection: %v", err)
		}
	}
}

// getEnv gets an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
