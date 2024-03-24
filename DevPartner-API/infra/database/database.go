package database

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewConnect(db_connect_url string) *pgxpool.Pool {

	pool, err := pgxpool.New(context.Background(), db_connect_url)

	if err != nil {
		log.Fatal(err)
	}

	return pool
}
