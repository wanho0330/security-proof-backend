// Package chain is a package for handling blockchain processes.
package chain

import (
	"context"

	v1 "buf.build/gen/go/wanho/security-proof-api/protocolbuffers/go/chain/v1"
	"connectrpc.com/connect"
)

// MockChain struct is used for testing the ProofServiceClient structure.
type MockChain struct {
	ConfirmProofFn       func(ctx context.Context, req *connect.Request[v1.ConfirmProofRequest]) (*connect.Response[v1.ConfirmProofResponse], error)
	ConfirmUpdateProofFn func(ctx context.Context, req *connect.Request[v1.ConfirmUpdateProofRequest]) (*connect.Response[v1.ConfirmUpdateProofResponse], error)
	ReadImageHashesFn    func(ctx context.Context, req *connect.Request[v1.ReadImageHashesRequest]) (*connect.Response[v1.ReadImageHashesResponse], error)
	ReadLastImageHashFn  func(ctx context.Context, req *connect.Request[v1.ReadLastImageHashRequest]) (*connect.Response[v1.ReadLastImageHashResponse], error)
}

// ConfirmProof method is the mock test function for ConfirmProof.
func (m *MockChain) ConfirmProof(ctx context.Context, req *connect.Request[v1.ConfirmProofRequest]) (*connect.Response[v1.ConfirmProofResponse], error) {
	return m.ConfirmProofFn(ctx, req)
}

// ConfirmUpdateProof method is the mock test function for ConfirmUpdateProof.
func (m *MockChain) ConfirmUpdateProof(ctx context.Context, req *connect.Request[v1.ConfirmUpdateProofRequest]) (*connect.Response[v1.ConfirmUpdateProofResponse], error) {
	return m.ConfirmUpdateProofFn(ctx, req)
}

// ReadImageHashes method is the mock test function for ReadImageHashes.
func (m *MockChain) ReadImageHashes(ctx context.Context, req *connect.Request[v1.ReadImageHashesRequest]) (*connect.Response[v1.ReadImageHashesResponse], error) {
	return m.ReadImageHashesFn(ctx, req)
}

// ReadLastImageHash method is the mock test function for ReadLastImageHash.
func (m *MockChain) ReadLastImageHash(ctx context.Context, req *connect.Request[v1.ReadLastImageHashRequest]) (*connect.Response[v1.ReadLastImageHashResponse], error) {
	return m.ReadLastImageHashFn(ctx, req)
}
