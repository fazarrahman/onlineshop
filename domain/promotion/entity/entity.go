package entity

// Promotion ...
type Promotion struct {
	ID                int64
	ApplyToSKU        string
	MinRequiredQty    int64
	Description       *string
	FreeItemSKU       *string
	NewPriceQty       *int64
	DiscountInPercent *int64
}
