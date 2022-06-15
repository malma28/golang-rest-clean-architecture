package database

import (
	"context"
)

type Result interface {
	LastInsertId() (int64, error)
	RowsAffected() (int64, error)
}

type Row interface {
	Scan(dest ...any) error
}

type Rows interface {
	Scan(dest ...any) error
	Next() bool
	Err() error
	Close() error
}

type Tx interface {
	Exec(ctx context.Context, query string, args ...any) (Result, error)
	Query(ctx context.Context, query string, args ...any) (Rows, error)
	QueryRow(ctx context.Context, query string, args ...any) (Row, error)
	Commit() error
	Rollback() error
}

type SQL interface {
	Exec(ctx context.Context, query string, args ...any) (Result, error)
	Query(ctx context.Context, query string, args ...any) (Rows, error)
	QueryRow(ctx context.Context, query string, args ...any) (Row, error)
	Tx(ctx context.Context) (Tx, error)
	Close() error
}
