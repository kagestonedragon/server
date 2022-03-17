package user

import "context"

type Repository interface {
	Add(ctx context.Context, u *User) error
	Update(ctx context.Context, u *User) error
	DeleteById(ctx context.Context, id uint64) error
	GetById(ctx context.Context, id uint64) (*User, error)
	GetList(ctx context.Context) ([]*User, error)
}

type CacheableRepository interface {
	Repository

	Reset(ctx context.Context) error
}
