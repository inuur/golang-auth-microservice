package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database

func InitClient(ctx context.Context, host, port, database string) {
	mongoDBURL := fmt.Sprintf("mongodb://%s:%s/%s", host, port, database)
	Client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoDBURL))

	if err != nil {
		panic(err)
	}

	db = Client.Database("auth-service")

	mod := mongo.IndexModel{
		Keys:    bson.M{"username": 1},
		Options: options.Index().SetUnique(true),
	}

	_, err = db.Collection("users").Indexes().CreateOne(context.Background(), mod)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func GetDatabase() *mongo.Database {
	if db == nil {
		panic("Database connection is not established!")
	}
	return db
}
