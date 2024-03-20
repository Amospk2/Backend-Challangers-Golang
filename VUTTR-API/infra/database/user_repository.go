package database

import (
	"VUTTR-API/domain/user"
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepositoryImp struct {
	pool *mongo.Client
}

func (db UserRepositoryImp) Get() ([]user.Users, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	users := make([]user.Users, 0)

	rows, err := db.pool.Database("vuttr").Collection("users").Find(ctx, bson.D{})

	if err != nil {
		return nil, err
	}

	for rows.Next(ctx) {
		var user user.Users

		err = rows.Decode(&user)

		if err != nil {
			log.Fatal(err)
		}

		users = append(users, user)
	}

	return users, nil
}

func (db UserRepositoryImp) GetById(id string) (user.Users, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	var userFinded user.Users

	filter := bson.D{{Key: "id", Value: id}}

	err := db.pool.Database("vuttr").Collection("users").FindOne(ctx, filter).Decode(&userFinded)

	if err != nil {
		return user.Users{}, err
	}

	return userFinded, nil
}

func (db UserRepositoryImp) GetByEmail(email string) (user.Users, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	var userFinded user.Users

	filter := bson.D{{Key: "email", Value: email}}

	err := db.pool.Database("vuttr").Collection("users").FindOne(ctx, filter).Decode(&userFinded)

	if err != nil {
		return user.Users{}, err
	}

	return userFinded, nil
}

func (db UserRepositoryImp) Update(data user.Users) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	filter := bson.D{{Key: "id", Value: data.Id}}

	content := bson.D{{
		Key: "$set",
		Value: bson.D{
			{Key: "email", Value: data.Email},
			{Key: "name", Value: data.Name},
			{Key: "password", Value: data.Password},
		},
	}}

	_, err := db.pool.Database("vuttr").Collection("users").UpdateByID(ctx, filter, content)

	if err != nil {
		return err
	}

	return nil
}

func (db UserRepositoryImp) Create(data user.Users) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	content := bson.D{
		{Key: "id", Value: data.Id},
		{Key: "email", Value: data.Email},
		{Key: "name", Value: data.Name},
		{Key: "password", Value: data.Password},
	}

	_, err := db.pool.Database("vuttr").Collection("users").InsertOne(ctx, content)

	if err != nil {
		return err
	}

	return nil
}

func (db UserRepositoryImp) Delete(id string) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	filter := bson.D{{Key: "id", Value: id}}

	_, err := db.pool.Database("vuttr").Collection("users").DeleteOne(ctx, filter)

	if err != nil {
		return err
	}

	return nil
}

func NewUserRepositoryImp(pool *mongo.Client) user.UserRepository {
	return UserRepositoryImp{pool: pool}
}
