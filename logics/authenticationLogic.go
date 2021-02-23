package logics

import (
	logicinterface "auth-service/logic/interfaces"
)

type AuthenticationLogic struct {
	username           string
	password           string
	authenticationData logicinterface.IAuthenticationData
}

func (authLogic AuthenticationLogic) SetUserName(username string) {
	authLogic.username = username
}

func (authLogic AuthenticationLogic) SetPassword(password string) {
	authLogic.password = password
}

func (authLogic AuthenticationLogic) Authenticate() (int, string) {
	authLogic.authenticationData.SetUserName(authLogic.username)

	passwordHash, active := authLogic.authenticationData.FetchUserHashPasswordAndActive()

	if passwordHash == authLogic.password {
		if active {
			token := "d6sd4f5d456asd1f5a1f.f5asd16a1d.4da6sd4"
			return 1, token
		} else {
			return 3, ""
		}
	} else {
		return 2, ""
	}
}
