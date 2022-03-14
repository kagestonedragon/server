package user

import "context"

type cache interface {
	Reset(ctx context.Context) error
}
