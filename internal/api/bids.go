package api

import (
	"context"
	"net/url"
	"strconv"

	"github.com/cs2cap/cli/internal/models"
)

type ListBidsParams struct {
	ItemID         *int
	MarketHashName *string
	Phase          *string
	Providers      []string
	Currency       string
	Limit          int
	Offset         int
}

func (c *Client) ListBids(ctx context.Context, params ListBidsParams) (*models.BidsResponse, error) {
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

	var resp models.BidsResponse
	if err := c.get(ctx, "/v1/bids", q, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) BatchBids(ctx context.Context, req models.BatchBidsRequest) (*models.BatchBidsResponse, error) {
	var resp models.BatchBidsResponse
	if err := c.post(ctx, "/v1/bids/batch", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
