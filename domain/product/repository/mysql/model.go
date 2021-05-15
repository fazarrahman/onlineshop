package mysql

// Product ...
type Product struct {
	SKU   string  `db:"sku"`
	Name  string  `db:"name"`
	Price float64 `db:"price"`
	Qty   int64   `db:"qty"`
}
