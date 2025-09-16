package postgres

import (
	"context"
	"database/sql"

	"github.com/Ucell/client_manager/storage/repo"
)

type AdminRepository struct {
	Db *sql.DB
}

func NewAdminRepository(db *sql.DB) repo.IAdminStorage {
	return &AdminRepository{Db: db}
}

func (AdminRepository) Login(ctx context.Context, user *repo.User) error {
	return nil
}
