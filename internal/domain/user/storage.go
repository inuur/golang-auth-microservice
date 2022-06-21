package user

import (
	"authService/pkg/client/mongodb"
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Storage interface {
	Create(ctx context.Context, user User) (*User, error)
	FindOne(ctx context.Context, id string) (*User, error)
}

type db struct {
	collection *mongo.Collection
}

func (d db) Create(ctx context.Context, user User) (*User, error) {
	result, err := d.collection.InsertOne(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("failed to create user due to errors: %v", err)
	}

	oid, ok := result.InsertedID.(primitive.ObjectID)

	if !ok {
		return nil, fmt.Errorf("failed to convert object id to HEX. oid: %s", oid)
	}

	return d.FindOne(ctx, oid.Hex())
}

func (d db) FindOne(ctx context.Context, id string) (u *User, err error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return u, fmt.Errorf("failed to convert hex to objectID. hex: %s", id)
	}

	filter := bson.M{"_id": oid}

	result := d.collection.FindOne(ctx, filter)
	if result.Err() != nil {
		if errors.Is(result.Err(), mongo.ErrNoDocuments) {
			return u, fmt.Errorf("ErrEntityNotFound")
		}
		return u, fmt.Errorf("failed to find one user by id: %s", id)
	}
	if err := result.Decode(&u); err != nil {
		return u, fmt.Errorf("failed to decode user(id:%s) from database: %s", id, err)
	}
	return u, nil
}

func NewStorage() Storage {
	return &db{
		collection: mongodb.GetDatabase().Collection("users"),
	}
}
