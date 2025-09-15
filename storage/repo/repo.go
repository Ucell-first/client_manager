package repo

import "context"

type User struct {
	UserID   string
	MSISDN   string
	Name     string
	IsActive bool
}

type IUserStorage interface {
	Create(ctx context.Context, user *User) error
	GetAll(ctx context.Context) ([]*User, error)
	GetByID(ctx context.Context, id string) (*User, error)
	Update(ctx context.Context, user *User) error
	Delete(ctx context.Context, id string) error
}

type IAdminStorage interface {
}
