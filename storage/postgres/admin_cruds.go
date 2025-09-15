package postgres

import (
	"database/sql"

	"github.com/Ucell/client_manager/storage/repo"
)

type AdminRepository struct {
	Db *sql.DB
}

func NewAdminRepository(db *sql.DB) repo.IAdminStorage {
	return &AdminRepository{Db: db}
}
