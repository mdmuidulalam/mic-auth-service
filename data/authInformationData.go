package data

import (
	"context"
	"time"

	datainterface "auth-service/data/interfaces"
)

type AuthInformationData struct {
	username     string
	passwordHash []byte
	createOn     time.Time
	MongoData    datainterface.IMongoData
}

func (authData *AuthInformationData) SetUserName(username string) {
	authData.username = username
}

func (authData *AuthInformationData) GetUserName() string {
	return authData.username
}

func (authData *AuthInformationData) SetPasswordHash(passwordHash []byte) {
	authData.passwordHash = passwordHash
}

func (authData *AuthInformationData) GetPasswordHash() []byte {
	return authData.passwordHash
}

func (authData *AuthInformationData) SetCreateOn(createOn time.Time) {
	authData.createOn = createOn
}

func (authData *AuthInformationData) GetCreateOn() time.Time {
	return authData.createOn
}

func (authData *AuthInformationData) InsertAuthInformation() {
	authData.MongoData.Connect()
	collection := authData.MongoData.GetDatabaseInstance().Collection("authInformation")

	if _, err := collection.InsertOne(context.TODO(), authData); err != nil {
		//TODO logging need to be done
		panic(err)
	}

	authData.MongoData.Disconnect()
}
