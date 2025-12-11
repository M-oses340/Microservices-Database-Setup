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
func (ms *MySQLStorer) CreateUser(ctx context.Context, u *User) (*User, error) {
	res, err := ms.db.NamedExecContext(ctx, "INSERT INTO users (name, email, password, is_admin) VALUES (:name, :email, :password, :is_admin)", u)
	if err != nil {
		return nil, fmt.Errorf("error inserting user: %w", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("error getting last insert ID: %w", err)
	}
	u.ID = id

	return u, nil
}

func (ms *MySQLStorer) GetUser(ctx context.Context, email string) (*User, error) {
	var u User
	err := ms.db.GetContext(ctx, &u, "SELECT * FROM users WHERE email=?", email)
	if err != nil {
		return nil, fmt.Errorf("error getting user: %w", err)
	}

	return &u, nil
}

func (ms *MySQLStorer) ListUsers(ctx context.Context) ([]*User, error) {
	var users []*User
	err := ms.db.SelectContext(ctx, &users, "SELECT * FROM users")
	if err != nil {
		return nil, fmt.Errorf("error listing users: %w", err)
	}

	return users, nil
}

func (ms *MySQLStorer) UpdateUser(ctx context.Context, u *User) (*User, error) {
	_, err := ms.db.NamedExecContext(ctx, "UPDATE users SET name=:name, email=:email, password=:password, is_admin=:is_admin, updated_at=:updated_at WHERE id=:id", u)
	if err != nil {
		return nil, fmt.Errorf("error updating user: %w", err)
	}

	return u, nil
}

func (ms *MySQLStorer) DeleteUser(ctx context.Context, id int64) error {
	_, err := ms.db.ExecContext(ctx, "DELETE FROM users WHERE id=?", id)
	if err != nil {
		return fmt.Errorf("error deleting user: %w", err)
	}

	return nil
}
