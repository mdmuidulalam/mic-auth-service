package data

import (
	"context"
	"fmt"
	"time"

	datainterface "auth-service/data/interfaces"
	logicinterface "auth-service/logics/interfaces"

	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type AuthInformationData struct {
	Username     string
	PasswordHash []byte
	CreateOn     time.Time
	siteGroup    string                   `bson:"-"`
	MongoData    datainterface.IMongoData `bson:"-"`
}

func (authData *AuthInformationData) SetSiteGroup(siteGroup string) {
	authData.siteGroup = siteGroup
}

func (authData *AuthInformationData) SetUserName(username string) {
	authData.Username = username
}

func (authData *AuthInformationData) GetUserName() string {
	return authData.Username
}

func (authData *AuthInformationData) SetPasswordHash(passwordHash []byte) {
	authData.PasswordHash = passwordHash
}

func (authData *AuthInformationData) GetPasswordHash() []byte {
	return authData.PasswordHash
}

func (authData *AuthInformationData) SetCreateOn(createOn time.Time) {
	authData.CreateOn = createOn
}

func (authData *AuthInformationData) GetCreateOn() time.Time {
	return authData.CreateOn
}

func (authData *AuthInformationData) InsertAuthInformation() {
	authData.MongoData.SetSiteGroup(authData.siteGroup)
	authData.MongoData.Connect()
	defer authData.MongoData.Disconnect()

	collection := authData.MongoData.GetDatabaseInstance().Collection("authInformation")

	fmt.Println(authData)
	if _, err := collection.InsertOne(context.TODO(), authData); err != nil {
		panic(err)
	}
}

func (authData *AuthInformationData) FindOneAuthInformation() *logicinterface.IAuthInformation {
	authData.MongoData.SetSiteGroup(authData.siteGroup)
	authData.MongoData.Connect()
	defer authData.MongoData.Disconnect()

	collection := authData.MongoData.GetDatabaseInstance().Collection("authInformation")

	if err := collection.FindOne(context.TODO(), bson.M{"username": authData.Username}).Decode(authData); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil
		} else {
			panic(err)
		}
	}

	var authInformation logicinterface.IAuthInformation = authData

	return &authInformation
}
