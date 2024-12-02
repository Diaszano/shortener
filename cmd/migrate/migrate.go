// Package main is the entry point for the database migration tool. It utilizes
// the golang-migrate library to manage schema migrations for a PostgreSQL database.
package main

import (
	"database/sql"
	"errors"
	"log"
	"os"
	"strings"

	"github.com/Diaszano/shortener/internal/config"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

// main is the entry point of the migration tool. It initializes the database connection,
// sets up the migration instance, and executes migration commands based on the provided arguments.
func main() {
	db, err := sql.Open("postgres", config.Env.Database.Dsn())
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalf("Failed to initialize database driver: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance("file:internal/database/migrations", "postgres", driver)
	if err != nil {
		log.Fatalf("Failed to initialize migrations: %v", err)
	}

	switch cmd := os.Args[len(os.Args)-1]; strings.ToLower(cmd) {
	case "up":
		if err := m.Up(); err != nil {
			if errors.Is(err, migrate.ErrNoChange) {
				log.Println("No migrations to apply (up-to-date).")
			} else {
				log.Fatalf("Failed to apply migrations: %v", err)
			}
		} else {
			log.Println("Migrations applied successfully.")
		}
	case "down":
		if err := m.Down(); err != nil {
			if errors.Is(err, migrate.ErrNoChange) {
				log.Println("No migrations to revert (already clean).")
			} else {
				log.Fatalf("Failed to revert migrations: %v", err)
			}
		} else {
			log.Println("Migrations reverted successfully.")
		}
	default:
		log.Fatalf("Invalid option: %s", cmd)
	}
}
