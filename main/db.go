package main

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

type AppData struct {
	DB *mongo.Client
}

type TestInsert struct {
	Content string    `json:"Content"`
	Created time.Time `json:"Created"`
}

func (d *AppData) Get() ([]*TestInsert, error) {
	var result []*TestInsert
	collection := d.DB.Database("test").Collection("examples")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	cur, err := collection.Find(ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}
	for cur.Next(context.TODO()) {
		var elem TestInsert
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, &elem)
	}
	cur.Close(context.TODO())
	return result, nil
}

func (d *AppData) Insert(content string) (*mongo.InsertOneResult, error) {
	collection := d.DB.Database("test").Collection("examples")
	insert := TestInsert{
		Content: content,
		Created: time.Now(),
	}
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	insertResult, err := collection.InsertOne(ctx, insert)
	if err != nil {
		return &mongo.InsertOneResult{}, err
	}
	return insertResult, nil
}
