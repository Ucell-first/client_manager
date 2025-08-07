package postgres

import (
	"database/sql"
	"fmt"

	"github.com/Ucell/client_manager/configuration"
	_ "github.com/lib/pq"
)

func ConnectPdb(conf *configuration.PostgresConfig) (*sql.DB, error) {
	conDb := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		conf.Host,
		conf.Port,
		conf.User,
		conf.Name,
		conf.Password,
	)
	db, err := sql.Open("postgres", conDb)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
