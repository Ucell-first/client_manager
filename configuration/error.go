package configuration

import "errors"

var (
	ErrConfiguration    = errors.New("Configuration")
	ErrValidate         = errors.New("vallidate method")
	ErrDb               = errors.New("database")
	ErrServer           = errors.New("server")
	ErrDBHostNotSet     = errors.New("database host is not set")
	ErrDBPortNotSet     = errors.New("database port is not set")
	ErrDBUserNotSet     = errors.New("database user is not set")
	ErrDBNameNotSet     = errors.New("database name is not set")
	ErrDBPasswordNotSet = errors.New("database password is not set")
	ErrServerPortNotSet = errors.New("server port is not set")
)
