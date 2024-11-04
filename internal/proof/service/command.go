// Package service returns data for the controller layer, processing business logic from the repository.
package service

import (
	"context"
	"errors"
	"strconv"
	"time"

	"buf.build/gen/go/wanho/security-proof-api/connectrpc/go/chain/v1/chainv1connect"
	apiv1 "buf.build/gen/go/wanho/security-proof-api/protocolbuffers/go/api/v1"
	chainv1 "buf.build/gen/go/wanho/security-proof-api/protocolbuffers/go/chain/v1"
	"connectrpc.com/connect"

	"security-proof/internal/proof/convert"
	"security-proof/internal/proof/repository"
	"security-proof/pkg/auth"
	"security-proof/pkg/constants"
	filemanage "security-proof/pkg/manage/file"
)

var conv = convert.ServiceConverterImpl{}

// ProofCommand struct is composed of a Token, a ProofCommander, a ProofQuerier and a chain.
type ProofCommand struct {
	token        *auth.Token
	proofCommand repository.ProofCommander
	proofQuery   repository.ProofQuerier
	chain        chainv1connect.ProofServiceClient
}

// NewProofCommand function is returning a ProofCommand interface, accepting a Token, a ProofCommander, a ProofQuerier and ProofServiceClient.
func NewProofCommand(
	token *auth.Token,
	proofCommander repository.ProofCommander,
	proofQuerier repository.ProofQuerier,
	chain chainv1connect.ProofServiceClient,
) *ProofCommand {
	return &ProofCommand{
		token:        token,
		proofCommand: proofCommander,
		proofQuery:   proofQuerier,
		chain:        chain,
	}
}

// CreateProof method is returning a created index and an error, accepting a context, a Proof and an access token.
func (c *ProofCommand) CreateProof(ctx context.Context, proof *apiv1.Proof, accessToken string) (int32, error) {
	userIdx, role, err := c.token.ValidateToken(accessToken)
	if err != nil {
		return 0, errors.Join(constants.ErrProofCreate, err)
	}

	if role != constants.RoleAdmin {
		return 0, errors.Join(constants.ErrProofCreate, constants.ErrTokenRoleAuth)
	}

	proof.CreatedUserIdx = auth.StrToInt32(userIdx)
	proof.CreatedAt = convert.TimeToPTimestamppb(time.Now())
	proof.UpdatedAt = convert.TimeToPTimestamppb(time.Now())

	idx, err := c.proofCommand.CreateProof(ctx, conv.ProtoToModel(proof), nil)
	if err != nil {
		return 0, errors.Join(constants.ErrProofCreate, err)
	}
	return idx, nil
}

// UpdateProof method is returning an updated index and an error, accepting a context, a Proof and an access token.
func (c *ProofCommand) UpdateProof(ctx context.Context, proof *apiv1.Proof, accessToken string) (int32, error) {
	userIdx, role, err := c.token.ValidateToken(accessToken)
	if err != nil {
		return 0, errors.Join(constants.ErrProofUpdate, err)
	}

	if role != constants.RoleAdmin {
		return 0, errors.Join(constants.ErrProofUpdate, constants.ErrTokenRoleAuth)
	}

	proof.UpdatedUserIdx = auth.StrToInt32(userIdx)
	proof.UpdatedAt = convert.TimeToPTimestamppb(time.Now())

	idx, err := c.proofCommand.UpdateProof(ctx, conv.ProtoToModel(proof), nil)
	if err != nil {
		return 0, errors.Join(constants.ErrProofUpdate, err)
	}
	return idx, nil
}

// DeleteProof method is returning an error, accepting a context, a deleting idx and an access token.
func (c *ProofCommand) DeleteProof(ctx context.Context, idx int32, accessToken string) error {
	userIdx, role, err := c.token.ValidateToken(accessToken)
	if err != nil {
		return errors.Join(constants.ErrProofDelete, err)
	}

	readProof, err := c.proofQuery.ReadProof(ctx, idx)
	if err != nil {
		return errors.Join(constants.ErrProofDelete, err)
	}

	if auth.Pint32ToStr(readProof.CreatedUserIdx) != userIdx && auth.Pint32ToStr(readProof.UpdatedUserIdx) != userIdx {
		return errors.Join(constants.ErrProofDelete, constants.ErrTokenRoleAuth)
	}

	if role != constants.RoleAdmin {
		return errors.Join(constants.ErrProofCreate, constants.ErrTokenRoleAuth)
	}

	_, err = c.proofQuery.ReadProof(ctx, idx)
	if err != nil {
		return errors.Join(constants.ErrProofDelete, err)
	}

	err = c.proofCommand.DeleteProof(ctx, idx, nil)
	if err != nil {
		return errors.Join(constants.ErrProofDelete, err)
	}

	return nil
}

