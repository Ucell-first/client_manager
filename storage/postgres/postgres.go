package postgres

import (
	"database/sql"

	"github.com/Ucell/client_manager/configuration"
	_ "github.com/lib/pq"
)

func ConnectPdb(conf *configuration.PostgresConfig) (*sql.DB, error) {
	db, err := sql.Open("postgres", conf.GetDSN())
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
