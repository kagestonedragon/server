package user

import (
	"context"
	"github.com/pkg/errors"
	"time"
)

type cacheableRepository struct {
	db    Repository
	cache cache
}

func NewCacheableRepository(ctx context.Context, db Repository, d time.Duration) (CacheableRepository, error) {
	users, err := db.GetList(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "initialize repository")
	}

	r := &cacheableRepository{
		db:    db,
		cache: initCacheRepository(users),
	}

	if d > 0 {
		go r.runCacheCleaner(ctx, d)
	}

	return r, nil
}

func (r *cacheableRepository) Add(ctx context.Context, u *User) error {
	if err := r.db.Add(ctx, u); err != nil {
		return err
	}

	if err := r.cache.Add(ctx, u); err != nil {
		return err
	}

	return nil
}

func (r *cacheableRepository) Update(ctx context.Context, user *User) error {
	if err := r.db.Update(ctx, user); err != nil {
		return err
	}

	if err := r.cache.Update(ctx, user); err != nil {
		return err
	}

	return nil
}

func (r *cacheableRepository) DeleteById(ctx context.Context, id uint64) error {
	if err := r.db.DeleteById(ctx, id); err != nil {
		return err
	}

	if err := r.cache.DeleteById(ctx, id); err != nil {
		return err
	}

	return nil
}

func (r *cacheableRepository) GetById(ctx context.Context, id uint64) (*User, error) {
	if user, err := r.cache.GetById(ctx, id); err == nil {
		return user, nil
	}

	user, err := r.db.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	if err := r.cache.Add(ctx, user); err != nil {
		return user, err
	}

	return user, nil
}

func (r *cacheableRepository) GetList(ctx context.Context) ([]*User, error) {
	if users, err := r.cache.GetList(ctx); len(users) > 0 && err == nil {
		return users, nil
	}

	users, err := r.db.GetList(ctx)
	if err != nil {
		return nil, err
	}

	for _, u := range users {
		if err := r.cache.Add(ctx, u); err != nil {
			return users, err
		}
	}

	return users, nil
}

func (r *cacheableRepository) Reset(ctx context.Context) error {
	users, err := r.db.GetList(ctx)
	if err != nil {
		return errors.Wrap(err, "repository reset")
	}

	if err := r.cache.Reset(ctx, users); err != nil {
		return errors.Wrap(err, "repository reset")
	}

	return nil
}

func (r *cacheableRepository) runCacheCleaner(ctx context.Context, d time.Duration) {
	for _ = range time.Tick(d) {
		if err := r.Reset(ctx); err != nil {
			// TODO something
		}
	}
}
