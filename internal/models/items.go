package models

type ItemsResponse struct {
	Items      []ItemOut          `json:"items"`
	Pagination ItemsPaginationMeta `json:"pagination"`
}

type ItemOut struct {
	ItemID        *int     `json:"item_id,omitempty"`
	MarketHashName string  `json:"market_hash_name"`
	Phase         *string `json:"phase,omitempty"`
	ItemType      *string `json:"item_type,omitempty"`
	ItemSubtype   *string `json:"item_subtype,omitempty"`
	WeaponType    *string `json:"weapon_type,omitempty"`
	BaseName      *string `json:"base_name,omitempty"`
	SkinName      *string `json:"skin_name,omitempty"`
	WearName      *string `json:"wear_name,omitempty"`
	DefIndex      *string `json:"def_index,omitempty"`
	PaintIndex    *int    `json:"paint_index,omitempty"`
	Collection    *string `json:"collection,omitempty"`
	Crates        []string `json:"crates,omitempty"`
	RarityName    *string `json:"rarity_name,omitempty"`
	RarityColor   *string `json:"rarity_color,omitempty"`
	StyleName     *string `json:"style_name,omitempty"`
	IsStatTrak    *bool   `json:"is_stattrak,omitempty"`
	IsSouvenir    *bool   `json:"is_souvenir,omitempty"`
	MinFloat      *float64 `json:"min_float,omitempty"`
	MaxFloat      *float64 `json:"max_float,omitempty"`
	ImageURL      *string `json:"image_url,omitempty"`
	Supply        *int    `json:"supply,omitempty"`
}

type ItemsPaginationMeta struct {
	Limit   int  `json:"limit"`
	Offset  int  `json:"offset"`
	Total   int  `json:"total"`
	HasNext bool `json:"has_next"`
	HasPrev bool `json:"has_prev"`
}

type ItemsFilterMetadata struct {
	ItemTypes   []string `json:"item_types,omitempty"`
	RarityNames []string `json:"rarity_names,omitempty"`
	WeaponTypes []string `json:"weapon_types,omitempty"`
	Collections []string `json:"collections,omitempty"`
	Categories  []string `json:"categories,omitempty"`
}

type ItemsMetadataResponse struct {
	Filters ItemsFilterMetadata `json:"filters"`
}
