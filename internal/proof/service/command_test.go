package service

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	apiv1 "buf.build/gen/go/wanho/security-proof-api/protocolbuffers/go/api/v1"
	chainv1 "buf.build/gen/go/wanho/security-proof-api/protocolbuffers/go/chain/v1"
	"connectrpc.com/connect"
	"github.com/stretchr/testify/assert"

	"security-proof/internal/db/security_proof/proof/model"
	"security-proof/internal/proof/repository"
	"security-proof/pkg/auth"
	"security-proof/pkg/constants"
	chainmanage "security-proof/pkg/manage/chain"
)

var ctx, cancel = context.WithTimeout(context.Background(), time.Second*5)
var command = newMockCommand()

func TestProofCommand_CreateProof(t *testing.T) {
	defer cancel()

	accessToken, _, err := mockToken.CreateToken(ctx, "1", constants.RoleAdmin)
	assert.NoError(t, err, "토큰 생성 중 에러가 발생하지 않았습니다.")

	t.Run("증적 추가 케이스", func(t *testing.T) {
		proof := &apiv1.Proof{
			Category:    "test",
			Description: "test",
		}

		idx, err := command.CreateProof(ctx, proof, accessToken)
		assert.NoError(t, err, "에러가 발생하지 않았습니다.")
		assert.Equal(t, int32(1), idx, "테스트 증적이 정상적으로 생성되었습니다.")
	})
}

func TestProofCommand_UpdateProof(t *testing.T) {
	defer cancel()

	accessToken, _, err := mockToken.CreateToken(ctx, "1", constants.RoleAdmin)
	assert.NoError(t, err, "토큰 생성 중 에러가 발생하지 않았습니다.")

	t.Run("증적 업데이트 케이스", func(t *testing.T) {
		proof := &apiv1.Proof{Idx: 1}

		idx, err := command.UpdateProof(ctx, proof, accessToken)
		assert.NoError(t, err, "에러가 발생하지 않았습니다.")
		assert.Equal(t, int32(1), idx, "테스트 증적이 정상적으로 업데이트되었습니다.")
	})
}

func TestProofCommand_DeleteProof(t *testing.T) {
	defer cancel()

	accessToken, _, err := mockToken.CreateToken(ctx, "1", constants.RoleAdmin)
	assert.NoError(t, err, "토큰 생성 중 에러가 발생하지 않았습니다.")

	t.Run("증적 삭제 케이스", func(t *testing.T) {
		err := command.DeleteProof(ctx, 1, accessToken)
		assert.NoError(t, err, "에러가 발생하지 않았습니다.")
		assert.Equal(t, err, nil, "테스트 증적이 정상적으로 삭제되었습니다")
	})
}

func TestProofCommand_UploadProof(t *testing.T) {
	defer cancel()

	accessToken, _, err := mockToken.CreateToken(ctx, "1", constants.RoleEngineer)
	assert.NoError(t, err)

	t.Run("증적 업로드 케이스", func(t *testing.T) {
		idx, err := command.UploadProof(ctx, 1, []byte{}, []byte{}, accessToken)
		assert.NoError(t, err, "에러가 발생하지 않았습니다.")
		assert.Equal(t, int32(1), idx, "테스트 증적이 정상적으로 업데이트되었습니다.")
	})
}

func TestProofCommand_ConfirmProof(t *testing.T) {
	defer cancel()

	accessToken, _, err := mockToken.CreateToken(ctx, "1", constants.RoleAdmin)
	assert.NoError(t, err)

	t.Run("증적 확정 케이스", func(t *testing.T) {
		err := command.ConfirmProof(ctx, 1, accessToken)
		assert.NoError(t, err, "에러가 발생하지 않았습니다.")
		assert.Equal(t, err, nil, "테스트 증적이 정상적으로 컨펌되었습니다.")
	})
}

func TestProofCommand_ConfirmUpdateProof(t *testing.T) {
	defer cancel()

	accessToken, _, err := mockToken.CreateToken(ctx, "1", constants.RoleAdmin)
	assert.NoError(t, err)

	t.Run("증적 확정 케이스", func(t *testing.T) {
		err := command.ConfirmUpdateProof(ctx, 1, accessToken)
		assert.NoError(t, err, "에러가 발생하지 않았습니다.")
		assert.Equal(t, err, nil, "테스트 증적이 정상적으로 컨펌되었습니다.")
	})
}

func newMockCommand() *ProofCommand {
	return NewProofCommand(mockToken, mockCommand, mockQuery, mockChainClient)
}

var mockTokenRepo = &auth.MockTokenRepo{
	SaveTokenFn:      func(ctx context.Context, token string) error { return nil },
	ReadTokenByIdxFn: func(ctx context.Context, idx string) (string, error) { return "", nil },
	DeleteTokenFn:    func(ctx context.Context, idx string) error { return nil },
}

var mockToken = auth.NewToken(mockTokenRepo)

var mockCommand = &repository.MockProofCommand{
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
	CreateProofFn: func(ctx context.Context, proof *model.Proof, tx *sql.Tx) (int32, error) {

		if proof.Category != "test" || proof.Description != "test" {
			return 0, constants.ErrProofCreate
		}
		return 1, nil
	},
	UpdateProofFn: func(ctx context.Context, proof *model.Proof, tx *sql.Tx) (int32, error) {
		if proof.Idx == 0 {
			return 0, constants.ErrProofUpdate
		}
		return 1, nil
	},
	UploadProofFn: func(ctx context.Context, proof *model.Proof, tx *sql.Tx) (int32, error) {
		if proof.Idx == 0 {
			return 0, constants.ErrProofUpload
		}
		return 1, nil

	},
	DeleteProofFn: func(ctx context.Context, idx int32, tx *sql.Tx) error {
		if idx == 0 {
			return constants.ErrProofDelete
		}
		return nil
	},
	ConfirmProofFn: func(ctx context.Context, proof *model.Proof, tx *sql.Tx) error {
		if proof.Idx == 0 {
			return constants.ErrProofConfirm
		}
		return nil
	},
	ConfirmUpdateProofFn: func(ctx context.Context, proof *model.Proof, tx *sql.Tx) error {
		if proof.Idx == 0 {
			return constants.ErrProofConfirm
		}
		return nil
	},
}

var mockChainClient = &chainmanage.MockChain{
	ConfirmProofFn: func(ctx context.Context, req *connect.Request[chainv1.ConfirmProofRequest]) (*connect.Response[chainv1.ConfirmProofResponse], error) {
		return connect.NewResponse(&chainv1.ConfirmProofResponse{
			TokenId: 1,
		}), nil
	},
	ConfirmUpdateProofFn: func(ctx context.Context, req *connect.Request[chainv1.ConfirmUpdateProofRequest]) (*connect.Response[chainv1.ConfirmUpdateProofResponse], error) {
		return connect.NewResponse(&chainv1.ConfirmUpdateProofResponse{}), nil
	},
	ReadImageHashesFn: func(ctx context.Context, req *connect.Request[chainv1.ReadImageHashesRequest]) (*connect.Response[chainv1.ReadImageHashesResponse], error) {
		return nil, nil
	},
	ReadLastImageHashFn: func(ctx context.Context, req *connect.Request[chainv1.ReadLastImageHashRequest]) (*connect.Response[chainv1.ReadLastImageHashResponse], error) {
		return nil, nil
	},
}
