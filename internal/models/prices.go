package models

type PricesResponse struct {
	Meta       PricesMeta    `json:"meta"`
	Items      []MarketItem  `json:"items"`
	Pagination PaginationMeta `json:"pagination"`
}

type MarketItem struct {
	Provider       string  `json:"provider"`
	ItemID         int     `json:"item_id"`
	MarketHashName string  `json:"market_hash_name"`
	Phase          *string `json:"phase,omitempty"`
	LowestAsk      int     `json:"lowest_ask"`
	Quantity       int     `json:"quantity"`
	Link           *string `json:"link,omitempty"`
	URL            *string `json:"url,omitempty"`
	Timestamp      *string `json:"timestamp,omitempty"`
	LastUpdated    *string `json:"last_updated,omitempty"`
}

type PricesMeta struct {
	Currency          string           `json:"currency"`
	Filters           PricesFilterMeta `json:"filters"`
	ReturnedProviders []string         `json:"returned_providers"`
}

type PricesFilterMeta struct {
	ItemID         *int     `json:"item_id,omitempty"`
	MarketHashName *string  `json:"market_hash_name,omitempty"`
	Phase          *string  `json:"phase,omitempty"`
	Providers      []string `json:"providers,omitempty"`
	Currency       string   `json:"currency,omitempty"`
}

type BatchPricesRequest struct {
	ItemIDs []int    `json:"item_ids,omitempty"`
	Names   []string `json:"names,omitempty"`
}

type BatchPricesResponse struct {
	Meta          BatchPricesMeta  `json:"meta"`
	Items         []BatchPriceItem `json:"items"`
	ItemsNotFound []int            `json:"items_not_found"`
	NamesNotFound []string         `json:"names_not_found,omitempty"`
}

type BatchPriceItem struct {
	ItemID         int               `json:"item_id"`
	MarketHashName string            `json:"market_hash_name"`
	Quotes         []BatchPriceQuote `json:"quotes"`
}

type BatchPriceQuote struct {
	Provider  string `json:"provider"`
	LowestAsk int    `json:"lowest_ask"`
	Quantity  int    `json:"quantity"`
}

type BatchPricesMeta struct {
	Currency string `json:"currency"`
}
