package user

import (
	"context"

	apiv1 "buf.build/gen/go/wanho/security-proof-api/protocolbuffers/go/api/v1"
	"connectrpc.com/connect"
)

// MockUser struct is used for testing the UserServiceClient structure.
type MockUser struct {
	CreateUserFn    func(ctx context.Context, req *connect.Request[apiv1.CreateUserRequest]) (*connect.Response[apiv1.CreateUserResponse], error)
	ReadUserFn      func(ctx context.Context, req *connect.Request[apiv1.ReadUserRequest]) (*connect.Response[apiv1.ReadUserResponse], error)
	UpdateUserFn    func(ctx context.Context, req *connect.Request[apiv1.UpdateUserRequest]) (*connect.Response[apiv1.UpdateUserResponse], error)
	DeleteUserFn    func(ctx context.Context, req *connect.Request[apiv1.DeleteUserRequest]) (*connect.Response[apiv1.DeleteUserResponse], error)
	ListUserFn      func(ctx context.Context, req *connect.Request[apiv1.ListUserRequest]) (*connect.Response[apiv1.ListUserResponse], error)
	SignInFn        func(ctx context.Context, req *connect.Request[apiv1.SignInRequest]) (*connect.Response[apiv1.SignInResponse], error)
	SignOutFn       func(ctx context.Context, req *connect.Request[apiv1.SignOutRequest]) (*connect.Response[apiv1.SignOutResponse], error)
	RotationTokenFn func(ctx context.Context, req *connect.Request[apiv1.RotationTokenRequest]) (*connect.Response[apiv1.RotationTokenResponse], error)
	DuplicateIDFn   func(ctx context.Context, req *connect.Request[apiv1.DuplicateIDRequest]) (*connect.Response[apiv1.DuplicateIDResponse], error)
	ReadUserByIdxFn func(ctx context.Context, idx int32, accessToken string) (*apiv1.User, error)
}

// CreateUser method is the mock test function for CreateUser.
func (m *MockUser) CreateUser(ctx context.Context, req *connect.Request[apiv1.CreateUserRequest]) (*connect.Response[apiv1.CreateUserResponse], error) {
	return m.CreateUserFn(ctx, req)
}

// ReadUser method is the mock test function for ReadUser.
func (m *MockUser) ReadUser(ctx context.Context, req *connect.Request[apiv1.ReadUserRequest]) (*connect.Response[apiv1.ReadUserResponse], error) {
	return m.ReadUserFn(ctx, req)
}

// UpdateUser method is the mock test function for UpdateUser.
func (m *MockUser) UpdateUser(ctx context.Context, req *connect.Request[apiv1.UpdateUserRequest]) (*connect.Response[apiv1.UpdateUserResponse], error) {
	return m.UpdateUserFn(ctx, req)
}

// DeleteUser method is the mock test function for DeleteUser.
func (m *MockUser) DeleteUser(ctx context.Context, req *connect.Request[apiv1.DeleteUserRequest]) (*connect.Response[apiv1.DeleteUserResponse], error) {
	return m.DeleteUserFn(ctx, req)
}

// ListUser method is the mock test function for ListUser.
func (m *MockUser) ListUser(ctx context.Context, req *connect.Request[apiv1.ListUserRequest]) (*connect.Response[apiv1.ListUserResponse], error) {
	return m.ListUserFn(ctx, req)
}

// SignIn method is the mock test function for SignIn.
func (m *MockUser) SignIn(ctx context.Context, req *connect.Request[apiv1.SignInRequest]) (*connect.Response[apiv1.SignInResponse], error) {
	return m.SignInFn(ctx, req)
}

// SignOut method is the mock test function for SignOut.
func (m *MockUser) SignOut(ctx context.Context, c *connect.Request[apiv1.SignOutRequest]) (*connect.Response[apiv1.SignOutResponse], error) {
	return m.SignOutFn(ctx, c)
}

// RotationToken method is the mock test function for RotationToken.
func (m *MockUser) RotationToken(ctx context.Context, req *connect.Request[apiv1.RotationTokenRequest]) (*connect.Response[apiv1.RotationTokenResponse], error) {
	return m.RotationTokenFn(ctx, req)
}

// DuplicateID method is the mock test function for DuplicateID.
func (m *MockUser) DuplicateID(ctx context.Context, req *connect.Request[apiv1.DuplicateIDRequest]) (*connect.Response[apiv1.DuplicateIDResponse], error) {
	return m.DuplicateIDFn(ctx, req)
}

// ReadUserByIdx method is the mock test function for ReadUserByIdx.
func (m *MockUser) ReadUserByIdx(ctx context.Context, idx int32, accessToken string) (*apiv1.User, error) {
	return m.ReadUserByIdxFn(ctx, idx, accessToken)
}
