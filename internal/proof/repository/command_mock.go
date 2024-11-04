package repository

import (
	"context"
	"database/sql"
	"log"

	"security-proof/internal/db/security_proof/proof/model"
)

// MockProofCommand struct is used for testing the proofCommand structure.
type MockProofCommand struct {
	BeginFn              func(ctx context.Context) (*sql.Tx, error)
	CommitFn             func(ctx context.Context, tx *sql.Tx) error
	RollbackFn           func(ctx context.Context, tx *sql.Tx) error
	CreateProofFn        func(ctx context.Context, proof *model.Proof, tx *sql.Tx) (int32, error)
	UpdateProofFn        func(ctx context.Context, proof *model.Proof, tx *sql.Tx) (int32, error)
	UploadProofFn        func(ctx context.Context, proof *model.Proof, tx *sql.Tx) (int32, error)
	DeleteProofFn        func(ctx context.Context, idx int32, tx *sql.Tx) error
	ConfirmProofFn       func(ctx context.Context, proof *model.Proof, tx *sql.Tx) error
	ConfirmUpdateProofFn func(ctx context.Context, proof *model.Proof, tx *sql.Tx) error
}

// Begin method is the mock test function for Begin.
func (m *MockProofCommand) Begin(ctx context.Context) (*sql.Tx, error) {
	if m.BeginFn == nil {
		log.Fatal("mock BeginFn is nil")
	}
	return m.BeginFn(ctx)
}

// Commit method is the mock test function for Commit.
func (m *MockProofCommand) Commit(ctx context.Context, tx *sql.Tx) error {
	if m.CommitFn == nil {
		log.Fatal("mock CommitFn is nil")
	}
	return m.CommitFn(ctx, tx)
}

// Rollback method is the mock test function for Rollback.
func (m *MockProofCommand) Rollback(ctx context.Context, tx *sql.Tx) error {
	if m.RollbackFn == nil {
		log.Fatal("mock RollbackFn is nil")
	}
	return m.RollbackFn(ctx, tx)
}

// CreateProof method is the mock test function for CreateProof.
func (m *MockProofCommand) CreateProof(ctx context.Context, proof *model.Proof, tx *sql.Tx) (int32, error) {
	if m.CreateProofFn == nil {
		log.Fatal("mock CreateProofFn is nil")
	}
	return m.CreateProofFn(ctx, proof, tx)
}

// UpdateProof method is the mock test function for UpdateProof.
func (m *MockProofCommand) UpdateProof(ctx context.Context, proof *model.Proof, tx *sql.Tx) (int32, error) {
	if m.UpdateProofFn == nil {
		log.Fatal("mock UpdateProofFn is nil")
	}
	return m.UpdateProofFn(ctx, proof, tx)
}

// UploadProof method is the mock test function for UploadProof.
func (m *MockProofCommand) UploadProof(ctx context.Context, proof *model.Proof, tx *sql.Tx) (int32, error) {
	if m.UploadProofFn == nil {
		log.Fatal("mock UploadProofFn is nil")
	}
	return m.UploadProofFn(ctx, proof, tx)
}

// DeleteProof method is the mock test function for DeleteProof.
func (m *MockProofCommand) DeleteProof(ctx context.Context, idx int32, tx *sql.Tx) error {
	if m.DeleteProofFn == nil {
		log.Fatal("mock DeleteProofFn is nil")
	}
	return m.DeleteProofFn(ctx, idx, tx)
}

// ConfirmProof method is the mock test function for ConfirmProof.
func (m *MockProofCommand) ConfirmProof(ctx context.Context, proof *model.Proof, tx *sql.Tx) error {
	if m.ConfirmProofFn == nil {
		log.Fatal("mock ConfirmProofFn is nil")
	}
	return m.ConfirmProofFn(ctx, proof, tx)
}

// ConfirmUpdateProof method is the mock test function for ConfirmUpdateProof.
func (m *MockProofCommand) ConfirmUpdateProof(ctx context.Context, proof *model.Proof, tx *sql.Tx) error {
	if m.ConfirmUpdateProofFn == nil {
		log.Fatal("mock ConfirmUpdateProofFn is nil")
	}
	return m.ConfirmUpdateProofFn(ctx, proof, tx)
}
