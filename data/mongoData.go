package data

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoData struct {
	client   *mongo.Client
	Database *mongo.Database
}

func (mongoDt *MongoData) GetDatabaseInstance() *mongo.Database {
	return mongoDt.Database
}

func (mongoDt *MongoData) Connect() {
	var err error

	if mongoDt.client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017")); err != nil {
		panic(err)
	}

	mongoDt.Database = mongoDt.client.Database("authService")
}

func (mongoDt *MongoData) Disconnect() {
	if err := mongoDt.client.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}
