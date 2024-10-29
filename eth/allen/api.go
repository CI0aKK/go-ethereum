package allen

import (
	"context"
	"github.com/ethereum/go-ethereum/core/cache"
)

// API is the struct that defines your custom API.
type API struct{}

// NewAPI returns a new instance of the custom API.
func NewAPI() *API {
	return &API{}
}

// Mikasa test api
func (api *API) Mikasa(_ context.Context, input string) (string, error) {
	return "Hello " + input, nil
}

// GetMempoolTxByHash get mempool tx
func (api *API) GetMempoolTxByHash(ctx context.Context, hash string) (string, error) {
	res := cache.RedisCli.Get(ctx, cache.KeyNamePrefixHash+hash).Val()
	return res, nil
}
