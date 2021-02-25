package logics

import (
	logicinterface "auth-service/logics/interfaces"

	"golang.org/x/crypto/bcrypt"
)

type AuthenticationLogic struct {
	username           string
	password           string
	AuthenticationData logicinterface.IAuthenticationData
}

func (authLogic AuthenticationLogic) SetUserName(username string) {
	authLogic.username = username
}

func (authLogic AuthenticationLogic) SetPassword(password string) {
	authLogic.password = password
}

func (authLogic AuthenticationLogic) Authenticate() (int, string) {
	authLogic.AuthenticationData.SetUserName(authLogic.username)

	hashedPassword, active := authLogic.AuthenticationData.FetchUserHashPasswordAndActive()

	// hashedPassword, err := bcrypt.GenerateFromPassword([]byte(authLogic.password), bcrypt.DefaultCost)
	// if err != nil {
	// 	panic(err)
	// }

	if bcrypt.CompareHashAndPassword(hashedPassword, []byte(authLogic.password)) == nil {
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
