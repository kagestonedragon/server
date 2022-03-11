package user

import "context"

type Repository interface {
	Add(ctx context.Context, user *User) error
	GetById(ctx context.Context, id uint64) (*User, error)
	Update(ctx context.Context, user *User) error
	Delete(ctx context.Context, user *User) error
}
