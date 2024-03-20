package database

import (
	"VUTTR-API/domain/tool"
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ToolRepositoryImp struct {
	pool *mongo.Client
}

func (db ToolRepositoryImp) Get() ([]tool.Tool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	tools := make([]tool.Tool, 0)

	rows, err := db.pool.Database("vuttr").Collection("tools").Find(ctx, bson.D{})

	if err != nil {
		return nil, err
	}

	for rows.Next(ctx) {
		var tool tool.Tool

		err = rows.Decode(&tool)

		if err != nil {
			log.Fatal(err)
		}

		tools = append(tools, tool)
	}

	return tools, nil
}

func (db ToolRepositoryImp) GetById(id string) (tool.Tool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	var toolFinded tool.Tool

	filter := bson.D{{Key: "id", Value: id}}

	err := db.pool.Database("vuttr").Collection("tools").FindOne(ctx, filter).Decode(&toolFinded)

	if err != nil {
		return tool.Tool{}, err
	}

	return toolFinded, nil
}

func (db ToolRepositoryImp) GetByTag(tag string) ([]tool.Tool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	tools := make([]tool.Tool, 0)

	filter := bson.M{"tags": bson.M{"$in": []string{tag}}}

	rows, err := db.pool.Database("vuttr").Collection("tools").Find(ctx, filter)

	if err != nil {
		return nil, err
	}

	for rows.Next(ctx) {
		var tool tool.Tool

		err = rows.Decode(&tool)

		if err != nil {
			log.Fatal(err)
		}

		tools = append(tools, tool)
	}

	return tools, nil

}

func (db ToolRepositoryImp) Update(data tool.Tool) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	filter := bson.D{{Key: "id", Value: data.Id}}

	content := bson.D{{
		Key: "$set",
		Value: bson.D{
			{Key: "title", Value: data.Title},
			{Key: "description", Value: data.Description},
			{Key: "link", Value: data.Link},
			{Key: "tags", Value: data.Tags},
		},
	}}

	_, err := db.pool.Database("vuttr").Collection("tools").UpdateByID(ctx, filter, content)

	if err != nil {
		return err
	}

	return nil
}

func (db ToolRepositoryImp) Create(data tool.Tool) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	content := bson.D{
		{Key: "id", Value: data.Id},
		{Key: "title", Value: data.Title},
		{Key: "description", Value: data.Description},
		{Key: "link", Value: data.Link},
		{Key: "tags", Value: data.Tags},
	}

	_, err := db.pool.Database("vuttr").Collection("tools").InsertOne(ctx, content)

	if err != nil {
		return err
	}

	return nil
}

func (db ToolRepositoryImp) Delete(id string) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	filter := bson.D{{Key: "id", Value: id}}

	_, err := db.pool.Database("vuttr").Collection("tools").DeleteOne(ctx, filter)

	if err != nil {
		return err
	}

	return nil
}

func NewToolRepositoryImp(pool *mongo.Client) tool.ToolRepository {
	return ToolRepositoryImp{pool: pool}
}
