package main

import (
	"database/sql"
	"log"
	"os"
	"strings"

	"github.com/Diaszano/shortener/internal/config"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", config.Env.Database.Dsn())
	if err != nil {
		panic(err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		panic(err)
	}

	m, err := migrate.NewWithDatabaseInstance("file:internal/database/migrations", "postgres", driver)
	if err != nil {
		panic(err)
	}

	switch cmd := os.Args[len(os.Args)-1]; strings.ToLower(cmd) {
	case "up":
		err = m.Up()
		if err != nil {
			panic(err)
		}
	case "down":
		err = m.Down()
		if err != nil {
			panic(err)
		}
	default:
		log.Fatalf("invalid option: %s", cmd)
	}
}
