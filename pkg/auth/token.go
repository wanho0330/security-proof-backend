package auth

import (
	"context"
	"errors"
	"time"

	"github.com/Netflix/go-env"
	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwt"

	"security-proof/pkg/constants"
	"security-proof/pkg/manage/db"
)

// Token struct is composed of a TokenRepo.
type Token struct {
	tokenRepo TokenRepo
}

// NewToken function is returning a Token accepting a TokenRepo.
func NewToken(tokenRepo TokenRepo) *Token {
	return &Token{tokenRepo: tokenRepo}
}

// CreateToken method is returning an access token and a refresh token, accepting a context, an index and role.
func (t *Token) CreateToken(ctx context.Context, idx string, role int32) (accessToken string, refreshToken string, err error) {
	config := jwtConfig{}
	_, err = env.UnmarshalFromEnviron(&config)
	if err != nil {
		return "", "", errors.Join(constants.ErrTokenCreate, err)
	}

	key := []byte(config.SecretKey)

	access := jwt.New()
	if err = access.Set(jwt.SubjectKey, idx); err != nil {
		return "", "", errors.Join(constants.ErrTokenCreate, err)
	}
	if err = access.Set(jwtRole, role); err != nil {
		return "", "", errors.Join(constants.ErrTokenCreate, err)
	}
	if err = access.Set(jwt.ExpirationKey, time.Now().Add(config.AccessTokenTime).Unix()); err != nil {
		return "", "", errors.Join(constants.ErrTokenCreate, err)
	}
	signedAccess, err := jwt.Sign(access, jwt.WithKey(jwa.HS256, key))
	if err != nil {
		return "", "", errors.Join(constants.ErrTokenCreate, err)
	}

	refresh := jwt.New()
	if err = refresh.Set(jwt.SubjectKey, idx); err != nil {
		return "", "", errors.Join(constants.ErrTokenCreate, err)
	}
	if err = refresh.Set(jwtRole, role); err != nil {
		return "", "", errors.Join(constants.ErrTokenCreate, err)
	}
	if err = refresh.Set(jwt.ExpirationKey, time.Now().Add(config.RefreshTokenTime).Unix()); err != nil {
		return "", "", errors.Join(constants.ErrTokenCreate, err)
	}
	signedRefresh, err := jwt.Sign(refresh, jwt.WithKey(jwa.HS256, key))
	if err != nil {
		return "", "", errors.Join(constants.ErrTokenCreate, err)
	}

	err = t.tokenRepo.SaveToken(ctx, string(signedRefresh))
	if err != nil {
		return "", "", errors.Join(constants.ErrTokenCreate, err)
	}

	return string(signedAccess), string(signedRefresh), nil
}

// ValidateToken method is returning an index, a role and an error, accepting signed token.
func (t *Token) ValidateToken(signedToken string) (idx string, role int32, err error) {
	config := jwtConfig{}
	_, err = env.UnmarshalFromEnviron(&config)
	if err != nil {
		return "", 0, errors.Join(constants.ErrTokenValidate, err)
	}

	key := []byte(config.SecretKey)

	token, err := jwt.Parse([]byte(signedToken), jwt.WithKey(jwa.HS256, key))
	if err != nil {
		return "", 0, errors.Join(constants.ErrTokenValidate, err)
	}

	if err = jwt.Validate(token); err != nil {
		return "", 0, errors.Join(constants.ErrTokenValidate, err)
	}

	roleAny, exist := token.Get(jwtRole)
	if !exist {
		return "", 0, errors.Join(constants.ErrTokenValidate, constants.ErrTokenRoleMissing)
	}

	return token.Subject(), int32(roleAny.(float64)), nil
}

// RotateRefreshToken method is returning a new access token, a new refresh token and an error, accepting a context and a refresh token.
func (t *Token) RotateRefreshToken(ctx context.Context, refreshToken string) (newAccessToken string, newRefreshToken string, err error) {
	tokenConfig := db.TokenConfig{}
	_, err = env.UnmarshalFromEnviron(&tokenConfig)
	if err != nil {
		return "", "", errors.Join(constants.ErrTokenRotation, err)
	}

	jwtConfig := jwtConfig{}
	_, err = env.UnmarshalFromEnviron(&jwtConfig)
	if err != nil {
		return "", "", errors.Join(constants.ErrTokenRotation, err)
	}

	idx, role, err := t.ValidateToken(refreshToken)
	if err != nil {
		return "", "", errors.Join(constants.ErrTokenRotation, err)
	}

	savedRefreshToken, err := t.tokenRepo.ReadTokenByIdx(ctx, idx)
	if err != nil {
		return "", "", errors.Join(constants.ErrTokenRotation, err)
	}

	if refreshToken != savedRefreshToken {
		return "", "", errors.Join(constants.ErrTokenRotation, constants.ErrTokenDoesNotMatch)
	}

	newAccessToken, newRefreshToken, err = t.CreateToken(ctx, idx, role)
	if err != nil {
		return "", "", errors.Join(constants.ErrTokenRotation, err)
	}

	err = t.tokenRepo.SaveToken(ctx, newRefreshToken)
	if err != nil {
		return "", "", errors.Join(constants.ErrTokenRotation, err)
	}

	return newAccessToken, newRefreshToken, nil
}

// DeleteToken method is returning an error, accepting a context and a signed token.
func (t *Token) DeleteToken(ctx context.Context, signedToken string) (err error) {
	config := jwtConfig{}
	_, err = env.UnmarshalFromEnviron(&config)
	if err != nil {
		return errors.Join(constants.ErrTokenValidate, err)
	}

	key := []byte(config.SecretKey)

	token, err := jwt.Parse([]byte(signedToken), jwt.WithKey(jwa.HS256, key))
	if err != nil {
		return errors.Join(constants.ErrTokenValidate, err)
	}

	if err = jwt.Validate(token); err != nil {
		return errors.Join(constants.ErrTokenValidate, err)
	}

	idxInterface, exist := token.Get(jwt.SubjectKey)
	if !exist {
		return errors.Join(constants.ErrTokenValidate, constants.ErrTokenRoleMissing)
	}

	err = t.DeleteToken(ctx, idxInterface.(string))
	if err != nil {
		return errors.Join(constants.ErrTokenDelete, err)
	}

	return nil
}
