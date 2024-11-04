package service

import (
	"context"
	"errors"

	"buf.build/gen/go/wanho/security-proof-api/connectrpc/go/api/v1/apiv1connect"
	apiv1 "buf.build/gen/go/wanho/security-proof-api/protocolbuffers/go/api/v1"
	"connectrpc.com/connect"

	"security-proof/internal/db/security_proof/proof/model"
	"security-proof/internal/proof/repository"
	"security-proof/pkg/auth"
	"security-proof/pkg/constants"
)

// ProofQuery struct is composed of a Token, a ProofQuerier and an UserServiceClient.
type ProofQuery struct {
	token      *auth.Token
	proofQuery repository.ProofQuerier
	user       apiv1connect.UserServiceClient
}

// NewProofQuery function is returning a ProofQuery accepting a Token, a ProofQuerier and an UserServiceClient.
func NewProofQuery(token *auth.Token, proofQuery repository.ProofQuerier, user apiv1connect.UserServiceClient) *ProofQuery {
	return &ProofQuery{token: token, proofQuery: proofQuery, user: user}
}

// ReadProof method is returning a Proof and an error, accepting a context, a reading index and an access token.
func (q *ProofQuery) ReadProof(ctx context.Context, idx int32, accessToken string) (*apiv1.Proof, error) {
	_, _, err := q.token.ValidateToken(accessToken)
	if err != nil {
		return nil, errors.Join(constants.ErrProofRead, err)
	}

	proof, err := q.proofQuery.ReadProof(ctx, idx)
	if err != nil {
		return nil, errors.Join(constants.ErrProofRead, err)
	}

	req := connect.NewRequest(&apiv1.ReadUserRequest{Idx: *proof.UploadedUserIdx})
	req.Header().Set("accessToken", accessToken)
	uploadUser, err := q.user.ReadUser(ctx, req)

	if err != nil {
		return nil, errors.Join(constants.ErrProofRead, err)
	}

	result := conv.ModelToProto(proof)
	result.UploadedUserId = uploadUser.Msg.User.Id
	return result, nil
}

// ReadFirstProofImage method is returning a first image path and an error, accepting a context, a reading index and an access token.
func (q *ProofQuery) ReadFirstProofImage(ctx context.Context, idx int32, accessToken string) (string, error) {
	_, _, err := q.token.ValidateToken(accessToken)
	if err != nil {
		return "", errors.Join(constants.ErrProofReadFirstImage, err)
	}

	proof, err := q.proofQuery.ReadFirstProofImage(ctx, idx)
	if err != nil {
		return "", errors.Join(constants.ErrProofReadFirstImage, err)
	}

	return *proof.FirstImagePath, nil
}

// ReadSecondProofImage method is returning a first image path and an error, accepting a context, a reading index and an access token.
func (q *ProofQuery) ReadSecondProofImage(ctx context.Context, idx int32, accessToken string) (string, error) {
	_, _, err := q.token.ValidateToken(accessToken)
	if err != nil {
		return "", errors.Join(constants.ErrProofReadSecondImage, err)
	}

	proof, err := q.proofQuery.ReadSecondProofImage(ctx, idx)
	if err != nil {
		return "", errors.Join(constants.ErrProofReadSecondImage, err)
	}

	return *proof.SecondImagePath, nil
}

// ReadProofLog method is returning a proof and an error, accepting a context, a reading index and an access token.
func (q *ProofQuery) ReadProofLog(ctx context.Context, idx int32, accessToken string) (*apiv1.Proof, error) {
	_, _, err := q.token.ValidateToken(accessToken)
	if err != nil {
		return nil, errors.Join(constants.ErrProofReadLog, err)
	}

	proof, err := q.proofQuery.ReadProofLog(ctx, idx)
	if err != nil {
		return nil, errors.Join(constants.ErrProofReadLog, err)
	}
	return conv.ModelToProto(proof), nil
}

// ListProofs method is returning proofs and an error, accepting a context, a category and an access token.
func (q *ProofQuery) ListProofs(ctx context.Context, category string, accessToken string) ([]*apiv1.Proof, error) {
	_, _, err := q.token.ValidateToken(accessToken)
	if err != nil {
		return nil, errors.Join(constants.ErrProofList, err)
	}

	var proofs []*model.Proof
	if category == "" {
		proofs, err = q.proofQuery.AllProofs(ctx)
	} else {
		proofs, err = q.proofQuery.SearchProofs(ctx, category)
	}

	if err != nil {
		return nil, errors.Join(constants.ErrProofList, err)
	}

	result := make([]*apiv1.Proof, len(proofs))
	for i, proof := range proofs {
		result[i] = conv.ModelToProto(proof)
	}

	return result, nil
}
