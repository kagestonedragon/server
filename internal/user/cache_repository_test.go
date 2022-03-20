package user

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCacheRepository(t *testing.T) {
	u1 := &User{
		Id:     1,
		Name:   "1",
		Active: true,
	}
	u2 := &User{
		Id:     2,
		Name:   "2",
		Active: false,
	}
	users := []*User{u1, u2}

	a := assert.New(t)
	ctx := context.Background()
	c := initCacheRepository(users)

	t.Run("GetById", func(t *testing.T) {
		u, err := c.GetById(ctx, 1)
		a.Equal(u1, u)
		a.Nil(err)
	})

	t.Run("Add", func(t *testing.T) {
		user := &User{
			Id:     3,
			Name:   "3",
			Active: true,
		}

		err := c.Add(ctx, user)
		a.Nil(err)

		u, err := c.GetById(ctx, 3)
		a.Equal(user, u)
		a.Nil(err)
	})

	t.Run("GetList", func(t *testing.T) {
		u, err := c.GetList(ctx)
		a.True(len(u) > 0)
		a.Nil(err)
	})

	t.Run("DeleteById", func(t *testing.T) {
		err := c.DeleteById(ctx, 1)
		a.Nil(err)

		u, err := c.GetById(ctx, 1)
		a.Nil(u)
		a.NotNil(err)
	})

	t.Run("Update", func(t *testing.T) {
		u, err := c.GetById(ctx, 2)
		a.Equal(u2, u)
		a.Nil(err)

		user := &User{
			Id:     2,
			Name:   "test_name",
			Active: false,
		}

		err = c.Update(ctx, user)
		a.Nil(err)

		u, err = c.GetById(ctx, 2)
		a.Equal(user, u)
		a.Nil(err)
	})

	t.Run("Reset", func(t *testing.T) {
		err := c.Reset(ctx, []*User{})
		a.Nil(err)

		u, err := c.GetList(ctx)
		a.True(len(u) == 0)
		a.Nil(err)
	})
}
