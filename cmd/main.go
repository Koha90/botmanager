package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"botmanager/internal/config"
	"botmanager/pkg/logger"
	"botmanager/pkg/migrator"
)

func main() {
	_ = godotenv.Load()
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbDatabase := os.Getenv("DB_DATABASE")

	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		dbUser,
		dbPass,
		dbHost,
		dbPort,
		dbDatabase,
	)

	if err := migrator.MigratePostgres(dsn, "./migrations"); err != nil {
		log.Fatalf("Failed to migrate: %v", err)
	}

	cfg := config.MustLoad()

	logger, _ := logger.Setup(cfg.Env)
	logger.Debug("debug mode is enabled")

	log.Println("All migrations applied successfully!")
}
