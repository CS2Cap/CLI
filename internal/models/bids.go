package models

type BidsResponse struct {
	Meta       BidsMeta       `json:"meta"`
	Items      []BuyOrderItem `json:"items"`
	Pagination PaginationMeta `json:"pagination"`
}

type BuyOrderItem struct {
	ItemID         int     `json:"item_id"`
	MarketHashName string  `json:"market_hash_name"`
	Phase          *string `json:"phase,omitempty"`
	Provider       string  `json:"provider"`
	HighestBid     int     `json:"highest_bid"`
	NumBids        int     `json:"num_bids"`
	Timestamp      *string `json:"timestamp,omitempty"`
	LastUpdated    *string `json:"last_updated,omitempty"`
}

type BidsMeta struct {
	Currency          string        `json:"currency"`
	Filters           BidsFilterMeta `json:"filters"`
	ReturnedProviders []string      `json:"returned_providers"`
}

type BidsFilterMeta struct {
	ItemID         *int     `json:"item_id,omitempty"`
	MarketHashName *string  `json:"market_hash_name,omitempty"`
	Phase          *string  `json:"phase,omitempty"`
	Providers      []string `json:"providers,omitempty"`
	Currency       string   `json:"currency,omitempty"`
}

type BatchBidsRequest struct {
	ItemIDs []int    `json:"item_ids,omitempty"`
	Names   []string `json:"names,omitempty"`
}

type BatchBidsResponse struct {
	Meta          BatchBidsMeta  `json:"meta"`
	Items         []BatchBidItem `json:"items"`
	ItemsNotFound []int          `json:"items_not_found"`
	NamesNotFound []string       `json:"names_not_found,omitempty"`
}

type BatchBidItem struct {
	ItemID         int            `json:"item_id"`
	MarketHashName string         `json:"market_hash_name"`
	Quotes         []BatchBidQuote `json:"quotes"`
}

type BatchBidQuote struct {
	Provider   string `json:"provider"`
	HighestBid int    `json:"highest_bid"`
	NumBids    int    `json:"num_bids"`
}

type BatchBidsMeta struct {
	Currency string `json:"currency"`
}
