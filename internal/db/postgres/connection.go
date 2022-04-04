package postgres

import (
	"context"
	"fmt"

	"github.com/kagestonedragon/server/configs"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
)

const (
	DefaultMaxConnsPool = 5
)

type connection struct {
	Master  *pgxpool.Pool
	Replica *pgxpool.Pool
}

func NewConnection(ctx context.Context, cfg *configs.Config) (Connection, error) {
	var c connection
	var err error
	c.Master, err = c.connect(ctx, &cfg.Postgres.Master)
	if err != nil {
		return nil, errors.Wrap(err, "Master DB connect")
	}

	c.Replica, err = c.connect(ctx, &cfg.Postgres.Replica)
	if err != nil {
		return nil, errors.Wrap(err, "Replica DB connect")
	}

	if res := c.Ping(ctx); res != nil {
		return nil, errors.Wrap(err, "Ping DB connect")
	}

	return &c, nil
}

func (c *connection) connect(ctx context.Context, db *configs.Database) (conn *pgxpool.Pool, err error) {
	if db == nil {
		return
	}

	config, err := pgxpool.ParseConfig(fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		db.User,
		db.Password,
		db.Host,
		db.Port,
		db.DatabaseName,
		db.Secure,
	))

	if err != nil {
		return nil, errors.Wrap(err, "parse connection string")
	}

	if db.MaxConnsPool != 0 {
		config.MaxConns = int32(db.MaxConnsPool)
	} else {
		config.MaxConns = DefaultMaxConnsPool
	}

	return pgxpool.ConnectConfig(ctx, config)
}

func (c *connection) GetMasterConn(ctx context.Context) (*pgxpool.Conn, error) {
	return c.Master.Acquire(ctx)
}

func (c *connection) GetReplicaConn(ctx context.Context) (*pgxpool.Conn, error) {
	return c.Replica.Acquire(ctx)
}

func (c *connection) Ping(ctx context.Context) (err error) {
	if conn, err := c.Master.Acquire(ctx); err != nil {
		return errors.Wrap(err, "master ping error")
	} else {
		conn.Release()
	}

	if conn, err := c.Replica.Acquire(ctx); err != nil {
		return errors.Wrap(err, "replica ping error")
	} else {
		conn.Release()
	}
	return
}

func (c *connection) Close() {
	c.Master.Close()
	c.Replica.Close()
}
