package service

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"security-proof/internal/db/security_proof/user/model"
	"security-proof/internal/user/repository"
	"security-proof/pkg/constants"
)

var query = newMockQuery()

func TestQuery_ReadUserByIdx(t *testing.T) {
	accessToken, _, err := mockToken.CreateToken(ctx, "1", constants.RoleAdmin)
	assert.NoError(t, err, "토큰 생성 중 에러가 발생하지 않았습니다.")

	t.Run("유저 인덱스로 읽기", func(t *testing.T) {
		user, err := query.ReadUserByIdx(ctx, 1, accessToken)
		assert.NoError(t, err)
		assert.NotNil(t, user)
	})
}

func TestQuery_ReadUserByID(t *testing.T) {
	accessToken, _, err := mockToken.CreateToken(ctx, "1", constants.RoleAdmin)
	assert.NoError(t, err, "토큰 생성 중 에러가 발생하지 않았습니다.")

	t.Run("유저 아이디로 읽기", func(t *testing.T) {
		user, err := query.ReadUserByID(ctx, "test", accessToken)
		assert.NoError(t, err)
		assert.NotNil(t, user)
	})
}

func TestQuery_DuplicateUserID(t *testing.T) {
	t.Run("중복 아이디 확인", func(t *testing.T) {
		err := query.DuplicateUserID(ctx, "test")
		assert.True(t, errors.Is(err, constants.ErrUserIDDuplicate), "발생한 에러는 ID Duplicate 에러입니다.")
	})
}

func newMockQuery() *UserQuery {
	return NewUserQuery(mockToken, mockQuery)
}

var mockQuery = &repository.MockUserQuery{
	ReadUserByIdxFn: func(ctx context.Context, idx int32) (user *model.User, err error) {
		if idx == 1 {
			return &model.User{
				Idx:    1,
				ID:     "test",
				Passwd: "test",
				Name:   "test",
				Email:  "test",
				Role:   0,
			}, nil
		}
		return nil, constants.ErrItemNotFound

	},
	ReadUserByIDFn: func(ctx context.Context, id string) (*model.User, error) {
		if id == "admin" {
			return nil, constants.ErrUserIDDuplicate
		} else if id == "test" {
			return &model.User{
				Idx:    1,
				ID:     "test",
				Passwd: "test",
				Name:   "test",
				Email:  "test",
				Role:   0,
			}, nil
		}
		return nil, nil

	},
	AllUsersFn: func(ctx context.Context) ([]*model.User, error) {
		return nil, nil
	},
	SearchUsersFn: func(ctx context.Context, id string) ([]*model.User, error) {
		return nil, nil
	},
	SignInUserFn: func(ctx context.Context, id string) (user *model.User, err error) {
		if id == "test" {
			return &model.User{
				Idx:    1,
				ID:     "test",
				Passwd: "test",
			}, nil
		}

		return nil, constants.ErrItemNotFound

	},
}
