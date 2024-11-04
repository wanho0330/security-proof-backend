// Package service returns data for the controller layer, processing business logic from the repository.
package service

import (
	"context"
	"errors"

	"buf.build/gen/go/wanho/security-proof-api/connectrpc/go/api/v1/apiv1connect"
	apiv1 "buf.build/gen/go/wanho/security-proof-api/protocolbuffers/go/api/v1"
	"connectrpc.com/connect"

	"security-proof/internal/dashboard/convert"
	"security-proof/internal/dashboard/repository"
	"security-proof/internal/db/security_proof/proof/model"
	"security-proof/internal/db/security_proof/proof/table"
	"security-proof/pkg/auth"
	"security-proof/pkg/constants"
	elasticmanage "security-proof/pkg/manage/elastic"
)

var conv = convert.ServiceConverterImpl{}

// DashboardQuery struct is composed of a token, a DashboardQuerier, a user client and an elastic.
type DashboardQuery struct {
	token          *auth.Token
	dashboardQuery repository.DashboardQuerier
	elastic        elasticmanage.Elastic
	user           apiv1connect.UserServiceClient
}

// NewDashboardService function is returning a DashboardQuery interface accepting a token, a DashboardQuerier, a user client and an Elastic.
func NewDashboardService(token *auth.Token, dashboardQuery repository.DashboardQuerier, elastic elasticmanage.Elastic, user apiv1connect.UserServiceClient) *DashboardQuery {
	return &DashboardQuery{token: token, dashboardQuery: dashboardQuery, elastic: elastic, user: user}
}

// ReadDashboard method is returning a NotConfirmProofs, a NotUploadProofs, a CountUploadProofs, an error, accepting a context and access token.
func (q *DashboardQuery) ReadDashboard(ctx context.Context, accessToken string) ([]*apiv1.NotConfirmProof, []*apiv1.NotUploadProof, []*apiv1.CountUploadProof, error) {
	_, _, err := q.token.ValidateToken(accessToken)
	if err != nil {
		return nil, nil, nil, errors.Join(constants.ErrDashboardRead, err)
	}

	notConfirmProof, err := q.dashboardQuery.NotConfirmProof(ctx)
	if err != nil {
		return nil, nil, nil, errors.Join(constants.ErrDashboardRead, err)
	}

	notConfirmResult, err := q.readNotConfirmUser(ctx, accessToken, notConfirmProof)
	if err != nil {
		return nil, nil, nil, errors.Join(constants.ErrDashboardRead, err)
	}

	notUploadProofs, err := q.dashboardQuery.NotUploadProof(ctx)
	if err != nil {
		return nil, nil, nil, errors.Join(constants.ErrDashboardRead, err)
	}

	notUploadResult, err := q.readNotUploadUser(ctx, accessToken, notUploadProofs)
	if err != nil {
		return nil, nil, nil, errors.Join(constants.ErrDashboardRead, err)
	}

	uploadProofs, err := q.elastic.CountExist(ctx, table.Proof.UploadedAt.Name())
	if err != nil {
		return nil, nil, nil, errors.Join(constants.ErrDashboardRead, err)
	}

	countUploadProofs := make([]*apiv1.CountUploadProof, 0)
	countUploadProofs = append(countUploadProofs, &apiv1.CountUploadProof{
		Title: "업로드된 증적",
		Count: uploadProofs,
	})

	allProofs, err := q.elastic.CountAll(ctx)
	if err != nil {
		return nil, nil, nil, errors.Join(constants.ErrDashboardRead, err)
	}
	countUploadProofs = append(countUploadProofs, &apiv1.CountUploadProof{
		Title: "모든 증적",
		Count: allProofs,
	})

	return notConfirmResult, notUploadResult, countUploadProofs, nil
}

func (q *DashboardQuery) readNotConfirmUser(ctx context.Context, accessToken string, notConfirmProofs []*model.Proof) ([]*apiv1.NotConfirmProof, error) {
	notConfirmResult := make([]*apiv1.NotConfirmProof, len(notConfirmProofs))
	for i, proof := range notConfirmProofs {

		notConfirmResult[i] = conv.ModelToNotConfirmProto(proof)

		UploadedUserReq := connect.NewRequest(&apiv1.ReadUserRequest{Idx: *proof.UploadedUserIdx})
		UploadedUserReq.Header().Set("accessToken", accessToken)
		UploadedUser, err := q.user.ReadUser(ctx, UploadedUserReq)
		if err != nil {
			return nil, errors.Join(constants.ErrDashboardRead, err)
		}

		notConfirmResult[i].UploadedUserId = UploadedUser.Msg.User.Id

		CreatedUserReq := connect.NewRequest(&apiv1.ReadUserRequest{Idx: *proof.CreatedUserIdx})
		CreatedUserReq.Header().Set("accessToken", accessToken)
		CreatedUser, err := q.user.ReadUser(ctx, CreatedUserReq)
		if err != nil {
			return nil, errors.Join(constants.ErrDashboardRead, err)
		}

		notConfirmResult[i].CreatedUserId = CreatedUser.Msg.User.Id
	}

	return notConfirmResult, nil
}

func (q *DashboardQuery) readNotUploadUser(ctx context.Context, accessToken string, notUploadProofs []*model.Proof) ([]*apiv1.NotUploadProof, error) {
	notUploadResult := make([]*apiv1.NotUploadProof, len(notUploadProofs))
	for i, proof := range notUploadProofs {
		notUploadResult[i] = conv.ModelToNotUploadProto(proof)

		UploadedUserReq := connect.NewRequest(&apiv1.ReadUserRequest{Idx: *proof.UploadedUserIdx})
		UploadedUserReq.Header().Set("accessToken", accessToken)
		UploadedUser, err := q.user.ReadUser(ctx, UploadedUserReq)
		if err != nil {
			return nil, errors.Join(constants.ErrDashboardRead, err)
		}

		notUploadResult[i].UploadedUserId = UploadedUser.Msg.User.Id

		CreatedUserReq := connect.NewRequest(&apiv1.ReadUserRequest{Idx: *proof.CreatedUserIdx})
		CreatedUserReq.Header().Set("accessToken", accessToken)
		CreatedUser, err := q.user.ReadUser(ctx, CreatedUserReq)
		if err != nil {
			return nil, errors.Join(constants.ErrDashboardRead, err)
		}

		notUploadResult[i].CreatedUserId = CreatedUser.Msg.User.Id
	}

	return notUploadResult, nil
}
