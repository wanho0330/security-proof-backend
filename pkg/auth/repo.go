// Package auth is a package for handling authentication-related processes.
package auth

import (
	"context"
	"errors"

	"github.com/redis/go-redis/v9"

	"security-proof/pkg/constants"
)

// TokenRepo interface is defining data related to manage token.
type TokenRepo interface {
	SaveToken(ctx context.Context, token string) error
	ReadTokenByIdx(ctx context.Context, idx string) (token string, err error)
	DeleteToken(ctx context.Context, idx string) error
}

type tokenRepo struct {
	rdb *redis.Client
}

// NewTokenRepo function is returning a TokenRepo accepting a redis client.
func NewTokenRepo(rdb *redis.Client) TokenRepo {
	return &tokenRepo{rdb: rdb}
}

// SaveToken method is returning an error accepting a context and a token.
func (r *tokenRepo) SaveToken(ctx context.Context, token string) error {
	config := &jwtConfig{}

	idx, _, err := parseToken(token)
	if err != nil {
		return errors.Join(constants.ErrTokenSaveRefresh, err)
	}

	err = r.rdb.Set(ctx, idx, token, config.RefreshTokenTime).Err()
	if err != nil {
		return errors.Join(constants.ErrTokenSaveRefresh, err)
	}

	return nil
}

// ReadTokenByIdx method is returning a token and an error, accepting a context and index.
func (r *tokenRepo) ReadTokenByIdx(ctx context.Context, idx string) (string, error) {
	token, err := r.rdb.Get(ctx, idx).Result()
	if err != nil {
		return "", errors.Join(constants.ErrTokenRead, err)
	}

	return token, nil
}

// DeleteToken method is returning an error accepting a context and index.
func (r *tokenRepo) DeleteToken(ctx context.Context, idx string) error {
	err := r.rdb.Del(ctx, idx).Err()
	if err != nil {
		return errors.Join(constants.ErrTokenDelete, err)
	}

	return nil
}
