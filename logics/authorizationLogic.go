package logics

import (
	"auth-service/config"

	"github.com/dgrijalva/jwt-go"
)

type AuthorizationLogic struct {
	authToken string
	Config    *config.Config
}

func (authorizeLogic *AuthorizationLogic) SetAuthToken(authToken string) {
	authorizeLogic.authToken = authToken
}

func (authorizeLogic *AuthorizationLogic) Authorize() int {
	configInstance := authorizeLogic.Config.GetConfig()
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(
		(*authorizeLogic).authToken,
		claims,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(configInstance.AuthTokenSecreteCode), nil
		},
	)
	if err != nil {
		return 2
	}
	if !token.Valid {
		return 2
	}
	return 1
}
