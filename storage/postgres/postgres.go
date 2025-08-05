package postgres

import (
	"database/sql"
	"fmt"

	"github.com/Ucell/client_manager/config"
	_ "github.com/lib/pq"
)

func ConnectPdb(conf *config.Config) (*sql.DB, error) {
	conDb := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		conf.Postgres.PDB_HOST,
		conf.Postgres.PDB_PORT,
		conf.Postgres.PDB_USER,
		conf.Postgres.PDB_NAME,
		conf.Postgres.PDB_PASSWORD,
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
