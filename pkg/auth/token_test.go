package auth

import (
	"context"
	"errors"
	constants "security-proof/pkg/constants"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var ctx, cancel = context.WithTimeout(context.Background(), time.Second*5)
var savedRefreshToken = ""

var mockToken = initMockToken()

func TestToken_Create(t *testing.T) {
	defer cancel()

	t.Run("토큰 생성 케이스", func(t *testing.T) {
		accessToken, refreshToken, err := mockToken.CreateToken(ctx, "1", 1)
		assert.NoError(t, err, "에러가 발생하지 않았습니다.")
		assert.NotEmpty(t, accessToken, "액세스토큰이 생성되었습니다.")
		assert.NotEmpty(t, refreshToken, "리프레쉬 토큰이 생성되었습니다.")
	})
}

func TestToken_Validate(t *testing.T) {
	// 전제조건 : TokenCreate 가 정상 동작해야합니다.
	defer cancel()

	t.Run("액세스 토큰 유효성 검증 케이스", func(t *testing.T) {
		accessToken, _, err := mockToken.CreateToken(ctx, "1", 1)
		assert.NoError(t, err, "토큰 생성 중 에러가 발생하지 않았습니다.")

		idx, role, err := mockToken.ValidateToken(accessToken)
		assert.NoError(t, err, "에러가 발생하지 않았습니다.")
		assert.Equal(t, "1", idx, "액세스토큰에서 확인된 ID 동일합니다.")
		assert.Equal(t, int32(1), role, "액세스토큰에서 확인된 Role 동일합니다.")
	})

	t.Run("리프레쉬 토큰 유효성 검증 케이스", func(t *testing.T) {
		_, refreshToken, err := mockToken.CreateToken(ctx, "1", 1)
		assert.NoError(t, err, "토큰 생성 중 에러가 발생하지 않았습니다.")

		idx, role, err := mockToken.ValidateToken(refreshToken)
		assert.NoError(t, err, "에러가 발생하지 않았습니다.")
		assert.Equal(t, "1", idx, "액세스토큰에서 확인된 ID 동일합니다.")
		assert.Equal(t, int32(1), role, "액세스토큰에서 확인된 role 동일합니다.")
	})
}

func TestToken_RotateRefresh(t *testing.T) {
	// 전제조건 : TokenCreate 가 정상 동작해야합니다.
	// 새로운 시간값을 이용하여 토큰을 생성하여 비교하므로 테스트 시간이 1초 이상 걸립니다.
	defer cancel()

	t.Run("리프레쉬 토큰 재발급 케이스", func(t *testing.T) {
		_, refreshToken, err := mockToken.CreateToken(ctx, "1", 1)
		assert.NoError(t, err, "토큰 생성 중 에러가 발생하지 않았습니다.")

		savedRefreshToken = refreshToken
		time.Sleep(1 * time.Second) // 새로운 시간값으로 토큰 발급을 위해 1초 대기

		newAccessToken, newRefreshToken, err := mockToken.RotateRefreshToken(ctx, refreshToken)
		assert.NoError(t, err, "에러가 발생하지 않았습니다.")
		assert.NotEmpty(t, newAccessToken, "액세스토큰이 생성되었습니다.")
		assert.NotEmpty(t, newRefreshToken, "리프레쉬 토큰이 생성되었습니다.")
		assert.NotEqual(t, refreshToken, newRefreshToken, "리프레쉬 토큰이 갱신되었습니다.")
	})

	t.Run("리프레쉬 토큰이 만료된 케이스", func(t *testing.T) {
		_, refreshToken, err := mockToken.CreateToken(ctx, "1", 1)
		assert.NoError(t, err)

		_, _, err = mockToken.RotateRefreshToken(ctx, refreshToken)
		assert.Error(t, err, "예상된 에러가 발생하였습니다.")
		assert.True(t, errors.Is(err, constants.ErrTokenDoesNotMatch), "발생한 에러는 ErrTokenDoesNotMatch 입니다.")
	})
}

func initMockToken() *Token {
	var mockTokenRepo = &MockTokenRepo{
		SaveTokenFn: func(ctx context.Context, token string) error {
			return nil
		},
		ReadTokenByIdxFn: func(ctx context.Context, idx string) (string, error) {
			return savedRefreshToken, nil
		},
		DeleteTokenFn: func(ctx context.Context, idx string) error {
			return nil
		},
	}

	return NewToken(mockTokenRepo)
}
