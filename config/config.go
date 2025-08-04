// Package config provides configuration loading and management
package config

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
	Token    TokensConfig
}

// PostgresConfig holds PostgreSQL database configuration.
type PostgresConfig struct {
	PDB_NAME     string
	PDB_PORT     string
	PDB_PASSWORD string
	PDB_USER     string
	PDB_HOST     string
}

// ServerConfig holds server-related configuration.
type ServerConfig struct {
	USER_ROUTER string
}

// TokensConfig holds JWT token configuration.
type TokensConfig struct {
	ACCES_TOKEN_KEY   string
	REFRESH_TOKEN_KEY string
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
		Token: TokensConfig{
			ACCES_TOKEN_KEY:   cast.ToString(coalesce("ACCES_TOKEN_KEY", "your_secret_key1")),
			REFRESH_TOKEN_KEY: cast.ToString(coalesce("REFRESH_TOKEN_KEY", "your_secret_key2")),
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
