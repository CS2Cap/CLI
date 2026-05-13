package api

import (
	"context"
	"net/url"
	"strconv"

	"github.com/cs2cap/cli/internal/models"
)

type ListSalesParams struct {
	ItemID         *int
	MarketHashName *string
	Providers      []string
	Limit          int
}

func (c *Client) ListSales(ctx context.Context, params ListSalesParams) (*models.SalesResponse, error) {
	q := url.Values{}
	if params.ItemID != nil {
		q.Set("item_id", strconv.Itoa(*params.ItemID))
	}
	if params.MarketHashName != nil {
		q.Set("market_hash_name", *params.MarketHashName)
	}
	for _, p := range params.Providers {
		q.Add("providers", p)
	}
	if params.Limit > 0 {
		q.Set("limit", strconv.Itoa(params.Limit))
	}

	var resp models.SalesResponse
	if err := c.get(ctx, "/v1/sales", q, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