// UploadProof method is returning an uploaded index and an error, accepting context, an uploading index, a first image byte, a second image byte and access token.
func (c *ProofCommand) UploadProof(ctx context.Context, idx int32, firstImage []byte, secondImage []byte, accessToken string) (int32, error) {
	userIdx, role, err := c.token.ValidateToken(accessToken)
	if err != nil {
		return 0, errors.Join(constants.ErrProofUpload, err)
	}

	readProof, err := c.proofQuery.ReadProof(ctx, idx)
	if err != nil {
		return 0, errors.Join(constants.ErrProofUpload, err)
	}

	if auth.Pint32ToStr(readProof.UploadedUserIdx) != userIdx {
		return 0, errors.Join(constants.ErrProofUpload, constants.ErrTokenRoleAuth)
	}

	if role != constants.RoleEngineer {
		return 0, errors.Join(constants.ErrProofUpload, constants.ErrTokenRoleAuth)
	}

	fileName := strconv.Itoa(int(idx)) + "_" + strconv.FormatInt(time.Now().Unix(), 10) + "_"

	firstImagePath, err := filemanage.SaveFile(fileName+"1", firstImage)
	if err != nil {
		return 0, errors.Join(constants.ErrProofUpload, err)
	}

	secondImagePath, err := filemanage.SaveFile(fileName+"2", secondImage)
	if err != nil {
		return 0, errors.Join(constants.ErrProofUpload, err)
	}

	uploadedUserIdx64, err := strconv.ParseInt(userIdx, 10, 32)
	if err != nil {
		return 0, errors.Join(constants.ErrProofUpload, err)
	}
	proof := &apiv1.Proof{
		Idx:             idx,
		FirstImagePath:  firstImagePath,
		SecondImagePath: secondImagePath,
		UploadedUserIdx: int32(uploadedUserIdx64),
		UploadedAt:      convert.TimeToPTimestamppb(time.Now()),
		Confirm:         constants.NotConfirm,
	}

	_, err = c.proofCommand.UploadProof(ctx, conv.ProtoToModel(proof), nil)

	if err != nil {
		return 0, errors.Join(constants.ErrProofUpload, err)
	}

	return idx, nil
}

// ConfirmProof method is returning an error accepting a context, a confirmed index and access token.
func (c *ProofCommand) ConfirmProof(ctx context.Context, idx int32, accessToken string) error {
	userIdx, role, err := c.token.ValidateToken(accessToken)
	if err != nil {
		return errors.Join(constants.ErrProofConfirm, err)
	}

	readFirstProofImage, err := c.proofQuery.ReadFirstProofImage(ctx, idx)
	if err != nil {
		return errors.Join(constants.ErrProofConfirm, err)
	}

	readSecondProofImage, err := c.proofQuery.ReadSecondProofImage(ctx, idx)
	if err != nil {
		return errors.Join(constants.ErrProofConfirm, err)
	}

	if role != constants.RoleAdmin {
		return errors.Join(constants.ErrProofConfirm, constants.ErrTokenRoleAuth)
	}

	firstImageHash, err := filemanage.ImageToHash(readFirstProofImage.FirstImagePath)
	if err != nil {
		return errors.Join(constants.ErrProofConfirm, err)
	}

	secondImageHash, err := filemanage.ImageToHash(readSecondProofImage.SecondImagePath)
	if err != nil {
		return errors.Join(constants.ErrProofConfirm, err)
	}

	res, err := c.chain.ConfirmProof(
		ctx,
		connect.NewRequest(&chainv1.ConfirmProofRequest{
			Idx:             readFirstProofImage.Idx,
			FirstImageHash:  firstImageHash,
			SecondImageHash: secondImageHash,
		}),
	)

	if err != nil {
		return errors.Join(constants.ErrProofConfirm, err)
	}

	updatedUserIdx64, err := strconv.ParseInt(userIdx, 10, 32)
	if err != nil {
		return errors.Join(constants.ErrProofConfirm, err)
	}

	proof := &apiv1.Proof{
		Idx:            idx,
		UpdatedUserIdx: int32(updatedUserIdx64),
		UpdatedAt:      convert.TimeToPTimestamppb(time.Now()),
		Confirm:        constants.Confirm,
		TokenId:        res.Msg.TokenId,
	}

	err = c.proofCommand.ConfirmProof(ctx, conv.ProtoToModel(proof), nil)
	if err != nil {
		return errors.Join(constants.ErrProofConfirm, err)
	}

	return nil
}

// ConfirmUpdateProof method is returning an error accepting a context, a confirming index and an access token.
func (c *ProofCommand) ConfirmUpdateProof(ctx context.Context, idx int32, accessToken string) error {
	userIdx, role, err := c.token.ValidateToken(accessToken)
	if err != nil {
		return errors.Join(constants.ErrProofUpdateConfirm, err)
	}

	readProof, err := c.proofQuery.ReadProof(ctx, idx)
	if err != nil {
		return errors.Join(constants.ErrProofUpdateConfirm, err)
	}

	if role != constants.RoleAdmin {
		return errors.Join(constants.ErrProofUpdateConfirm, constants.ErrTokenRoleAuth)
	}

	firstImageHash, err := filemanage.ImageToHash(readProof.FirstImagePath)
	if err != nil {
		return errors.Join(constants.ErrProofConfirm, err)
	}

	secondImageHash, err := filemanage.ImageToHash(readProof.SecondImagePath)
	if err != nil {
		return errors.Join(constants.ErrProofUpdateConfirm, err)
	}

	_, err = c.chain.ConfirmUpdateProof(
		ctx,
		connect.NewRequest(&chainv1.ConfirmUpdateProofRequest{
			TokenId:         *readProof.TokenID,
			FirstImageHash:  firstImageHash,
			SecondImageHash: secondImageHash,
		}),
	)
	if err != nil {
		return errors.Join(constants.ErrProofUpdateConfirm, err)
	}

	updatedUserIdx64, err := strconv.ParseInt(userIdx, 10, 32)
	if err != nil {
		return errors.Join(constants.ErrProofUpdateConfirm, err)
	}

	proof := &apiv1.Proof{
		Idx:            idx,
		UpdatedUserIdx: int32(updatedUserIdx64),
		UpdatedAt:      convert.TimeToPTimestamppb(time.Now()),
		Confirm:        constants.Confirm,
	}

	err = c.proofCommand.ConfirmUpdateProof(ctx, conv.ProtoToModel(proof), nil)
	if err != nil {
		return errors.Join(constants.ErrProofUpdateConfirm, err)
	}

	return nil
}
