package service

import (
	"context"
	"fmt"
	"testing"

	apiv1 "buf.build/gen/go/wanho/security-proof-api/protocolbuffers/go/api/v1"

	"security-proof/internal/db/security_proof/proof/model"
	"security-proof/internal/proof/repository"
	"security-proof/pkg/constants"
	usermanage "security-proof/pkg/manage/user"
)

var query = newMockQuery()

func TestProofQuery_ReadProof(t *testing.T) {

}

func TestProofQuery_ReadFirstProofImage(t *testing.T) {

}

func TestProofQuery_ReadSecondProofImage(t *testing.T) {

}

func TestProofQuery_ReadProofLog(t *testing.T) {

}

func TestProofQuery_ListProofs(t *testing.T) {

}

func TestProofQuery_ReadUserByIdx(t *testing.T) {

}

func newMockQuery() *ProofQuery {
	return NewProofQuery(mockToken, mockQuery, mockUserClient)
}

var mockQuery = &repository.MockProofQuery{
	ReadProofFn: func(ctx context.Context, idx int32) (*model.Proof, error) {

		i := int32(1)

		if idx != 1 {
			return nil, constants.ErrProofRead
		}
		return &model.Proof{
			Idx:             1,
			CreatedUserIdx:  &i,
			UploadedUserIdx: &i,
			UpdatedUserIdx:  &i,
			TokenID:         &i,
		}, nil
	},
	AllProofsFn: func(ctx context.Context) ([]*model.Proof, error) {
		return nil, nil
	},
	SearchProofsFn: func(ctx context.Context, category string) ([]*model.Proof, error) {
		return nil, nil
	},
	ReadFirstProofImageFn: func(ctx context.Context, idx int32) (proof *model.Proof, err error) {
		return nil, nil
	},
	ReadSecondProofImageFn: func(ctx context.Context, idx int32) (proof *model.Proof, err error) {
		return nil, nil
	},
	ReadProofLogFn: func(ctx context.Context, idx int32) (*model.Proof, error) {
		return nil, nil
	},
}

var mockUserClient = &usermanage.MockUser{
	ReadUserByIdxFn: func(ctx context.Context, idx int32, accessToken string) (*apiv1.User, error) {
		fmt.Printf("Mock 유저 불러오기")
		return nil, nil
	},
}
