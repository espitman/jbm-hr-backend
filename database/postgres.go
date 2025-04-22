package database

import (
	"context"
	"fmt"
	"log"

	"github.com/espitman/jbm-hr-backend/ent"

	_ "github.com/lib/pq"
)

var Client *ent.Client

// InitDB initializes the database connection using Ent
func InitDB() (*ent.Client, error) {
	// Connection string for Liara PostgreSQL with SSL disabled
	connStr := "postgresql://root:agk1Lnu4byF6ifGzEgJREJs3@cho-oyu.liara.cloud:30205/postgres?search_path=jbmhr&sslmode=disable"

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
