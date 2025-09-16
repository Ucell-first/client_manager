package configuration

import "fmt"

type TokenConfig struct {
	TOKEN string `json:"token"`
}

// Validate validates token configuration
func (s *TokenConfig) Validate() error {
	if s.TOKEN == "" {
		return fmt.Errorf("%w: %w", ErrConfiguration, ErrTokenNotSet)
	}
	return nil
}
