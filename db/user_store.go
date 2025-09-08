package db

import (
	"context"

	"github.com/fulltimegodev/hotel-reservation/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)


const userColl = "users"

type UserStore interface {
	GetUserByID(context.Context, string) (types.User, error)
}

type MongoUserStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewMongoUserStore(client *mongo.Client) *MongoUserStore {
	return &MongoUserStore{
		client: client,
		coll:   client.Database(DBNAME).Collection(userColl),
	}
}

func (s *MongoUserStore) GetUserByID(ctx context.Context, id string) (*types.User, error) {
	var user types.User
	if err := s.coll.FindOne(ctx, bson.M{"_id": }).Decode(&user); err != nil {
		return nil, err
	}
	return &user,nil
}
