// Package controller is returning the data needed for the dashboard based on user input.
package controller

import (
	"context"

	apiv1 "buf.build/gen/go/wanho/security-proof-api/protocolbuffers/go/api/v1"
	"connectrpc.com/connect"

	"security-proof/internal/dashboard/service"
)

// DashboardController struct is composed of query from the service layer.
type DashboardController struct {
	dashboardQuery *service.DashboardQuery
}

// NewDashboardController function is returning a DashboardController struct that accept query from the service layer.
func NewDashboardController(dashboardQuery *service.DashboardQuery) *DashboardController {
	return &DashboardController{dashboardQuery: dashboardQuery}
}

// ReadDashboard method is returning a ReadDashboardResponse and an error, accepting a ReadDashboardRequest and a context.
func (c *DashboardController) ReadDashboard(ctx context.Context, req *connect.Request[apiv1.ReadDashboardRequest]) (*connect.Response[apiv1.ReadDashboardResponse], error) {
	accessToken := req.Header().Get("accessToken")

	notConfirmProofs, notUploadProofs, countUploadProofs, err := c.dashboardQuery.ReadDashboard(ctx, accessToken)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	res := connect.NewResponse(&apiv1.ReadDashboardResponse{
		NotConfirmProofs:  notConfirmProofs,
		NotUploadProofs:   notUploadProofs,
		CountUploadProofs: countUploadProofs,
	})

	return res, nil
}
