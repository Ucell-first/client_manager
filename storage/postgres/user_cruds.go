package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Ucell/client_manager/storage/repo"
)

type UserRepository struct {
	Db *sql.DB
}

func NewUserRepository(db *sql.DB) repo.IUserStorage {
	return &UserRepository{Db: db}
}

func (r *UserRepository) Create(ctx context.Context, user *repo.User) error {
	query := `
		INSERT INTO users (msisdn, name, is_active)
		VALUES ($1, $2, $3)
		RETURNING user_id
	`
	err := r.Db.QueryRowContext(ctx, query, user.MSISDN, user.Name, user.IsActive).Scan(&user.UserID)
	if err != nil {
		return fmt.Errorf("foydalanuvchini qo'shishda xatolik: %w", err)
	}
	return nil
}

func (r *UserRepository) GetAll(ctx context.Context) ([]*repo.User, error) {
	query := `
		SELECT user_id, msisdn, name, is_active
		FROM users
		ORDER BY name
	`
	rows, err := r.Db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("foydalanuvchilarni olishda xatolik: %w", err)
	}
	defer rows.Close()

	var users []*repo.User
	for rows.Next() {
		u := repo.User{}
		if err := rows.Scan(&u.UserID, &u.MSISDN, &u.Name, &u.IsActive); err != nil {
			return nil, err
		}
		users = append(users, &u)
	}
	return users, nil
}

func (r *UserRepository) GetByID(ctx context.Context, id string) (*repo.User, error) {
	query := `
		SELECT user_id, msisdn, name, is_active
		FROM users
		WHERE user_id = $1
	`
	u := repo.User{}
	err := r.Db.QueryRowContext(ctx, query, id).Scan(&u.UserID, &u.MSISDN, &u.Name, &u.IsActive)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("foydalanuvchini olishda xatolik: %w", err)
	}
	return &u, nil
}

func (r *UserRepository) Update(ctx context.Context, user *repo.User) error {
	query := `
		UPDATE users
		SET msisdn = $1, name = $2, is_active = $3
		WHERE user_id = $4
	`
	res, err := r.Db.ExecContext(ctx, query, user.MSISDN, user.Name, user.IsActive, user.UserID)
	if err != nil {
		return fmt.Errorf("foydalanuvchini yangilashda xatolik: %w", err)
	}
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("foydalanuvchi topilmadi")
	}
	return nil
}

func (r *UserRepository) Delete(ctx context.Context, id string) error {
	query := `
		DELETE FROM users
		WHERE user_id = $1
	`
	res, err := r.Db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("foydalanuvchini o'chirishda xatolik: %w", err)
	}
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("foydalanuvchi topilmadi")
	}
	return nil
}
