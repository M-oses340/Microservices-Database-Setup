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
func (ms *MySQLStore) GetProduct(ctx context.Context, id int64) (*Product, error) {
	var p Product

	err := ms.db.GetContext(ctx, &p, `SELECT * FROM products WHERE id = ?`, id)
	if err != nil {
		return nil, fmt.Errorf("error getting product: %v", err)
	}

	return &p, nil
}
func (ms *MySQLStore) UpdateProduct(ctx context.Context, p *Product) (*Product, error) {
	query := `
		UPDATE products
		SET name = :name,
		    image = :image,
		    category = :category,
		    description = :description,
		    rating = :rating,
		    num_reviews = :num_reviews,
		    price = :price,
		    count_in_stock = :count_in_stock
		WHERE id = :id;
	`

	// using NamedExec because we are passing a struct with tags
	_, err := ms.db.NamedExecContext(ctx, query, p)
	if err != nil {
		return nil, fmt.Errorf("error updating product: %v", err)
	}

	return p, nil
}
func (ms *MySQLStore) DeleteProduct(ctx context.Context, id int64) error {
	query := `DELETE FROM products WHERE id = ?`

	res, err := ms.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("error deleting product: %v", err)
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("error checking delete result: %v", err)
	}

	if affected == 0 {
		return fmt.Errorf("no product found with id %d", id)
	}

	return nil
}
