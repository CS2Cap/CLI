package api

import (
	"context"
	"net/url"
	"strconv"

	"github.com/cs2cap/cli/internal/models"
)

type ListPricesParams struct {
	ItemID         *int
	MarketHashName *string
	Phase          *string
	Providers      []string
	Currency       string
	Limit          int
	Offset         int
}

func (c *Client) ListPrices(ctx context.Context, params ListPricesParams) (*models.PricesResponse, error) {
	q := url.Values{}
	if params.ItemID != nil {
		q.Set("item_id", strconv.Itoa(*params.ItemID))
	}
	if params.MarketHashName != nil {
		q.Set("market_hash_name", *params.MarketHashName)
	}
	if params.Phase != nil {
		q.Set("phase", *params.Phase)
	}
	for _, p := range params.Providers {
		q.Add("providers", p)
	}
	if params.Currency != "" {
		q.Set("currency", params.Currency)
	}
	if params.Limit > 0 {
		q.Set("limit", strconv.Itoa(params.Limit))
	}
	if params.Offset > 0 {
		q.Set("offset", strconv.Itoa(params.Offset))
	}

	var resp models.PricesResponse
	if err := c.get(ctx, "/v1/prices", q, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) BatchPrices(ctx context.Context, req models.BatchPricesRequest) (*models.BatchPricesResponse, error) {
	var resp models.BatchPricesResponse
	if err := c.post(ctx, "/v1/prices/batch", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
