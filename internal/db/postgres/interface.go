package postgres

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)


type Postgres interface {
	Connection
	Constraint
}

//go:generate mockgen -destination instance_mock.go -package postgres  git.repo.services.lenvendo.ru/grade-factor/ss/internal/postgres Connection
type Connection interface {
	Ping(ctx context.Context) error
	GetMasterConn(context.Context) (*pgxpool.Conn, error)
	GetReplicaConn(context.Context) (*pgxpool.Conn, error)
	Close()
}


type Constraint interface {
	GetConstraintName(context.Context, string, []string) (string, error)
}
