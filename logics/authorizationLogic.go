package logics

import (
	"github.com/dgrijalva/jwt-go"
)

type AuthorizationLogic struct {
	authToken string
}

func (authorizeLogic *AuthorizationLogic) SetAuthToken(authToken string) {
	authorizeLogic.authToken = authToken
}

func (authorizeLogic *AuthorizationLogic) Authorize() int {
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(
		(*authorizeLogic).authToken,
		claims,
		func(token *jwt.Token) (interface{}, error) {
			return []byte("Secrate_Code"), nil
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
