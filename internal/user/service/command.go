// Package service returns data for the controller layer, processing business logic from the repository.
package service

import (
	"context"
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"strconv"
	"time"

	apiv1 "buf.build/gen/go/wanho/security-proof-api/protocolbuffers/go/api/v1"

	"security-proof/internal/user/convert"
	"security-proof/internal/user/repository"
	"security-proof/pkg/auth"
	"security-proof/pkg/constants"
)

var conv = convert.ServiceConverterImpl{}

// UserCommand struct is composed of a Token, a UserCommander and a UserQuerier.
type UserCommand struct {
	token         *auth.Token
	userCommander repository.UserCommander
	userQuerier   repository.UserQuerier
}

// NewUserCommand function is returning a UserCommand interface, accepting a Token, a UserCommander and UserQuerier.
func NewUserCommand(token *auth.Token, userCommander repository.UserCommander, userQuerier repository.UserQuerier) *UserCommand {
	return &UserCommand{
		token:         token,
		userCommander: userCommander,
		userQuerier:   userQuerier,
	}
}

// CreateUser method is returning a created index and an error, accepting a context, a user and an access token.
func (c *UserCommand) CreateUser(ctx context.Context, user *apiv1.User, accessToken string) (int32, error) {
	_, role, err := c.token.ValidateToken(accessToken)
	if err != nil {
		return 0, errors.Join(constants.ErrUserCreate, err)
	}

	if role != constants.RoleAdmin {
		return 0, errors.Join(constants.ErrUserCreate, constants.ErrTokenRoleAuth)
	}

	userModel, err := c.userQuerier.ReadUserByID(ctx, user.Id)
	if err != nil && !errors.Is(err, constants.ErrItemNotFound) {
		return 0, errors.Join(constants.ErrUserCreate, err)
	}

	if userModel != nil {
		return 0, errors.Join(constants.ErrUserCreate, constants.ErrUserIDDuplicate)
	}

	user.CreatedAt = convert.TimeToPTimestamppb(time.Now())
	user.Passwd = c.strToSha(user.Passwd)

	idx, err := c.userCommander.CreateUser(ctx, conv.ProtoToModel(user), nil)
	if err != nil {
		return 0, errors.Join(constants.ErrUserCreate, err)
	}

	return idx, nil
}

// UpdateUser method is returning an updated index and an error, accepting a context, an updating user, an access token.
func (c *UserCommand) UpdateUser(ctx context.Context, user *apiv1.User, accessToken string) (int32, error) {
	_, role, err := c.token.ValidateToken(accessToken)
	if err != nil {
		return 0, errors.Join(constants.ErrUserUpdate, err)
	}

	if role != constants.RoleAdmin {
		return 0, errors.Join(constants.ErrUserUpdate, constants.ErrTokenRoleAuth)
	}

	_, err = c.userQuerier.ReadUserByIdx(ctx, user.Idx)
	if err != nil {
		return 0, errors.Join(constants.ErrUserUpdate, err)
	}

	user.UpdatedAt = convert.TimeToPTimestamppb(time.Now())
	idx, err := c.userCommander.UpdateUser(ctx, conv.ProtoToModel(user), nil)
	if err != nil {
		return 0, errors.Join(constants.ErrUserUpdate, err)
	}

	return idx, nil
}

// DeleteUser method is returning an error accepting a context, a deleting index and an access token.
func (c *UserCommand) DeleteUser(ctx context.Context, idx int32, accessToken string) error {
	_, role, err := c.token.ValidateToken(accessToken)
	if err != nil {
		return errors.Join(constants.ErrUserDelete, err)
	}

	if role != constants.RoleAdmin {
		return errors.Join(constants.ErrUserDelete, constants.ErrTokenRoleAuth)
	}

	_, err = c.userQuerier.ReadUserByIdx(ctx, idx)
	if err != nil {
		return errors.Join(constants.ErrUserDelete, err)
	}

	err = c.userCommander.DeleteUser(ctx, idx, nil)
	if err != nil {
		return errors.Join(constants.ErrUserDelete, err)
	}

	return nil
}

// SignInUser method is returning an access token, a refresh token and an error, accepting a context and a user.
func (c *UserCommand) SignInUser(ctx context.Context, user *apiv1.User) (string, string, error) {
	readUser, err := c.userQuerier.SignInUser(ctx, user.Id)
	if err != nil {
		return "", "", errors.Join(constants.ErrUserSignIn, err)
	}

	if c.strToSha(readUser.Passwd) != c.strToSha(user.Passwd) {
		return "", "", errors.Join(constants.ErrUserSignIn, constants.ErrItemNotFound)
	}

	idxStr := strconv.Itoa(int(readUser.Idx))
	accessToken, refreshToken, err := c.token.CreateToken(ctx, idxStr, readUser.Role)
	if err != nil {
		return "", "", errors.Join(constants.ErrUserSignIn, err)
	}

	return accessToken, refreshToken, nil
}

// SingOutUser method is returning an error, accepting a context and an access token.
func (c *UserCommand) SingOutUser(ctx context.Context, accessToken string) error {
	err := c.token.DeleteToken(ctx, accessToken)

	if err != nil {
		return errors.Join(constants.ErrUserSignIn, err)
	}

	return nil
}

// RotateRefreshToken method is returning a new access token, a new refresh token and an error, accepting a context and a old refresh token.
func (c *UserCommand) RotateRefreshToken(ctx context.Context, refreshToken string) (string, string, error) {
	if refreshToken == "" {
		return "", "", errors.Join(constants.ErrUserToken, constants.ErrItemNotFound)
	}

	newAccessToken, newRefreshToken, err := c.token.RotateRefreshToken(ctx, refreshToken)
	if err != nil {
		return "", "", errors.Join(constants.ErrUserToken, err)
	}

	return newAccessToken, newRefreshToken, nil
}

func (c *UserCommand) strToSha(s string) string {
	hash := sha512.New()
	hash.Write([]byte(s))

	hashBytes := hash.Sum(nil)
	hashString := hex.EncodeToString(hashBytes)

	return hashString

}
