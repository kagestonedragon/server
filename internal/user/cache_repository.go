package user

import (
	"context"
	"errors"
	"sync"
)

type cacheRepository struct {
	sync.RWMutex

	users    []*User
	usersMap map[uint64]*User
}

func initCacheRepository(users []*User) cache {
	c := &cacheRepository{}

	c.clean()
	c.addList(users)

	return c
}

func (c *cacheRepository) Add(ctx context.Context, u *User) error {
	c.RWMutex.Lock()
	defer c.RWMutex.Unlock()

	c.add(u)

	return nil
}

// TODO не работает нормально из-за DeleteById
func (c *cacheRepository) Update(ctx context.Context, u *User) error {
	if _, e := c.usersMap[u.Id]; e {
		if err := c.DeleteById(ctx, u.Id); err != nil {
			return err
		}
	}

	if err := c.Add(ctx, u); err != nil {
		return err
	}

	return nil
}

// TODO не работает
func (c *cacheRepository) DeleteById(ctx context.Context, id uint64) error {
	delete(c.usersMap, id)

	return nil
}

func (c *cacheRepository) GetById(ctx context.Context, id uint64) (*User, error) {
	c.RWMutex.RLock()
	defer c.RWMutex.RUnlock()

	if u, e := c.usersMap[id]; e {
		return u, nil
	}

	return nil, errors.New("user not found in cache")
}

func (c *cacheRepository) GetList(ctx context.Context) ([]*User, error) {
	return c.users, nil
}

func (c *cacheRepository) Reset(ctx context.Context, users []*User) error {
	c.RWMutex.Lock()
	defer c.RWMutex.Unlock()

	c.clean()
	c.addList(users)

	return nil
}

func (c *cacheRepository) add(u *User) {
	c.users = append(c.users, u)
	c.usersMap[u.Id] = u
}

func (c *cacheRepository) addList(users []*User) {
	for _, u := range users {
		c.add(u)
	}
}

func (c *cacheRepository) clean() {
	c.users = make([]*User, 0)
	c.usersMap = make(map[uint64]*User, 0)
}
