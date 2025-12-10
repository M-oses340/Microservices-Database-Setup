package store

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type MySQLStore struct {
	db *sqlx.DB
}

func NewMySQLStore(db *sqlx.DB) *MySQLStore {
	return &MySQLStore{db: db}
}

func (ms *MySQLStore) CreateProduct(ctx context.Context, p *Product) (*Product, error) {

	res, err := ms.db.NamedExecContext(ctx,
		`INSERT INTO products(name, image, category, description, rating, num_reviews, price, count_in_stock)
         VALUES (:name, :image, :category, :description, :rating, :num_reviews, :price, :count_in_stock)`, p)

	if err != nil {
		return nil, fmt.Errorf("error inserting product: %v", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("error getting last insert id: %v", err)
	}

	p.ID = int64(int(id)) // assign the generated ID into the struct

	return p, nil
}
