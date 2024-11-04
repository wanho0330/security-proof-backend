package auth

import "context"

// MockTokenRepo struct is used for testing the tokenRepo structure.
type MockTokenRepo struct {
	SaveTokenFn      func(ctx context.Context, token string) error
	ReadTokenByIdxFn func(ctx context.Context, idx string) (string, error)
	DeleteTokenFn    func(ctx context.Context, idx string) error
}

// SaveToken method is the mock test function for SaveToken.
func (m *MockTokenRepo) SaveToken(ctx context.Context, token string) error {
	return m.SaveTokenFn(ctx, token)
}

// ReadTokenByIdx method is the mock test function for ReadTokenByIdx.
func (m *MockTokenRepo) ReadTokenByIdx(ctx context.Context, idx string) (string, error) {
	return m.ReadTokenByIdxFn(ctx, idx)
}

// DeleteToken method is the mock test function for DeleteToken.
func (m *MockTokenRepo) DeleteToken(ctx context.Context, idx string) error {
	return m.DeleteTokenFn(ctx, idx)
}
