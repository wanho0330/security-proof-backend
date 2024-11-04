package service

import (
	"context"
	"errors"

	apiv1 "buf.build/gen/go/wanho/security-proof-api/protocolbuffers/go/api/v1"

	"security-proof/internal/db/security_proof/user/model"
	"security-proof/internal/user/repository"
	"security-proof/pkg/auth"
	"security-proof/pkg/constants"
)

// UserQuery struct is composed of a Token and a UserQuerier.
type UserQuery struct {
	token       *auth.Token
	userQuerier repository.UserQuerier
}

// NewUserQuery function is returning a UserQuery accepting a Token and a UserQuerier.
func NewUserQuery(token *auth.Token, userQuerier repository.UserQuerier) *UserQuery {
	return &UserQuery{
		token:       token,
		userQuerier: userQuerier,
	}
}

// ReadUserByIdx method is retuning a user and an error, accepting a context, reading index and an access token.
func (q *UserQuery) ReadUserByIdx(ctx context.Context, idx int32, accessToken string) (*apiv1.User, error) {
	_, _, err := q.token.ValidateToken(accessToken)
	if err != nil {
		return nil, errors.Join(constants.ErrUserRead, err)
	}

	user, err := q.userQuerier.ReadUserByIdx(ctx, idx)
	if err != nil {
		return nil, errors.Join(constants.ErrUserRead, err)
	}

	return conv.ModelToProto(user), nil
}

// ReadUserByID method is returning a user and an error, accepting a context, a reading id and an access token.
func (q *UserQuery) ReadUserByID(ctx context.Context, id string, accessToken string) (*apiv1.User, error) {
	_, _, err := q.token.ValidateToken(accessToken)
	if err != nil {
		return nil, errors.Join(constants.ErrUserRead, err)
	}

	user, err := q.userQuerier.ReadUserByID(ctx, id)
	if err != nil {
		return nil, errors.Join(constants.ErrUserRead, err)
	}

	return conv.ModelToProto(user), nil
}

// DuplicateUserID method is returning error accepting a context and a reading id.
func (q *UserQuery) DuplicateUserID(ctx context.Context, id string) error {
	_, err := q.userQuerier.ReadUserByID(ctx, id)
	if errors.Is(err, constants.ErrItemNotFound) {
		return nil
	}

	return constants.ErrUserIDDuplicate
}

// ListUsers method is returning users and an error, accepting a context, a reading id and an access token.
func (q *UserQuery) ListUsers(ctx context.Context, id string, accessToken string) ([]*apiv1.User, error) {
	_, _, err := q.token.ValidateToken(accessToken)
	if err != nil {
		return nil, errors.Join(constants.ErrUsersList, err)
	}

	var users []*model.User
	if id == "" {
		users, err = q.userQuerier.AllUsers(ctx)
	} else {
		users, err = q.userQuerier.SearchUsers(ctx, id)
	}

	if err != nil {
		return nil, errors.Join(constants.ErrUsersList, err)
	}

	result := make([]*apiv1.User, len(users))
	for i, user := range users {
		result[i] = conv.ModelToProto(user)
	}

	return result, nil
}
