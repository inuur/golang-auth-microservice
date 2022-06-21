package token

import (
	"authService/pkg/client/mongodb"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type db struct {
	collection *mongo.Collection
}

type Storage interface {
	SaveRefreshToken(ctx context.Context, refreshToken string, atUUID string) error
	FindRefreshToken(ctx context.Context, atUUID string) (*RefreshToken, error)
	DeleteRefreshToken(ctx context.Context, atUUID string) error
}

func (d db) SaveRefreshToken(ctx context.Context, refreshToken string, atUUID string) error {
	_, err := d.collection.InsertOne(ctx, &RefreshToken{
		AccessTokenID: atUUID,
		Token:         refreshToken,
	})
	if err != nil {
		return err
	}
	return nil
}

func (d db) FindRefreshToken(ctx context.Context, atUUID string) (*RefreshToken, error) {
	rt := &RefreshToken{}
	filter := bson.M{"access_token_id": atUUID}

	result := d.collection.FindOne(ctx, filter)
	err := result.Decode(rt)
	if err != nil {
		return nil, err
	}

	return rt, nil
}

func (d db) DeleteRefreshToken(ctx context.Context, atUUID string) error {
	filter := bson.M{"access_token_id": atUUID}
	_, err := d.collection.DeleteOne(ctx, filter)
	return err
}

func NewStorage() Storage {
	return &db{
		collection: mongodb.GetDatabase().Collection("tokens"),
	}
}
