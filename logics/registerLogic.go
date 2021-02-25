package logics

import (
	logicinterface "auth-service/logics/interfaces"

	"golang.org/x/crypto/bcrypt"
)

type RegisterLogic struct {
	username     string
	password     string
	RegisterData logicinterface.IRegisterData
}

func (regLogic *RegisterLogic) SetUserName(username string) {
	regLogic.username = username
}

func (regLogic *RegisterLogic) SetPassword(password string) {
	regLogic.password = password
}

func (regLogic *RegisterLogic) Register() int {
	regLogic.RegisterData.SetUserName(regLogic.username)

	authInfo := regLogic.RegisterData.FindOneAuthInformation()
	if authInfo == nil {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(regLogic.password), bcrypt.DefaultCost)
		if err != nil {
			//TODO logging need to be done
			panic(err)
		}

		regLogic.RegisterData.SetPasswordHash(hashedPassword)
		regLogic.RegisterData.InsertAuthInformation()

		return 1
	} else {
		return 2
	}
}
