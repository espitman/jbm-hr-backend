package database

import (
	"log"

	"github.com/espitman/jbm-hr-backend/ent"
	_ "github.com/lib/pq"
)

// NewClient creates a new Ent client
func NewClient() (*ent.Client, error) {
	client, err := ent.Open("postgres", "postgresql://root:agk1Lnu4byF6ifGzEgJREJs3@cho-oyu.liara.cloud:30205/postgres?search_path=jbmhr&sslmode=disable")
	if err != nil {
		log.Printf("failed opening connection to postgres: %v", err)
		return nil, err
	}
	return client, nil
}
