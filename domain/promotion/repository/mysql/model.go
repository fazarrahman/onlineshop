package mysql

// Promotion ...
type Promotion struct {
	ID                int64   `db:"id"`
	ApplyToSKU        string  `db:"apply_to_sku"`
	MinRequiredQty    int64   `db:"min_required_qty"`
	Description       *string `db:"description"`
	FreeItemSKU       *string `db:"free_item_sku"`
	NewPriceQty       *int64  `db:"new_price_qty"`
	DiscountInPercent *int64  `db:"discount_in_percent"`
}
