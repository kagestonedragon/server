package user

import (
	"context"
	"github.com/kagestonedragon/server/internal/db/postgres"
	"github.com/pkg/errors"
)

const (
	AddUserSqlTemplate        = `insert into "user"."users" ("name", "is_active") values ($1, $2) returning "id"`
	UpdateUserSqlTemplate     = `update "user"."users" set "name"=$1, "is_active"=$2 where "id"=$3`
	DeleteUserByIdSqlTemplate = `delete from "user"."users" where "id"=$1`
	GetUserByIdSqlTemplate    = `select * from "user"."users" where "id"=$1`
	GetUsersSqlTemplate       = `select * from "user"."users"`
)

type postgreSqlRepository struct {
	conn postgres.Connection
}

func NewPostgreSqlRepository(conn postgres.Connection) Repository {
	return &postgreSqlRepository{
		conn: conn,
	}
}

func (p *postgreSqlRepository) Add(ctx context.Context, u *User) error {
	conn, err := p.conn.GetMasterConn(ctx)
	if err != nil {
		return errors.Wrap(err, "get master connection")
	}

	defer conn.Release()

	if err = conn.QueryRow(ctx, AddUserSqlTemplate, u.Name, u.Active).Scan(&u.Id); err != nil {
		return errors.Wrap(err, "add user to database")
	}

	return nil
}

func (p *postgreSqlRepository) Update(ctx context.Context, u *User) error {
	conn, err := p.conn.GetMasterConn(ctx)
	if err != nil {
		return errors.Wrap(err, "get master connection")
	}

	defer conn.Release()

	if r, err := conn.Exec(ctx, UpdateUserSqlTemplate, u.Name, u.Active, u.Id); err != nil {
		return errors.Wrap(err, "update user in database")
	} else if r.RowsAffected() == 0 {
		return errors.New("user not found")
	}

	return nil
}

func (p *postgreSqlRepository) DeleteById(ctx context.Context, id uint64) error {
	conn, err := p.conn.GetMasterConn(ctx)
	if err != nil {
		return errors.Wrap(err, "get master connection")
	}

	defer conn.Release()

	if r, err := conn.Exec(ctx, DeleteUserByIdSqlTemplate, id); err != nil {
		return errors.Wrap(err, "delete user from database")
	} else if r.RowsAffected() == 0 {
		return errors.New("user not found")
	}

	return nil
}

func (p *postgreSqlRepository) GetById(ctx context.Context, id uint64) (*User, error) {
	conn, err := p.conn.GetMasterConn(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get master connection")
	}

	defer conn.Release()

	u := &User{}

	err = conn.QueryRow(ctx, GetUserByIdSqlTemplate, id).Scan(&u.Id, &u.Name, &u.Active)
	if err != nil {
		return nil, errors.Wrap(err, "get user from database")
	}

	return u, nil
}

func (p *postgreSqlRepository) GetList(ctx context.Context) ([]*User, error) {
	conn, err := p.conn.GetMasterConn(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get master connection")
	}

	defer conn.Release()

	users := make([]*User, 0)

	res, err := conn.Query(ctx, GetUsersSqlTemplate)
	if err != nil {
		return nil, errors.Wrap(err, "get users from database")
	}

	for res.Next() {
		user := &User{}

		if err := res.Scan(&user.Id, &user.Name, &user.Active); err != nil {
			return nil, errors.Wrap(err, "get users from database on fetch")
		}

		users = append(users, user)
	}

	return users, nil
}

func (p *postgreSqlRepository) Reset(ctx context.Context) error {
	return nil
}
