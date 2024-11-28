// Package config provides configuration structures and parsing logic
// for the application. It uses environment variables to load the
// configurations and applies default values when necessary.
package config

import (
	"fmt"
	"github.com/caarlos0/env/v11"
	_ "github.com/joho/godotenv/autoload"
)

// database represents the database configuration, which is loaded from
// environment variables prefixed with "DATABASE_".
type database struct {
	// Host specifies the hostname or IP address of the database server.
	// This field is required and cannot be empty.
	Host string `env:"HOST,required,notEmpty"`

	// Port specifies the port to connect to on the database server.
	// Defaults to "5432" if not explicitly set.
	Port uint `env:"PORT" envDefault:"5432"`

	// User specifies the username for the database connection.
	// This field is required and cannot be empty.
	User string `env:"USER,required,notEmpty"`

	// Pass specifies the password for the database connection.
	// This field is required and cannot be empty.
	Pass string `env:"PASS,required,notEmpty"`

	// Name specifies the name of the database to connect to.
	// This field is required and cannot be empty.
	Name string `env:"NAME,required,notEmpty"`

	// SslMode specifies the SSL mode for the database connection.
	// Defaults to "disable" if not explicitly set.
	SslMode string `env:"SSLMODE" envDefault:"disable"`
}

// Dsn generates a Data Source Name (DSN) string for connecting to the database.
// The DSN is formatted using the configuration fields.
func (d database) Dsn() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		d.Host, d.Port, d.User, d.Pass, d.Name, d.SslMode,
	)
}

// server represents the server configuration, which is loaded from
// environment variables prefixed with "SERVER_".
type server struct {
	// Port specifies the port the server will listen on for incoming connections.
	// Defaults to "2821" if not explicitly set.
	Port uint `env:"PORT" envDefault:"2821"`

	// LogLevel specifies the verbosity level of the application logs.
	// Defaults to "info" if not explicitly set.
	LogLevel string `env:"LOG_LEVEL" envDefault:"info"`
}

// environment aggregates all configurations for the application,
// grouping them under respective prefixes for database and server settings.
type environment struct {
	// Database holds the database configuration, with all fields
	// prefixed by "DATABASE_".
	Database database `envPrefix:"DATABASE_"`

	// Server holds the server configuration, with all fields
	// prefixed by "SERVER_".
	Server server `envPrefix:"SERVER_"`
}

// Env is the global instance of the environment configuration,
// populated during the initialization of the package.
var Env environment

// init initializes the package by parsing environment variables
// into the Env variable. If parsing fails, the application panics.
func init() {
	err := env.Parse(&Env)
	if err != nil {
		panic(err)
	}
}
