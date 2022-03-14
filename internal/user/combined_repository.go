package user

import (
	"context"
	"git.repo.services.lenvendo.ru/grade-factor/echo/internal/db/postgres"
	"github.com/pkg/errors"
)

type CombinedRepository struct {
	postgres Repository
	cache    Repository
}

func NewCombinedRepository(conn postgres.Connection) Repository {
	return &CombinedRepository{
		postgres: initPgRepository(conn),
		cache:    initCacheRepository(),
	}
}

// убрать переопредление переменных
func (c *CombinedRepository) Add(ctx context.Context, user User) (User, error) {
	if user, err := c.postgres.Add(ctx, user); err != nil {
		return user, errors.Wrap(err, "add user to database")
	}

	if user, err := c.cache.Add(ctx, user); err != nil {
		return user, errors.Wrap(err, "add user to cache")
	}

	return user, nil
}

func (c *CombinedRepository) Update(ctx context.Context, user User) error {
	if err := c.postgres.Update(ctx, user); err != nil {
		return errors.Wrap(err, "update user in database")
	}

	if err := c.cache.Update(ctx, user); err != nil {
		return errors.Wrap(err, "update user in cache")
	}

	return nil
}

func (c *CombinedRepository) DeleteById(ctx context.Context, id uint64) error {
	if err := c.postgres.DeleteById(ctx, id); err != nil {
		return errors.Wrap(err, "delete user from database")
	}

	if err := c.cache.DeleteById(ctx, id); err != nil {
		return errors.Wrap(err, "delete cache from database")
	}

	return nil
}

// убрать переопредление переменных
func (c *CombinedRepository) GetById(ctx context.Context, id uint64) (User, error) {
	if user, err := c.cache.GetById(ctx, id); err == nil {
		return user, nil
	}

	user, err := c.postgres.GetById(ctx, id)
	if err != nil {
		return user, errors.Wrap(err, "get user from database")
	}

	if user, err := c.cache.Add(ctx, user); err != nil {
		return user, errors.Wrap(err, "add user to cache")
	}

	return user, nil
}

func (c *CombinedRepository) GetList(ctx context.Context) ([]User, error) {
	if users, err := c.cache.GetList(ctx); len(users) > 0 && err == nil {
		return users, nil
	}

	users, err := c.postgres.GetList(ctx)
	if err != nil {
		return users, err
	}

	return users, nil
}

func (c *CombinedRepository) Reset(ctx context.Context) error {
	return c.cache.Reset(ctx)
}
