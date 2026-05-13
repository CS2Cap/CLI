package api

import (
	"context"

	"github.com/cs2cap/cli/internal/models"
)

func (c *Client) ListProviders(ctx context.Context) (map[string]models.ProviderInfo, error) {
	var resp map[string]models.ProviderInfo
	if err := c.get(ctx, "/v1/providers", nil, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}
