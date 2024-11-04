package elastic

import "context"

// MockElastic struct is used for testing the Elastic structure.
type MockElastic struct {
	CountGTFn  func(ctx context.Context, field string) (int32, error)
	CountAllFn func(ctx context.Context) (int32, error)
}

// CountExist method is the mock test function for CountExist.
func (m *MockElastic) CountExist(ctx context.Context, field string) (int32, error) {
	return m.CountGTFn(ctx, field)
}

// CountAll method is the mock test function for CountAll.
func (m *MockElastic) CountAll(ctx context.Context) (int32, error) {
	return m.CountAllFn(ctx)
}
