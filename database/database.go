package database

import (
	"fmt"
	"log"

	"github.com/espitman/jbm-hr-backend/ent"
	"github.com/espitman/jbm-hr-backend/utils/config"
	_ "github.com/lib/pq"
)

// NewClient creates a new Ent client
func NewClient() (*ent.Client, error) {
	// Load environment variables
	config.LoadEnv()

	// Construct connection string
	connStr := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?search_path=%s&sslmode=%s",
		config.GetConfig("DB_USER", "postgres"),
		config.GetConfig("DB_PASSWORD", "postgres"),
		config.GetConfig("DB_HOST", "localhost"),
		config.GetConfig("DB_PORT", "5432"),
		config.GetConfig("DB_NAME", "jbm_hr"),
		config.GetConfig("DB_SEARCH_PATH", "jbmhr"),
		config.GetConfig("DB_SSL_MODE", "disable"))

	client, err := ent.Open("postgres", connStr)
	if err != nil {
		log.Printf("failed opening connection to postgres: %v", err)
		return nil, err
	}
	return client, nil
}
