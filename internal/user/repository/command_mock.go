package repository

import (
	"context"
	"database/sql"
	"log"

	"security-proof/internal/db/security_proof/user/model"
)

// MockUserCommand struct is used for testing the userCommand structure.
type MockUserCommand struct {
	BeginFn      func(ctx context.Context) (*sql.Tx, error)
	CommitFn     func(ctx context.Context, tx *sql.Tx) error
	RollbackFn   func(ctx context.Context, tx *sql.Tx) error
	CreateUserFn func(ctx context.Context, user *model.User, tx *sql.Tx) (int32, error)
	UpdateUserFn func(ctx context.Context, user *model.User, tx *sql.Tx) (int32, error)
	DeleteUserFn func(ctx context.Context, idx int32, tx *sql.Tx) error
}

// Begin method is the mock test function for Begin.
func (m *MockUserCommand) Begin(ctx context.Context) (*sql.Tx, error) {
	if m.BeginFn == nil {
		log.Fatal("mock BeginFn is nil")
	}
	return m.BeginFn(ctx)
}

// Commit method is the mock test function for Commit.
func (m *MockUserCommand) Commit(ctx context.Context, tx *sql.Tx) error {
	if m.CommitFn == nil {
		log.Fatal("mock CommitFn is nil")
	}
	return m.CommitFn(ctx, tx)
}

// Rollback method is the mock test function for Rollback.
func (m *MockUserCommand) Rollback(ctx context.Context, tx *sql.Tx) error {
	if m.RollbackFn == nil {
		log.Fatal("mock RollbackFn is nil")
	}
	return m.RollbackFn(ctx, tx)
}

// CreateUser method is the mock test function for CreateUser.
func (m *MockUserCommand) CreateUser(ctx context.Context, user *model.User, tx *sql.Tx) (int32, error) {
	if m.CreateUserFn == nil {
		log.Fatal("mock CreateProofFn is nil")
	}
	return m.CreateUserFn(ctx, user, tx)
}

// UpdateUser method is the mock test function for UpdateUser.
func (m *MockUserCommand) UpdateUser(ctx context.Context, user *model.User, tx *sql.Tx) (int32, error) {
	if m.UpdateUserFn == nil {
		log.Fatal("mock UpdateProofFn is nil")
	}
	return m.UpdateUserFn(ctx, user, tx)
}

// DeleteUser method is the mock test function for DeleteUser.
func (m *MockUserCommand) DeleteUser(ctx context.Context, idx int32, tx *sql.Tx) (err error) {
	if m.DeleteUserFn == nil {
		log.Fatal("mock DeleteProofFn is nil")
	}
	return m.DeleteUserFn(ctx, idx, tx)
}
