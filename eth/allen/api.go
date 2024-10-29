package allen

import (
	"context"
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
