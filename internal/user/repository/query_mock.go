package repository

import (
	"context"

	"security-proof/internal/db/security_proof/user/model"
)

// MockUserQuery struct is used for testing the userQuery structure.
type MockUserQuery struct {
	ReadUserByIdxFn func(ctx context.Context, idx int32) (user *model.User, err error)
	ReadUserByIDFn  func(ctx context.Context, id string) (user *model.User, err error)
	AllUsersFn      func(ctx context.Context) ([]*model.User, error)
	SearchUsersFn   func(ctx context.Context, id string) ([]*model.User, error)
	SignInUserFn    func(ctx context.Context, id string) (user *model.User, err error)
}

// ReadUserByIdx method is the mock test function for ReadUserByIdx.
func (m *MockUserQuery) ReadUserByIdx(ctx context.Context, idx int32) (user *model.User, err error) {
	return m.ReadUserByIdxFn(ctx, idx)
}

// ReadUserByID method is the mock test function for ReadUserByID.
func (m *MockUserQuery) ReadUserByID(ctx context.Context, id string) (user *model.User, err error) {
	return m.ReadUserByIDFn(ctx, id)
}

// AllUsers method is the mock test function for AllUsers.
func (m *MockUserQuery) AllUsers(ctx context.Context) ([]*model.User, error) {
	return m.AllUsersFn(ctx)
}

// SearchUsers method is the mock test function for SearchUsers.
func (m *MockUserQuery) SearchUsers(ctx context.Context, id string) ([]*model.User, error) {
	return m.SearchUsersFn(ctx, id)
}

// SignInUser method is the mock test function for SignInUser.
func (m *MockUserQuery) SignInUser(ctx context.Context, id string) (user *model.User, err error) {
	return m.SignInUserFn(ctx, id)
}
