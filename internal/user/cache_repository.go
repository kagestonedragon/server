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

func (c *cacheRepository) Update(ctx context.Context, u *User) error {
	c.RWMutex.Lock()
	defer c.RWMutex.Unlock()

	if k, err := c.getStorageKeyById(u.Id); err == nil {
		c.users[k] = u
		c.usersMap[u.Id] = u
	}

	return nil
}

func (c *cacheRepository) DeleteById(ctx context.Context, id uint64) error {
	c.RWMutex.Lock()
	defer c.RWMutex.Unlock()

	// видится мне, что это плохо + долго
	if k, err := c.getStorageKeyById(id); err == nil {
		c.deleteFromStorageByKey(k)
		delete(c.usersMap, id)
	}

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
	c.RWMutex.RLock()
	defer c.RWMutex.RUnlock()

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

func (c *cacheRepository) getStorageKeyById(id uint64) (int, error) {
	for k, u := range c.users {
		if u.Id == id {
			return k, nil
		}
	}

	return 0, errors.New("user not found")
}

func (c *cacheRepository) deleteFromStorageByKey(k int) {
	copy(c.users[k:], c.users[k+1:])
	c.users[len(c.users)-1] = nil
	c.users = c.users[:len(c.users)-1]
}
