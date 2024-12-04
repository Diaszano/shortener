// Package main is the entry point for the database migration tool. It utilizes
// the golang-migrate library to manage schema migrations for a PostgreSQL database.
package main

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"strings"

	"go.uber.org/zap"

	"github.com/Diaszano/shortener/internal/config"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

// main is the entry point of the migration tool. It initializes the database connection,
// sets up the migration instance, and executes migration commands based on the provided arguments.
func main() {

	logger := config.GetLogger()
	defer config.CloseLogger()

	db, err := sql.Open("postgres", config.Env.Database.Dsn())
	if err != nil {
		logger.Fatal("Failed to connect to the database", zap.Error(err))
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		logger.Fatal("Failed to initialize database driver", zap.Error(err))
	}

	m, err := migrate.NewWithDatabaseInstance("file:internal/database/migrations", "postgres", driver)
	if err != nil {
		logger.Fatal("Failed to initialize migrations", zap.Error(err))
	}

	switch cmd := os.Args[len(os.Args)-1]; strings.ToLower(cmd) {
	case "up":
		if err := m.Up(); err != nil {
			if errors.Is(err, migrate.ErrNoChange) {
				logger.Info("No migrations to apply (up-to-date).")
			} else {
				logger.Fatal("Failed to apply migrations", zap.Error(err))
			}
		} else {
			logger.Info("Migrations applied successfully.")
		}
	case "down":
		if err := m.Down(); err != nil {
			if errors.Is(err, migrate.ErrNoChange) {
				logger.Info("No migrations to revert (already clean).")
			} else {
				logger.Fatal("Failed to revert migrations", zap.Error(err))
			}
		} else {
			logger.Info("Migrations reverted successfully.")
		}
	default:
		logger.Fatal(fmt.Sprintf("Invalid option: %s", cmd))
	}
}
