package user

import (
	"context"
	"git.repo.services.lenvendo.ru/grade-factor/echo/internal/db/postgres"
	"github.com/pkg/errors"
)

const (
	AddUserSqlTemplate        = `insert into "user"."users" ("name", "is_active") values ($1, $2) returning "id"`
	UpdateUserSqlTemplate     = `update "user"."users" set "name"=$1, "active"=$2 where "id"=$3`
	DeleteUserByIdSqlTemplate = `delete from "user"."users" where "id"=$1`
	GetUserByIdSqlTemplate    = `select * from "user"."users" where "id"=$1`
	GetUsersSqlTemplate       = `select * from "user"."users"`
)

type pgRepository struct {
	conn postgres.Connection
}

func initPgRepository(conn postgres.Connection) Repository {
	return &pgRepository{
		conn: conn,
	}
}

func (pg *pgRepository) Add(ctx context.Context, user User) (User, error) {
	conn, err := pg.conn.GetMasterConn(ctx)

	defer conn.Release()

	if err != nil {
		return user, errors.Wrap(err, "get master connection")
	}

	err = conn.QueryRow(ctx, AddUserSqlTemplate, user.Name, user.Active).Scan(&user.Id)
	if err != nil {
		return user, errors.Wrap(err, "add user to database")
	}

	return user, nil
}

func (pg *pgRepository) Update(ctx context.Context, user User) error {
	conn, err := pg.conn.GetMasterConn(ctx)

	defer conn.Release()

	if err != nil {
		return errors.Wrap(err, "get master connection")
	}

	_, err = conn.Exec(ctx, UpdateUserSqlTemplate, user.Name, user.Active)
	if err != nil {
		return errors.Wrap(err, "update user in database")
	}

	return nil
}

func (pg *pgRepository) DeleteById(ctx context.Context, id uint64) error {
	conn, err := pg.conn.GetMasterConn(ctx)

	defer conn.Release()

	if err != nil {
		return errors.Wrap(err, "get replica connection")
	}

	if _, err = conn.Exec(ctx, DeleteUserByIdSqlTemplate, id); err != nil {
		return errors.Wrap(err, "delete user from database")
	}

	return nil
}

func (pg *pgRepository) GetById(ctx context.Context, id uint64) (User, error) {
	conn, err := pg.conn.GetReplicaConn(ctx)

	defer conn.Release()

	user := User{}

	if err != nil {
		return user, errors.Wrap(err, "get replica connection")
	}

	err = conn.QueryRow(ctx, GetUserByIdSqlTemplate, id).Scan(&user.Id, &user.Name, &user.Active)
	if err != nil {
		return user, errors.Wrap(err, "get user from database")
	}

	return user, nil
}

// убрать переопредление переменных
// протестить
func (pg *pgRepository) GetList(ctx context.Context) ([]User, error) {
	conn, err := pg.conn.GetReplicaConn(ctx)

	defer conn.Release()

	users := make([]User, 0)

	if err != nil {
		return users, errors.Wrap(err, "get replica connection")
	}

	res, err := conn.Query(ctx, GetUsersSqlTemplate)

	for res.Next() {
		user := User{}

		if err := res.Scan(&user.Id, &user.Name, &user.Active); err != nil {
			return users, errors.Wrap(err, "scan error")
		}

		users = append(users, user)
	}

	return users, nil
}

func (pg *pgRepository) Reset(ctx context.Context) error {
	return errors.New("cannot reset")
}
