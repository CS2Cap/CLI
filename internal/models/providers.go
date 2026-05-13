package models

type ProviderInfo struct {
	Key             string            `json:"key"`
	Logo            string            `json:"logo"`
	Code            string            `json:"code"`
	MarketType      string            `json:"market_type"`
	DefaultCurrency string            `json:"default_currency"`
	Fees            ProviderFees      `json:"fees"`
	Features        ProviderFeatures  `json:"features"`
	Health          ProviderHealth    `json:"health"`
}

type ProviderFees struct {
	SellFee            *float64 `json:"sell_fee,omitempty"`
	InstaSellFee       *float64 `json:"insta_sell_fee,omitempty"`
	TradingSpreadFee   *float64 `json:"trading_spread_fee,omitempty"`
}

type ProviderFeatures struct {
	HasBuyOrders   bool `json:"has_buy_orders"`
	HasRecentSales bool `json:"has_recent_sales"`
}

type ProviderHealth struct {
	Status         string  `json:"status"`
	LastCheckedAt  string  `json:"last_checked_at,omitempty"`
	TotalOffers    int     `json:"total_offers,omitempty"`
	UniqueItems    int     `json:"unique_items,omitempty"`
	MarketCoverage float64 `json:"market_coverage,omitempty"`
	TotalValue     int64   `json:"total_value,omitempty"`
	TotalValueUSD  int64   `json:"total_value_usd,omitempty"`
}
