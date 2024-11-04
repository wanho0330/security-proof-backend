package db

import (
	"context"
	"database/sql"
	"errors"

	_ "github.com/lib/pq" // postgresql
	"github.com/redis/go-redis/v9"

	"security-proof/pkg/constants"
)

// NewDB function is returning a DB and an error, accepting context and dsn.
func NewDB(_ context.Context, dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, errors.Join(constants.ErrNewDB, err)
	}

	return db, nil
}

// NewRedis function is returning a Client and an error, accepting dsn.
func NewRedis(dsn string) (*redis.Client, error) {
	opt, err := redis.ParseURL(dsn)
	if err != nil {
		return nil, errors.Join(constants.ErrNewRedis, err)
	}

	return redis.NewClient(opt), nil
}
