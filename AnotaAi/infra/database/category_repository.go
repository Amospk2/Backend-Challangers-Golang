package database

import (
	"api/domain/category"
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

type CategoryRepository struct {
	pool *pgxpool.Pool
}

func (db *CategoryRepository) Get() ([]category.Category, error) {

	categories := make([]category.Category, 0)

	rows, err := db.pool.Query(context.Background(),
		`select id, title, description, ownerID 
		from public.categories`,
	)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var category category.Category

		err = rows.Scan(
			&category.Id,
			&category.Title,
			&category.Description,
			&category.OwnerID,
		)

		if err != nil {
			log.Fatal(err)
		}

		categories = append(categories, category)
	}

	return categories, nil
}

func (db *CategoryRepository) GetById(id string) (category.Category, error) {

	var categoryFinded category.Category

	err := db.pool.QueryRow(
		context.Background(),
		`select 
			id, title, description, category
		from 
			public.categories 
		where
			id=$1`,
		id,
	).Scan(
		&categoryFinded.Id,
		&categoryFinded.Title,
		&categoryFinded.Description,
		&categoryFinded.OwnerID,
	)

	if err != nil {
		return category.Category{}, err
	}

	return categoryFinded, nil
}

func (db *CategoryRepository) Update(data category.Category) error {
	_, err := db.pool.Exec(
		context.Background(),
		`UPDATE categories 
		SET title = $1, description = $2
		WHERE id = $3`,
		data.Title, data.Description, data.Id,
	)

	if err != nil {
		return err
	}

	return nil
}

func (db *CategoryRepository) Create(data category.Category) error {
	_, err := db.pool.Exec(
		context.Background(), "INSERT INTO categories VALUES($1,$2,$3)",
		data.Id, data.Title, data.Description,
	)

	if err != nil {
		return err
	}

	return nil
}

func (db *CategoryRepository) Delete(id string) error {

	_, err := db.pool.Exec(context.Background(), "DELETE FROM categories WHERE id=$1", id)

	if err != nil {
		return err
	}

	return nil
}

func NewCategoryRepository(pool *pgxpool.Pool) *CategoryRepository {
	return &CategoryRepository{pool: pool}
}
