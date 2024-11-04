// Package controller is returning the data needed for the dashboard based on user input.
package controller

import (
	"context"
	"errors"

	apiv1 "buf.build/gen/go/wanho/security-proof-api/protocolbuffers/go/api/v1"
	"connectrpc.com/connect"

	goverter "security-proof/internal/user/convert"
	"security-proof/internal/user/service"
	"security-proof/pkg/constants"
)

var conv = goverter.ControllerConverterImpl{}

// UserController struct is composed of command and query from the service layer.
type UserController struct {
	userCommand *service.UserCommand
	userQuery   *service.UserQuery
}

// NewUserController function is returning a UserController struct that accept command and query from the service layer.
func NewUserController(userCommand *service.UserCommand, userQuery *service.UserQuery) *UserController {
	return &UserController{userCommand: userCommand, userQuery: userQuery}
}

// CreateUser method is returning a CreateUserResponse and an error, accepting a context and a CreateUserRequest.
func (c *UserController) CreateUser(ctx context.Context, req *connect.Request[apiv1.CreateUserRequest]) (*connect.Response[apiv1.CreateUserResponse], error) {
	accessToken := req.Header().Get("accessToken")
	idx, err := c.userCommand.CreateUser(ctx, conv.CreateRequestToUser(req.Msg), accessToken)

	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	res := connect.NewResponse(&apiv1.CreateUserResponse{
		Idx: idx,
	})
	return res, nil
}

// ReadUser method is returning a ReadUserResponse and an error, accepting a context and a ReadUserRequest.
func (c *UserController) ReadUser(ctx context.Context, req *connect.Request[apiv1.ReadUserRequest]) (*connect.Response[apiv1.ReadUserResponse], error) {
	accessToken := req.Header().Get("accessToken")
	user, err := c.userQuery.ReadUserByIdx(ctx, req.Msg.Idx, accessToken)

	if err != nil {
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	res := connect.NewResponse(&apiv1.ReadUserResponse{
		User: user,
	})

	return res, nil
}

// UpdateUser method is returning a UpdateUserResponse and an error, accepting a context and a UpdateUserRequest.
func (c *UserController) UpdateUser(ctx context.Context, req *connect.Request[apiv1.UpdateUserRequest]) (*connect.Response[apiv1.UpdateUserResponse], error) {
	accessToken := req.Header().Get("accessToken")
	idx, err := c.userCommand.UpdateUser(ctx, conv.UpdateRequestToUser(req.Msg), accessToken)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	res := connect.NewResponse(&apiv1.UpdateUserResponse{
		Idx: idx,
	})

	return res, nil
}

// DeleteUser method is returning a DeleteUserResponse and an error, accepting a context and a DeleteUserRequest.
func (c *UserController) DeleteUser(ctx context.Context, req *connect.Request[apiv1.DeleteUserRequest]) (*connect.Response[apiv1.DeleteUserResponse], error) {
	accessToken := req.Header().Get("accessToken")
	err := c.userCommand.DeleteUser(ctx, conv.DeleteRequestToUser(req.Msg).Idx, accessToken)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	res := connect.NewResponse(&apiv1.DeleteUserResponse{})
	return res, nil
}

// ListUser method is returning a ListUserResponse and an error, accepting a context and a ListUserRequest.
func (c *UserController) ListUser(ctx context.Context, req *connect.Request[apiv1.ListUserRequest]) (*connect.Response[apiv1.ListUserResponse], error) {
	// TODO : 서버 페이징 작업 필요
	accessToken := req.Header().Get("accessToken")
	users, err := c.userQuery.ListUsers(ctx, req.Msg.Id, accessToken)
	if errors.Is(err, constants.ErrItemNotFound) {
		return nil, connect.NewError(connect.CodeNotFound, err)
	} else if err != nil {
		return nil, connect.NewError(connect.CodeUnknown, err)
	}

	res := connect.NewResponse(&apiv1.ListUserResponse{
		Users: users,
	})

	return res, nil
}

// DuplicateID method is returning a DuplicateIDResponse and an error, accepting a context and a DuplicateIDRequest.
func (c *UserController) DuplicateID(ctx context.Context, req *connect.Request[apiv1.DuplicateIDRequest]) (*connect.Response[apiv1.DuplicateIDResponse], error) {
	err := c.userQuery.DuplicateUserID(ctx, req.Msg.Id)
	if err != nil {
		return nil, connect.NewError(connect.CodeUnknown, err)
	}

	res := connect.NewResponse(&apiv1.DuplicateIDResponse{})

	return res, nil
}

// SignIn method is returning a SignInResponse and an error, accepting a context and a SignInRequest.
func (c *UserController) SignIn(ctx context.Context, req *connect.Request[apiv1.SignInRequest]) (*connect.Response[apiv1.SignInResponse], error) {
	accessToken, refreshToken, err := c.userCommand.SignInUser(ctx, conv.SignInRequestToUser(req.Msg))
	if errors.Is(err, constants.ErrItemNotFound) {
		return nil, connect.NewError(connect.CodeNotFound, err)
	} else if err != nil {
		return nil, connect.NewError(connect.CodeUnknown, err)
	}

	res := connect.NewResponse(&apiv1.SignInResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})

	return res, nil
}

// SignOut method is returning a SignOutResponse and an error, accepting a context and a SignOutRequest.
func (c *UserController) SignOut(ctx context.Context, req *connect.Request[apiv1.SignOutRequest]) (*connect.Response[apiv1.SignOutResponse], error) {
	accessToken := req.Header().Get("accessToken")

	err := c.userCommand.SingOutUser(ctx, accessToken)
	if err != nil {
		return nil, connect.NewError(connect.CodeUnauthenticated, err)
	}

	res := connect.NewResponse(&apiv1.SignOutResponse{})
	return res, nil
}

// RotationToken method is returning a RotationTokenResponse and an error, accepting a context and a RotationTokenRequest.
func (c *UserController) RotationToken(ctx context.Context, req *connect.Request[apiv1.RotationTokenRequest]) (*connect.Response[apiv1.RotationTokenResponse], error) {
	refreshToken := req.Header().Get("refreshToken")

	newAccessToken, newRefreshToken, err := c.userCommand.RotateRefreshToken(ctx, refreshToken)
	if err != nil {
		return nil, connect.NewError(connect.CodeUnauthenticated, err)
	}

	res := connect.NewResponse(&apiv1.RotationTokenResponse{
		AccessToken:  newAccessToken,
		RefreshToken: newRefreshToken,
	})
	return res, nil
}
