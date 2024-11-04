package elastic

import (
	"context"
	"errors"

	"github.com/elastic/go-elasticsearch/v8/typedapi/core/count"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"

	"security-proof/pkg/constants"
)

// Elastic interface is defining data related to managing elasticsearch.
type Elastic interface {
	CountExist(ctx context.Context, field string) (int32, error)
	CountAll(ctx context.Context) (int32, error)
}

// CountExist method is returning a count and an error, accepting a context and a field.
func (e *elastic) CountExist(ctx context.Context, field string) (int32, error) {
	req := &count.Request{
		Query: &types.Query{
			Exists: &types.ExistsQuery{
				Field: field,
			},
		},
	}
	res, err := e.client.Count().Index(Index).
		Request(req).
		Do(ctx)

	if err != nil {
		return 0, errors.Join(constants.ErrElasticCountExist, err)
	}

	return int32(res.Count), nil
}

// CountAll method is returning a count and an error, accepting a context.
func (e *elastic) CountAll(ctx context.Context) (int32, error) {
	req := &count.Request{
		Query: &types.Query{
			MatchAll: &types.MatchAllQuery{},
		},
	}
	res, err := e.client.Count().Index(Index).
		Request(req).
		Do(ctx)

	if err != nil {
		return 0, errors.Join(constants.ErrElasticCountAll, err)
	}

	return int32(res.Count), nil
}
