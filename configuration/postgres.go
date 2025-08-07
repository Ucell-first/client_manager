package configuration

import (
	"fmt"
	"strconv"
)

type PostgresConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

// Validate validates PostgreSQL configuration
func (p *PostgresConfig) Validate() error {
	if p.Host == "" {
		return fmt.Errorf("%w: %w: %w", ErrConfiguration, ErrDB, ErrDBHostNotSet)
	}

	if p.Port == "" {
		return fmt.Errorf("%w: %w: %w", ErrConfiguration, ErrDB, ErrDBPortNotSet)
	}

	// Port raqam ekanligini tekshirish
	if _, err := strconv.Atoi(p.Port); err != nil {
		return fmt.Errorf("%w: %w: invalid port number: %s", ErrConfiguration, ErrDB, p.Port)
	}

	if p.User == "" {
		return fmt.Errorf("%w: %w: %w", ErrConfiguration, ErrDB, ErrDBUserNotSet)
	}

	if p.Name == "" {
		return fmt.Errorf("%w: %w: %w", ErrConfiguration, ErrDB, ErrDBNameNotSet)
	}

	if p.Password == "" {
		return fmt.Errorf("%w: %w: %w", ErrConfiguration, ErrDB, ErrDBPasswordNotSet)
	}

	return nil
}

func (p *PostgresConfig) GetDSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		p.Host, p.Port, p.User, p.Name, p.Password)
}
