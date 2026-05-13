package api

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/cs2cap/cli/internal/models"
)

type SearchItemsParams struct {
	Query       string
	ItemType    string
	RarityName  string
	WeaponType  string
	Category    string
	Limit       int
	Offset      int
}

func (c *Client) SearchItems(ctx context.Context, params SearchItemsParams) (*models.ItemsResponse, error) {
	q := url.Values{}
	if params.Query != "" {
		q.Set("q", params.Query)
	}
	if params.ItemType != "" {
		q.Set("item_type", params.ItemType)
	}
	if params.RarityName != "" {
		q.Set("rarity_name", params.RarityName)
	}
	if params.WeaponType != "" {
		q.Set("weapon_type", params.WeaponType)
	}
	if params.Category != "" {
		q.Set("category", params.Category)
	}
	if params.Limit > 0 {
		q.Set("limit", strconv.Itoa(params.Limit))
	}
	if params.Offset > 0 {
		q.Set("offset", strconv.Itoa(params.Offset))
	}

	var resp models.ItemsResponse
	if err := c.get(ctx, "/v1/items", q, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) GetItem(ctx context.Context, itemID int) (*models.ItemOut, error) {
	var resp models.ItemOut
	if err := c.get(ctx, fmt.Sprintf("/v1/items/%d", itemID), nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
