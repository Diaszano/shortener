// Package config provides configuration structures and parsing logic
// for the application, utilizing environment variables with default
// values where applicable.
package config

import (
	"fmt"
	"strings"

	"github.com/caarlos0/env/v11"
	_ "github.com/joho/godotenv/autoload"
	"go.uber.org/zap/zapcore"
)

// cache holds the Redis cache configuration, loaded from environment
// variables prefixed with "CACHE_".
type cache struct {
	// Host specifies the Redis server's hostname or IP address (required).
	Host string `env:"HOST,required,notEmpty"`

	// Port specifies the Redis server's port (default: 6379).
	Port uint `env:"PORT" envDefault:"6379"`

	// User specifies the username for Redis authentication.
	User string `env:"USER"`

	// Pass specifies the password for Redis authentication.
	Pass string `env:"PASS"`

	// DB specifies the database index to use (default: 0).
	DB uint `env:"DB" envDefault:"0"`

	// Protocol defines the Redis protocol version (default: 3).
	Protocol uint `env:"PROTOCOL" envDefault:"3"`
}

// Addr generates a connection address combining the Host and Port fields.
func (c cache) Addr() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}

// database holds the database configuration, loaded from environment
// variables prefixed with "DATABASE_".
type database struct {
	// Host specifies the database server's hostname or IP address (required).
	Host string `env:"HOST,required,notEmpty"`

	// Port specifies the database server's port (default: 5432).
	Port uint `env:"PORT" envDefault:"5432"`

	// User specifies the username for database authentication (required).
	User string `env:"USER,required,notEmpty"`

	// Pass specifies the password for database authentication (required).
	Pass string `env:"PASS,required,notEmpty"`

	// Name specifies the database name to connect to (required).
	Name string `env:"NAME,required,notEmpty"`

	// SslMode specifies the SSL mode for the connection (default: "disable").
	SslMode string `env:"SSLMODE" envDefault:"disable"`
}

// Dsn builds a Data Source Name (DSN) string for the database connection.
func (d database) Dsn() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		d.Host, d.Port, d.User, d.Pass, d.Name, d.SslMode,
	)
}

// server holds the server configuration, loaded from environment
// variables prefixed with "SERVER_".
type server struct {
	// LogLevel specifies the logging verbosity (default: "info").
	LogLevel string `env:"LOG_LEVEL" envDefault:"info"`

	// AppEnv specifies the application environment (default: "production").
	AppEnv string `env:"APP_ENV" envDefault:"production"`

	// Version specifies the application version (default: "0.1.0").
	Version string `env:"VERSION" envDefault:"0.1.0"`
}

// IsDevelopment checks if the application is running in a development environment.
func (s server) IsDevelopment() bool {
	appEnv := strings.ToLower(s.AppEnv)
	return appEnv == "development" || appEnv == "dev"
}

// GetLogLevel parses and returns the configured log level as a zapcore.Level.
// Defaults to zapcore.InfoLevel if parsing fails.
func (s server) GetLogLevel() zapcore.Level {
	logLevel, err := zapcore.ParseLevel(strings.ToLower(s.LogLevel))
	if err != nil {
		return zapcore.InfoLevel
	}
	return logLevel
}

// environment aggregates all configuration settings grouped under
// respective prefixes: DATABASE_, SERVER_, and CACHE_.
type environment struct {
	Database database `envPrefix:"DATABASE_"`
	Server   server   `envPrefix:"SERVER_"`
	Cache    cache    `envPrefix:"CACHE_"`
}

// Env holds the global instance of the application's configuration,
// initialized during package setup.
var Env environment

// init parses environment variables into the Env structure,
// panicking if parsing fails.
func init() {
	if err := env.Parse(&Env); err != nil {
		panic(err)
	}
}
