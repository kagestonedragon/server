package user

import "errors"

type cache struct {
	users map[uint64]*User
}

func initCache() Cache {
	return &cache{
		users: make(map[uint64]*User, 0),
	}
}

func (c *cache) add(user *User) error {
	c.users[user.Id] = user

	return nil
}

func (c *cache) get(id uint64) (*User, error) {
	if c.users[id] == nil {
		return nil, errors.New("user not found")
	}

	return c.users[id], nil
}

func (c *cache) delete(id uint64) error {
	delete(c.users, id)

	return nil
}

func (c *cache) reset() error {
	c.users = make(map[uint64]*User, 0)

	return nil
}
