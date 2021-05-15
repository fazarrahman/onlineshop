package mysql

import (
	"context"
	"database/sql"

	"github.com/fazarrahman/onlineshop/domain/promotion/entity"
	"github.com/jmoiron/sqlx"
)

// MySQL ...
type MySQL struct {
	db *sqlx.DB
}

// New ...
func New(db *sqlx.DB) *MySQL {
	return &MySQL{
		db: db,
	}
}

// GetAll ...
func (m *MySQL) GetAll(ctx context.Context) ([]*entity.Promotion, error) {
	var (
		promotions []*Promotion
		query      = `
			SELECT
				id,
				apply_to_sku,
				min_required_qty,
				description,
				free_item_sku,
				new_price_qty,
				discount_in_percent
			FROM
				promotion
			WHERE
				status=1 ;
		`
	)

	err := m.db.SelectContext(ctx, &promotions, query)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	var entities []*entity.Promotion
	for _, p := range promotions {
		entities = append(entities, &entity.Promotion{
			ID:                p.ID,
			ApplyToSKU:        p.ApplyToSKU,
			MinRequiredQty:    p.MinRequiredQty,
			Description:       p.Description,
			FreeItemSKU:       p.FreeItemSKU,
			NewPriceQty:       p.NewPriceQty,
			DiscountInPercent: p.DiscountInPercent,
		})
	}

	return entities, nil
}
