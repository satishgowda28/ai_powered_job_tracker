package database

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/satishgowda28/ai_powered_job_tracker/internal/config"
)

var DB *pgxpool.Pool

func Connect() {
	cfg := config.Get()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err := pgxpool.New(ctx, cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Unable to create db pool %v", err)
	}

	if err := pool.Ping(ctx); err != nil {
		log.Fatalf("Unable to ping db %v", err)
	}

	DB = pool
	log.Println("conneted to postgresSQL")
}
