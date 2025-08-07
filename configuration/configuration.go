package configuration

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

// Config represents the main application configuration.
type Config struct {
	Postgres PostgresConfig
	Server   ServerConfig
}

// Load loads configuration from environment variables.
func Load() (*Config, error) {
	if err := godotenv.Load(".env"); err != nil {
		log.Printf("error while loading .env file: %v", err)
	}

	config := &Config{
		Postgres: PostgresConfig{
			Host:     cast.ToString(coalesce("PDB_HOST", "localhost")),
			Port:     cast.ToString(coalesce("PDB_PORT", "5432")),
			User:     cast.ToString(coalesce("PDB_USER", "postgres")),
			Name:     cast.ToString(coalesce("PDB_NAME", "postgres")),
			Password: cast.ToString(coalesce("PDB_PASSWORD", "3333")),
		},
		Server: ServerConfig{
			Port: cast.ToString(coalesce("SERVER_PORT", ":1234")),
		},
	}

	// Konfiguratsiyani validatsiya qilish
	if err := config.Validate(); err != nil {
		return nil, fmt.Errorf("failed to validate configuration: %w", err)
	}

	return config, nil
}

// Validate validates the entire configuration
func (c *Config) Validate() error {
	// PostgreSQL konfiguratsiyasini tekshirish
	if err := c.Postgres.Validate(); err != nil {
		return fmt.Errorf("%w: %w: %w", ErrConfiguration, ErrDB, err)
	}

	// Server konfiguratsiyasini tekshirish
	if err := c.Server.Validate(); err != nil {
		return fmt.Errorf("%w: %w: %w", ErrConfiguration, ErrServer, err)
	}

	return nil
}

func coalesce(key string, value interface{}) interface{} {
	val, exist := os.LookupEnv(key)
	if exist {
		return val
	}
	return value
}
