package models

type SalesResponse struct {
	Meta       SalesMeta          `json:"meta"`
	Items      []RecentSalesProvider `json:"items"`
	Pagination PaginationMeta     `json:"pagination"`
}

type RecentSalesProvider struct {
	Provider string            `json:"provider"`
	Sales    []SaleRecordDetail `json:"sales"`
}

type SaleRecordDetail struct {
	SaleID         string  `json:"sale_id"`
	ItemID         int     `json:"item_id"`
	MarketHashName string  `json:"market_hash_name"`
	Phase          *string `json:"phase,omitempty"`
	Price          int     `json:"price"`
	Currency       string  `json:"currency"`
	SoldAt         string  `json:"sold_at"`
	Age            *int    `json:"age,omitempty"`
}

type SalesMeta struct {
	Filters           SalesFiltersMeta `json:"filters"`
	ReturnedProviders []string         `json:"returned_providers"`
}

type SalesFiltersMeta struct {
	ItemID         *int     `json:"item_id,omitempty"`
	MarketHashName *string  `json:"market_hash_name,omitempty"`
	Providers      []string `json:"providers,omitempty"`
}

type SalesHistoryResponse struct {
	Items      []SaleRecordDetail `json:"items"`
	Pagination PaginationMeta     `json:"pagination"`
}
