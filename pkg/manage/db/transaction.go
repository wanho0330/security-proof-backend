package db

import (
	"context"
	"database/sql"
)

// Beginner interface is defining data to Begin Transaction.
type Beginner interface {
	Begin(ctx context.Context) (tx *sql.Tx, err error)
}

// Commiter interface is defining data to Commit Transaction.
type Commiter interface {
	Commit(ctx context.Context, tx *sql.Tx) error
}

// Rollbacker interface is defining data to Rollback Transaction.
type Rollbacker interface {
	Rollback(ctx context.Context, tx *sql.Tx) error
}
