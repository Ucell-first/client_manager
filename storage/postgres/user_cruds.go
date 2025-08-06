package postgres

import (
	"database/sql"

	"github.com/Ucell/client_manager/storage/repo"
)

type UserRepository struct {
	Db *sql.DB
}

func NewUserRepository(db *sql.DB) repo.IUserStorage {
	return &UserRepository{Db: db}
}
