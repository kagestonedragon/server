package user

import (
	"context"
	"errors"
)

type cacheRepository struct {
	users    []User
	usersMap map[uint64]*User
}

// Прочитать mutex-ы
func initCacheRepository() Repository {
	return &cacheRepository{
		users:    make([]User, 0),
		usersMap: make(map[uint64]*User, 0),
	}
}

func (c *cacheRepository) Add(ctx context.Context, user User) (User, error) {
	c.users = append(c.users, user)
	c.usersMap[user.Id] = &user

	return user, nil
}

func (c *cacheRepository) Update(ctx context.Context, user User) error {
	if user, exists := c.usersMap[user.Id]; exists {
		if err := c.DeleteById(ctx, user.Id); err != nil {
			return err
		}
	}

	if _, err := c.Add(ctx, user); err != nil {
		return err
	}

	return nil
}

func (c *cacheRepository) DeleteById(ctx context.Context, id uint64) error {
	delete(c.usersMap, id)

	return nil
}

func (c *cacheRepository) GetById(ctx context.Context, id uint64) (User, error) {
	if user, exists := c.usersMap[id]; exists {
		return *user, nil
	}

	return User{}, errors.New("user not found in cache")
}

func (c *cacheRepository) GetList(ctx context.Context) ([]User, error) {
	return c.users, nil
}

func (c *cacheRepository) Reset(ctx context.Context) error {
	c.users = make([]User, 0)
	c.usersMap = make(map[uint64]*User, 0)

	return nil
}
