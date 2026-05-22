package repository

import (
	"context"
	"database/sql"
	"time"
)

type ProductsRepository interface {
	GetProducts(context context.Context) ([]any, error)
}

type productsRepository struct {
	db *sql.DB
}

func NewProductsRepository(db *sql.DB) *productsRepository {
	return &productsRepository{db: db}
}

func (r *productsRepository) GetProducts(ctx context.Context) ([]any, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT * FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []any
	for rows.Next() {
		var product productRow
		if err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.PriceInCenters,
			&product.Quantity,
			&product.CreatedAt,
			&product.UpdatedAt,
		); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return products, nil
}

type productRow struct {
	ID             int64     `json:"id"`
	Name           string    `json:"name"`
	PriceInCenters int32     `json:"price_in_centers"`
	Quantity       int32     `json:"quantity"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
