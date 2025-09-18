package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Ucell/client_manager/storage/repo"
	"golang.org/x/crypto/bcrypt"
)

type AdminRepository struct {
	Db *sql.DB
}

func NewAdminRepository(db *sql.DB) repo.IAdminStorage {
	return &AdminRepository{Db: db}
}

func (r *AdminRepository) Login(ctx context.Context, login, password string) (*repo.Admin, error) {
	query := `
		SELECT id, login, hashed_password
		FROM admins
		WHERE login = $1
	`
	admin := repo.Admin{}
	var hashedPassword string

	err := r.Db.QueryRowContext(ctx, query, login).Scan(&admin.ID, &admin.Login, &hashedPassword)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("admin topilmadi")
	} else if err != nil {
		return nil, fmt.Errorf("database xatolik: %w", err)
	}

	// Parolni tekshirish
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		return nil, fmt.Errorf("parol noto'g'ri")
	}

	return &admin, nil
}

func (r *AdminRepository) GetByID(ctx context.Context, id string) (*repo.Admin, error) {
	query := `
		SELECT id, login
		FROM admins
		WHERE id = $1
	`
	admin := repo.Admin{}

	err := r.Db.QueryRowContext(ctx, query, id).Scan(&admin.ID, &admin.Login)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("database xatolik: %w", err)
	}

	return &admin, nil
}
