package datainterface

import "go.mongodb.org/mongo-driver/mongo"

type IMongoData interface {
	GetDatabaseInstance() *mongo.Database
	Connect()
	Disconnect()
}
