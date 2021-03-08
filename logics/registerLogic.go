package logics

import (
	"auth-service/config"
	logicinterface "auth-service/logics/interfaces"

	"golang.org/x/crypto/bcrypt"
)

type RegisterLogic struct {
	username     string
	password     string
	siteGroup    string
	RegisterData logicinterface.IRegisterData
	Config       *config.Config
}

func (regLogic *RegisterLogic) SetUserName(username string) {
	regLogic.username = username
}

func (regLogic *RegisterLogic) SetPassword(password string) {
	regLogic.password = password
}

func (regLogic *RegisterLogic) SetSiteGroup(siteGroup string) {
	regLogic.siteGroup = siteGroup
}

func (regLogic *RegisterLogic) Register() int {
	configInstance := regLogic.Config.GetConfig()
	regLogic.RegisterData.SetUserName(regLogic.username)
	regLogic.RegisterData.SetSiteGroup(regLogic.siteGroup)

	authInfo := regLogic.RegisterData.FindOneAuthInformation()
	if authInfo == nil {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(regLogic.password), configInstance.HashPasswordCost)
		if err != nil {
			panic(err)
		}

		regLogic.RegisterData.SetPasswordHash(hashedPassword)
		regLogic.RegisterData.InsertAuthInformation()

		return 1
	} else {
		return 2
	}
}
