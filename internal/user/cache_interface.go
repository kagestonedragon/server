package user

import "context"

type cache interface {
	Repository

	Reset(ctx context.Context, users []*User) error
}
