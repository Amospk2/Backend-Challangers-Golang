package repository

import (
	"context"
	"devpartner-api/domain/nota"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type NotaRepository struct {
	pool *pgxpool.Pool
}

func (db *NotaRepository) Get() ([]nota.Nota, error) {

	notas := make([]nota.Nota, 0)

	rows, err := db.pool.Query(context.Background(),
		`select notaFiscalId, numeroNf, valorTotal, dataNf, cnpjEmissorNf, cnpjDestinatarioNf
		from public.notas`,
	)

	if err != nil {

		return nil, err
	}

	for rows.Next() {
		var nota nota.Nota
		err = rows.Scan(
			&nota.Id,
			&nota.NumeroNf,
			&nota.ValorTotal,
			&nota.DataNf,
			&nota.CnpjEmissorNf,
			&nota.CnpjDestinatarioNf,
		)

		if err != nil {
			fmt.Println(err)
		}

		notas = append(notas, nota)
	}

	return notas, nil
}

func (db *NotaRepository) GetById(id string) (nota.Nota, error) {

	var notaFinded nota.Nota

	err := db.pool.QueryRow(
		context.Background(),
		`select 
			notaFiscalId, numeroNf, valorTotal, dataNf, cnpjEmissorNf, cnpjDestinatarioNf
		from 
			notas 
		where
			id=$1`,
		id,
	).Scan(
		&notaFinded.Id,
		&notaFinded.NumeroNf,
		&notaFinded.ValorTotal,
		&notaFinded.DataNf,
		&notaFinded.CnpjEmissorNf,
		&notaFinded.CnpjDestinatarioNf,
	)

	if err != nil {
		return nota.Nota{}, err
	}

	return notaFinded, nil
}

func (db *NotaRepository) Update(data nota.Nota) error {
	_, err := db.pool.Exec(
		context.Background(),
		`UPDATE notas 
		SET numeroNf = $1, valorTotal = $2, 
			dataNf = $3, cnpjEmissorNf = $4, cnpjDestinatarioNf = $5
		WHERE id = $6`,
		&data.NumeroNf,
		&data.ValorTotal,
		&data.DataNf,
		&data.CnpjEmissorNf,
		&data.CnpjDestinatarioNf,
		&data.Id,
	)

	if err != nil {
		return err
	}

	return nil
}

func (db *NotaRepository) Create(data nota.Nota) error {
	_, err := db.pool.Exec(
		context.Background(), "INSERT INTO notas VALUES($1,$2,$3,$4,$5,$6)",
		&data.Id,
		&data.NumeroNf,
		&data.ValorTotal,
		&data.DataNf,
		&data.CnpjEmissorNf,
		&data.CnpjDestinatarioNf,
	)

	if err != nil {
		return err
	}

	return nil
}

func (db *NotaRepository) Delete(id string) error {

	_, err := db.pool.Exec(context.Background(), "DELETE FROM notas WHERE id=$1", id)

	if err != nil {
		return err
	}

	return nil
}

func NewNotaRepository(pool *pgxpool.Pool) *NotaRepository {
	return &NotaRepository{pool: pool}
}
