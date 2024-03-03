package database

import (
	"api/domain/product"
	"context"
	"database/sql"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ProductRepository struct {
	pool *pgxpool.Pool
}

func (db *ProductRepository) Get() ([]product.Product, error) {

	products := make([]product.Product, 0)

	rows, err := db.pool.Query(context.Background(),
		`select id, title, description, price, category, ownerid
		from public.products`,
	)

	if err != nil {

		return nil, err
	}

	for rows.Next() {
		var (
			product  product.Product
			category sql.NullString
		)

		err = rows.Scan(
			&product.Id,
			&product.Title,
			&product.Description,
			&product.Price,
			&category,
			&product.OwnerID,
		)

		if err != nil {
			log.Fatal(err)
		}

		if category.Valid {
			product.Category = category.String
		}

		products = append(products, product)
	}

	return products, nil
}

func (db *ProductRepository) GetById(id string) (product.Product, error) {

	var productFinded product.Product

	err := db.pool.QueryRow(
		context.Background(),
		`select 
			id, title, description, price, category, ownerID 
		from 
			public.products 
		where
			id=$1`,
		id,
	).Scan(
		&productFinded.Id,
		&productFinded.Title,
		&productFinded.Description,
		&productFinded.Price,
		&productFinded.Category,
		&productFinded.OwnerID,
	)

	if err != nil {
		return product.Product{}, err
	}

	return productFinded, nil
}

func (db *ProductRepository) Update(data product.Product) error {
	_, err := db.pool.Exec(
		context.Background(),
		`UPDATE products 
		SET title = $1, description = $2, price = $3, ownerID = $4 
		WHERE id = $5`,
		data.Title, data.Description, data.Price, data.OwnerID, data.Id,
	)

	if err != nil {
		return err
	}

	return nil
}

func (db *ProductRepository) Create(data product.Product) error {
	_, err := db.pool.Exec(
		context.Background(), "INSERT INTO products VALUES($1,$2,$3,$4,$5,$6)",
		data.Id, data.Title, data.Description, data.Price, data.Category, data.OwnerID,
	)

	if err != nil {
		return err
	}

	return nil
}

func (db *ProductRepository) Delete(id string) error {

	_, err := db.pool.Exec(context.Background(), "DELETE FROM products WHERE id=$1", id)

	if err != nil {
		return err
	}

	return nil
}

func NewProductRepository(pool *pgxpool.Pool) *ProductRepository {
	return &ProductRepository{pool: pool}
}
