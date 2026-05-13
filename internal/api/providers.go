package api

import (
	"context"

	"github.com/cs2cap/cli/internal/models"
)

func (c *Client) ListProviders(ctx context.Context) (*models.AllProviders, error) {
	var resp models.AllProviders
	if err := c.get(ctx, "/v1/providers", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
