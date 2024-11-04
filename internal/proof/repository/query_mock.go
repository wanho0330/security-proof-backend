package repository

import (
	"golang.org/x/net/context"

	"security-proof/internal/db/security_proof/proof/model"
)

// MockProofQuery struct is used for testing the proofQuery structure.
type MockProofQuery struct {
	ReadProofFn            func(ctx context.Context, idx int32) (*model.Proof, error)
	AllProofsFn            func(ctx context.Context) ([]*model.Proof, error)
	SearchProofsFn         func(ctx context.Context, category string) ([]*model.Proof, error)
	ReadFirstProofImageFn  func(ctx context.Context, idx int32) (proof *model.Proof, err error)
	ReadSecondProofImageFn func(ctx context.Context, idx int32) (proof *model.Proof, err error)
	ReadProofLogFn         func(ctx context.Context, idx int32) (*model.Proof, error)
}

// ReadProof method is the mock test function for ReadProof.
func (m *MockProofQuery) ReadProof(ctx context.Context, idx int32) (*model.Proof, error) {
	return m.ReadProofFn(ctx, idx)
}

// AllProofs method is the mock test function for AllProofs.
func (m *MockProofQuery) AllProofs(ctx context.Context) ([]*model.Proof, error) {
	return m.AllProofsFn(ctx)
}

// SearchProofs method is the mock test function for SearchProofs.
func (m *MockProofQuery) SearchProofs(ctx context.Context, category string) ([]*model.Proof, error) {
	return m.SearchProofsFn(ctx, category)
}

// ReadProofLog method is the mock test function for ReadProofLog.
func (m *MockProofQuery) ReadProofLog(ctx context.Context, idx int32) (*model.Proof, error) {
	return m.ReadProofLogFn(ctx, idx)
}

// ReadFirstProofImage method is the mock test function for ReadFirstProofImage.
func (m *MockProofQuery) ReadFirstProofImage(ctx context.Context, idx int32) (proof *model.Proof, err error) {
	return m.ReadFirstProofImageFn(ctx, idx)
}

// ReadSecondProofImage method is the mock test function for ReadSecondProofImage.
func (m *MockProofQuery) ReadSecondProofImage(ctx context.Context, idx int32) (proof *model.Proof, err error) {
	return m.ReadSecondProofImageFn(ctx, idx)
}
