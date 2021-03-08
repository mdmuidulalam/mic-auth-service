package logics

import (
	"auth-service/config"
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
	Config             *config.Config
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
	configInstance := authLogic.Config.GetConfig()
	claims := jwt.MapClaims{}
	claims["username"] = authLogic.username
	claims["siteGroup"] = authLogic.siteGroup
	claims["exp"] = time.Now().Add(time.Minute * 60 * 24 * 30).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(configInstance.AuthTokenSecreteCode))
	if err != nil {
		panic(err)
	}
	return tokenString
}
