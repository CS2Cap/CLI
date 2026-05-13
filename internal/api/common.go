package api

import "github.com/cs2cap/cli/internal/models"

type BatchParams struct {
	ItemIDs []int
	Names   []string
}

func (p BatchParams) ToBatchPrices() models.BatchPricesRequest {
	return models.BatchPricesRequest{
		ItemIDs: p.ItemIDs,
		Names:   p.Names,
	}
}

func (p BatchParams) ToBatchBids() models.BatchBidsRequest {
	return models.BatchBidsRequest{
		ItemIDs: p.ItemIDs,
		Names:   p.Names,
	}
}
