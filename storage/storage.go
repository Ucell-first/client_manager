package storage

import (
	"database/sql"

	"github.com/Ucell/client_manager/storage/postgres"
	"github.com/Ucell/client_manager/storage/repo"
)

type IStorage interface {
	User() repo.IUserStorage
	Admin() repo.IAdminStorage
	ClosePDB() error
}

type databaseStorage struct {
	pdb *sql.DB
}

func NewStorage(pdb *sql.DB) IStorage {
	return &databaseStorage{
		pdb: pdb,
	}
}

func (p *databaseStorage) ClosePDB() error {
	err := p.pdb.Close()
	if err != nil {
		return err
	}
	return nil
}

func (p *databaseStorage) User() repo.IUserStorage {
	return postgres.NewUserRepository(p.pdb)
}

func (p *databaseStorage) Admin() repo.IAdminStorage { return postgres.NewAdminRepository(p.pdb) }
