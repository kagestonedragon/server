package user

import (
	"context"
	"errors"
	"sync"
)

type cacheRepository struct {
	sync.RWMutex

	users    []*User
	usersMap map[uint64]int
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

func (c *cacheRepository) DeleteById(ctx context.Context, id uint64) error {
	c.RWMutex.Lock()
	defer c.RWMutex.Unlock()

	if k, e := c.usersMap[id]; e {
		c.delete(k)
		delete(c.usersMap, id)
	}

	return nil
}

func (c *cacheRepository) GetById(ctx context.Context, id uint64) (*User, error) {
	c.RWMutex.RLock()
	defer c.RWMutex.RUnlock()

	if k, e := c.usersMap[id]; e {
		return c.users[k], nil
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
	c.usersMap[u.Id] = len(c.users) - 1
}

func (c *cacheRepository) addList(users []*User) {
	for _, u := range users {
		c.add(u)
	}
}

func (c *cacheRepository) delete(k int) {
	c.users[k] = c.users[len(c.users)-1]
	c.users[len(c.users)-1] = nil
	c.users = c.users[:len(c.users)-1]
}

func (c *cacheRepository) clean() {
	c.users = make([]*User, 0)
	c.usersMap = make(map[uint64]int, 0)
}
