package configuration

import (
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
func Load() *Config {
	if err := godotenv.Load(".env"); err != nil {
		log.Printf("error while loading .env file: %v", err)
	}

	return &Config{
		Postgres: PostgresConfig{
			PDB_HOST:     cast.ToString(coalesce("PDB_HOST", "localhost")),
			PDB_PORT:     cast.ToString(coalesce("PDB_PORT", "5432")),
			PDB_USER:     cast.ToString(coalesce("PDB_USER", "postgres")),
			PDB_NAME:     cast.ToString(coalesce("PDB_NAME", "postgres")),
			PDB_PASSWORD: cast.ToString(coalesce("PDB_PASSWORD", "3333")),
		},
		Server: ServerConfig{
			USER_ROUTER: cast.ToString(coalesce("USER_ROUTER", ":1234")),
		},
	}
}

func coalesce(key string, value interface{}) interface{} {
	val, exist := os.LookupEnv(key)
	if exist {
		return val
	}
	return value
}
