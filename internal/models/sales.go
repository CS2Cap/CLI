package models

type SalesResponse struct {
	Meta        SalesMeta              `json:"meta"`
	Items       []SaleRecordDetail      `json:"items"`
	CacheStatus map[string]string       `json:"cache_status"`
}

type SaleRecordDetail struct {
	Date           string   `json:"date"`
	Provider       string   `json:"provider"`
	Price          int      `json:"price"`
	Currency       string   `json:"currency"`
	ItemID         int      `json:"item_id"`
	MarketHashName string   `json:"market_hash_name"`
	Phase          *string  `json:"phase,omitempty"`
	Float          *float64 `json:"float,omitempty"`
	PaintSeed      *int     `json:"paint_seed,omitempty"`
}

type SalesMeta struct {
	Currency         string           `json:"currency"`
	Filters          SalesFiltersMeta `json:"filters"`
	ProvidersQueried []string         `json:"providers_queried"`
	ResultCount      int              `json:"result_count"`
}

type SalesFiltersMeta struct {
	ItemID             *int     `json:"item_id,omitempty"`
	MarketHashName     *string  `json:"market_hash_name,omitempty"`
	Phase              *string  `json:"phase,omitempty"`
	RequestedProviders []string `json:"requested_providers,omitempty"`
	Limit              int      `json:"limit"`
}
