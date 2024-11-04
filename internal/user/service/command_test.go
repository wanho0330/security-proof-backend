package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"testing"
	"time"

	apiv1 "buf.build/gen/go/wanho/security-proof-api/protocolbuffers/go/api/v1"
	"github.com/stretchr/testify/assert"

	"security-proof/internal/db/security_proof/user/model"
	"security-proof/internal/user/repository"
	"security-proof/pkg/auth"
	"security-proof/pkg/constants"
)

var ctx, cancel = context.WithTimeout(context.Background(), time.Second*5)

var command = newMockCommand()

func TestCommand_Create(t *testing.T) {
	defer cancel()

	accessToken, _, err := mockToken.CreateToken(ctx, "1", constants.RoleAdmin)
	assert.NoError(t, err, "토큰 생성 중 에러가 발생하지 않았습니다.")

	t.Run("유저 추가 케이스", func(t *testing.T) {
		user := &apiv1.User{
			Id:     "test",
			Passwd: "test",
			Name:   "test",
			Email:  "test",
		}
		idx, err := command.CreateUser(ctx, user, accessToken)

		assert.NoError(t, err, "에러가 발생하지 않았습니다.")
		assert.Equal(t, int32(1), idx, "테스트 유저가 정상적으로 생성되었습니다.")
	})

	t.Run("중복된 아이디 유저 추가 케이스", func(t *testing.T) {
		user := &apiv1.User{
			Id: "admin",
		}
		_, err := command.CreateUser(ctx, user, accessToken)

		assert.Error(t, err, "예상된 에러가 발생하였습니다.")
		assert.True(t, errors.Is(err, constants.ErrUserIDDuplicate), "발생한 에러는 ID Duplicate 에러입니다.")
	})
}

func TestCommand_Update(t *testing.T) {
	defer cancel()

	accessToken, _, err := mockToken.CreateToken(ctx, "1", constants.RoleAdmin)
	assert.NoError(t, err, "토큰 생성 중 에러가 발생하지 않았습니다.")

	t.Run("유저 업데이트 케이스", func(t *testing.T) {
		user := &apiv1.User{
			Idx:    int32(1),
			Id:     "test",
			Passwd: "test",
			Name:   "test1",
			Email:  "test",
		}
		idx, err := command.UpdateUser(ctx, user, accessToken)

		assert.NoError(t, err, "에러가 발생하지 않았습니다.")
		assert.Equal(t, int32(1), idx, "테스트 유저가 정상적으로 업데이트되었습니다.")
	})

	t.Run("업데이트할 계정이 존재하지 않는 케이스", func(t *testing.T) {
		user := &apiv1.User{
			Idx: int32(2),
		}
		_, err := command.UpdateUser(ctx, user, accessToken)

		assert.Error(t, err, "예상된 에러가 발생하였습니다.")
		assert.True(t, errors.Is(err, constants.ErrItemNotFound), "발생한 에러는 ItemNotFound 에러입니다.")
	})
}

func TestCommand_Delete(t *testing.T) {
	defer cancel()

	accessToken, _, err := mockToken.CreateToken(ctx, "1", constants.RoleAdmin)
	assert.NoError(t, err, "토큰 생성 중 에러가 발생하지 않았습니다.")

	t.Run("유저 삭제 케이스", func(t *testing.T) {
		err := command.DeleteUser(context.Background(), int32(1), accessToken)
		assert.NoError(t, err, "에러가 발생하지 않았습니다.")
		assert.Equal(t, err, nil, "테스트 유저가 정상적으로 삭제되었습니다.")
	})

	t.Run("삭제할 계정 이 존재하지 않는 케이스", func(t *testing.T) {
		err := command.DeleteUser(ctx, int32(2), accessToken)
		assert.Error(t, err, "예상된 에러가 발생하였습니다.")
		assert.True(t, errors.Is(err, constants.ErrItemNotFound), "발생한 에러는 ItemNotFound 에러입니다.")
	})
}

func TestCommand_SignIn(t *testing.T) {
	defer cancel()

	t.Run("로그인 케이스", func(t *testing.T) {
		user := &apiv1.User{
			Id:     "test",
			Passwd: "test",
		}

		_, _, err := command.SignInUser(ctx, user)
		assert.NoError(t, err, "에러가 발생하지 않았습니다.")
	})
}

func newMockCommand() *UserCommand {
	return NewUserCommand(mockToken, mockCommand, mockQuery)
}

var mockTokenRepo = &auth.MockTokenRepo{
	SaveTokenFn: func(ctx context.Context, token string) error {
		return nil
	},
	ReadTokenByIdxFn: func(ctx context.Context, idx string) (string, error) {
		return "", nil
	},
	DeleteTokenFn: func(ctx context.Context, idx string) error {
		return nil
	},
}

var mockToken = auth.NewToken(mockTokenRepo)

var mockCommand = &repository.MockUserCommand{
	BeginFn: func(ctx context.Context) (tx *sql.Tx, err error) {
		fmt.Println("mock begin")
		return nil, nil
	},
	CommitFn: func(ctx context.Context, tx *sql.Tx) error {
		fmt.Println("mock commit")
		return nil
	},
	RollbackFn: func(ctx context.Context, tx *sql.Tx) error {
		fmt.Println("mock rollback")
		return nil
	},
	CreateUserFn: func(ctx context.Context, user *model.User, tx *sql.Tx) (idx int32, err error) {
		if user.ID == "test" && user.Passwd == "test" && user.Name == "test" && user.Email == "test" {
			return 1, nil
		} else if user.ID == "admin" {
			return 0, errors.Join(constants.ErrUserCreate, constants.ErrUserIDDuplicate)
		}
		return 0, errors.Join(constants.ErrUserCreate, errors.New("test error"))

	},
	UpdateUserFn: func(ctx context.Context, user *model.User, tx *sql.Tx) (idx int32, err error) {
		if user.ID == "test" && user.Passwd == "test" && user.Name == "test1" && user.Email == "test" {
			return 1, nil
		}
		return 0, constants.ErrUserUpdate

	},
	DeleteUserFn: func(ctx context.Context, idx int32, tx *sql.Tx) (err error) {
		if idx == 1 {
			return nil
		}
		return constants.ErrUserDelete

	},
}
