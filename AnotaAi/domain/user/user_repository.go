package user

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository interface {
	Get() ([]Users, error)
	GetById(id string) (Users, error)
	GetByEmail(email string) (Users, error)
	Update(data Users) error
	Create(data Users) error
	Delete(id string) error
}

type UserRepositoryImp struct {
	pool *pgxpool.Pool
}

func (db UserRepositoryImp) Get() ([]Users, error) {

	users := make([]Users, 0)

	rows, err := db.pool.Query(context.Background(),
		"select id, name, email, age, password from public.users",
	)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var user Users

		err = rows.Scan(
			&user.Id,
			&user.Name,
			&user.Email,
			&user.Age,
			&user.Password,
		)

		if err != nil {
			log.Fatal(err)
		}

		users = append(users, user)
	}

	return users, nil
}

func (db UserRepositoryImp) GetById(id string) (Users, error) {

	var userFinded Users

	err := db.pool.QueryRow(
		context.Background(),
		"select id, name, email, age, password from public.users where id=$1", id,
	).Scan(
		&userFinded.Id,
		&userFinded.Name,
		&userFinded.Email,
		&userFinded.Age,
		&userFinded.Password,
	)

	if err != nil {
		return Users{}, err
	}

	return userFinded, nil
}

func (db UserRepositoryImp) GetByEmail(email string) (Users, error) {

	var userFinded Users

	err := db.pool.QueryRow(
		context.Background(),
		"select id, name, email, age, password from public.users where email=$1", email,
	).Scan(
		&userFinded.Id,
		&userFinded.Name,
		&userFinded.Email,
		&userFinded.Age,
		&userFinded.Password,
	)

	if err != nil {
		return Users{}, err
	}

	return userFinded, nil
}

func (db UserRepositoryImp) Update(data Users) error {
	_, err := db.pool.Exec(
		context.Background(),
		"UPDATE USERS SET name = $1, email = $2, age = $3, password = $4 WHERE id = $5",
		data.Name, data.Email, data.Age, data.Password, data.Id,
	)

	if err != nil {
		return err
	}

	return nil
}

func (db UserRepositoryImp) Create(data Users) error {
	_, err := db.pool.Exec(
		context.Background(), "INSERT INTO USERS VALUES($1,$2,$3,$4, $5)",
		data.Id, data.Name, data.Email, data.Age, data.Password,
	)

	if err != nil {
		return err
	}

	return nil
}

func (db UserRepositoryImp) Delete(id string) error {

	_, err := db.pool.Exec(context.Background(), "DELETE FROM USERS WHERE id=$1", id)

	if err != nil {
		return err
	}

	return nil
}

func NewUserRepositoryImp(pool *pgxpool.Pool) UserRepository {
	return UserRepositoryImp{pool: pool}
}
