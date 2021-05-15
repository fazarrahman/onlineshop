package mysql

import (
	"context"
	"database/sql"

	"github.com/fazarrahman/onlineshop/domain/product/entity"
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
func (m *MySQL) GetAll(ctx context.Context) ([]*entity.Product, error) {
	var (
		products []*Product
		query    = `
			SELECT
				sku,
				name,
				price,
				qty
			FROM
				product
			WHERE
				status=1 ;
		`
	)

	err := m.db.SelectContext(ctx, &products, query)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	var entities []*entity.Product
	for _, p := range products {
		entities = append(entities, &entity.Product{
			SKU:   p.SKU,
			Name:  p.Name,
			Price: p.Price,
			Qty:   p.Qty,
		})
	}

	return entities, nil
}
