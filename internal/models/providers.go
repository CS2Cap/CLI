package models

type ProviderInfo struct {
	Key       string           `json:"key"`
	Name      string           `json:"name"`
	Website   string           `json:"website"`
	Icon      string           `json:"icon"`
	Enabled   bool             `json:"enabled"`
	HasBids   bool             `json:"has_bids"`
	HasSales  bool             `json:"has_sales"`
	HasDirect bool             `json:"has_direct"`
	IsFree    bool             `json:"is_free"`
	Features  ProviderFeatures `json:"features"`
}

type ProviderFeatures struct {
	DirectURL     bool `json:"direct_url"`
	BuyOrders     bool `json:"buy_orders"`
	SalesHistory  bool `json:"sales_history"`
}

type AllProviders struct {
	Providers map[string]ProviderInfo `json:"providers"`
}
