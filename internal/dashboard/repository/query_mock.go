package repository

import (
	"context"

	"security-proof/internal/db/security_proof/proof/model"
)

// MockDashboardQuery struct is used for testing the dashboardQuery structure.
type MockDashboardQuery struct {
	NotConfirmProofFn func(ctx context.Context) ([]*model.Proof, error)
	NotUploadProofFn  func(ctx context.Context) ([]*model.Proof, error)
}

// NotConfirmProof method is the mock test function for NotConfirmProof.
func (m *MockDashboardQuery) NotConfirmProof(ctx context.Context) ([]*model.Proof, error) {
	return m.NotConfirmProofFn(ctx)
}

// NotUploadProof method is the mock test function for NotUploadProof.
func (m *MockDashboardQuery) NotUploadProof(ctx context.Context) ([]*model.Proof, error) {
	return m.NotUploadProofFn(ctx)
}
