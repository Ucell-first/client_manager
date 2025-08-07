package configuration

import "fmt"

// ServerConfig holds server-related configuration
type ServerConfig struct {
	Port string `json:"port"`
}

// Validate validates server configuration
func (s *ServerConfig) Validate() error {
	if s.Port == "" {
		return fmt.Errorf("%w: %w: %w", ErrConfiguration, ErrServer, ErrServerPortNotSet)
	}

	// Port formatini tekshirish (:1234 formatida bo'lishi kerak)
	if s.Port[0] != ':' {
		return fmt.Errorf("%w: %w: port should start with ':'", ErrConfiguration, ErrServer)
	}

	return nil
}

// GetAddress returns server address
func (s *ServerConfig) GetAddress() string {
	return s.Port
}
