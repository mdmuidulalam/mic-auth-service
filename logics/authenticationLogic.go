package logics

import (
	logicinterface "auth-service/logics/interfaces"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type AuthenticationLogic struct {
	username           string
	password           string
	siteGroup          string
	AuthenticationData logicinterface.IAuthenticationData
}

func (authLogic *AuthenticationLogic) SetUserName(username string) {
	authLogic.username = username
}

func (authLogic *AuthenticationLogic) SetPassword(password string) {
	authLogic.password = password
}

func (regLogic *AuthenticationLogic) SetSiteGroup(siteGroup string) {
	regLogic.siteGroup = siteGroup
}

func (authLogic *AuthenticationLogic) Authenticate() (int, string) {
	authLogic.AuthenticationData.SetUserName(authLogic.username)
	authLogic.AuthenticationData.SetSiteGroup(authLogic.siteGroup)

	authInformation := authLogic.AuthenticationData.FindOneAuthInformation()

	if authInformation != nil && bcrypt.CompareHashAndPassword((*authInformation).GetPasswordHash(), []byte(authLogic.password)) == nil {
		token := authLogic.createToken()
		return 1, token
	} else {
		return 2, ""
	}
}

func (authLogic *AuthenticationLogic) createToken() string {
	claims := jwt.MapClaims{}
	claims["username"] = authLogic.username
	claims["siteGroup"] = authLogic.siteGroup
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("Secrate_Code"))
	if err != nil {
		panic(err)
	}
	return tokenString
}
