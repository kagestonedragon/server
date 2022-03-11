package user

import (
	"context"
	"git.repo.services.lenvendo.ru/grade-factor/echo/internal/db/postgres"
)

const (
	AddUserSqlTemplate    = `insert into "user"."users" ("id", "name", "is_active") values ($1, $2, $3)`
	UpdateUserSqlTemplate = `update "user"."users" set name=$1, active=$2 where id=$3`
	GetUserSqlTemplate    = `select * from user.users where id=$1`
	DeleteUserSqlTemplate = `delete from user.users where id=$1`
)

type DefaultRepository struct {
	conn  postgres.Connection
	cache Cache
}

func NewRepository(conn postgres.Connection) Repository {
	return &DefaultRepository{
		conn:  conn,
		cache: initCache(),
	}
}

func (r *DefaultRepository) Add(ctx context.Context, user *User) error {
	if conn, err := r.conn.GetMasterConn(ctx); err == nil {
		if _, err := conn.Query(ctx, AddUserSqlTemplate, user.Id, user.Name, user.Active); err != nil {
			return err
		}
	} else {
		return err
	}

	if err := r.cache.add(user); err != nil {
		return err
	}

	return nil
}

func (r *DefaultRepository) GetById(ctx context.Context, id uint64) (*User, error) {
	if user, err := r.cache.get(id); err == nil {
		return user, nil
	}

	if err := r.cache.delete(id); err != nil {
		return nil, err
	}

	//if conn, err := r.conn.GetReplicaConn(ctx); err == nil {
	//	// TODO не ясно что дальше
	//	row := conn.QueryRow(ctx, GetUserSqlTemplate, id)
	//} else {
	//	return nil, err
	//}

	// TODO
	return nil, nil
}

func (r *DefaultRepository) Update(ctx context.Context, user *User) error {
	if conn, err := r.conn.GetMasterConn(ctx); err != nil {
		if _, err := conn.Exec(ctx, UpdateUserSqlTemplate, user.Name, user.Active, user.Id); err != nil {
			return err
		}
	} else {
		return err
	}

	if err := r.cache.add(user); err != nil {
		return err
	}

	return nil
}

func (r *DefaultRepository) Delete(ctx context.Context, user *User) error {
	if conn, err := r.conn.GetMasterConn(ctx); err != nil {
		if _, err := conn.Exec(ctx, DeleteUserSqlTemplate, user.Id); err != nil {
			return err
		}
	} else {
		return err
	}

	if err := r.cache.delete(user.Id); err != nil {
		return err
	}

	return nil
}
