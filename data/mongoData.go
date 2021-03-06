package data

import (
	"auth-service/config"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoData struct {
	client           *mongo.Client
	Database         *mongo.Database
	siteGroup        string
	SiteGroupsConfig *config.SiteGroupsConfig
}

func (mongoDt *MongoData) GetDatabaseInstance() *mongo.Database {
	return mongoDt.Database
}

func (mongoDt *MongoData) SetSiteGroup(siteGroup string) {
	mongoDt.siteGroup = siteGroup
}

func (mongoDt *MongoData) Connect() {
	var err error
	siteGroupConfig := (*mongoDt.SiteGroupsConfig.GetSiteGroupConfig())[mongoDt.siteGroup]

	if mongoDt.client, err = mongo.Connect(
		context.TODO(), options.Client().ApplyURI("mongodb://"+siteGroupConfig.Database.Host+":"+siteGroupConfig.Database.Port)); err != nil {
		panic(err)
	}

	mongoDt.Database = mongoDt.client.Database(siteGroupConfig.Database.Name)
}

func (mongoDt *MongoData) Disconnect() {
	if err := mongoDt.client.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}
