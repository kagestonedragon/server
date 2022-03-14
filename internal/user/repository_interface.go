package user

import "context"

type Repository interface {
	Add(ctx context.Context, user User) (User, error)
	Update(ctx context.Context, user User) error
	DeleteById(ctx context.Context, id uint64) error
	GetById(ctx context.Context, id uint64) (User, error)
	GetList(ctx context.Context) ([]User, error)

	cache
}
