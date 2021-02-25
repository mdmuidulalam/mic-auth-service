package routes

import (
	routesinterface "auth-service/routes/interfaces"

	"github.com/gin-gonic/gin"
)

type Auth struct {
	R                   *gin.Engine
	AuthenticationLogic routesinterface.IAuthenticationLogic
	RegisterLogic       routesinterface.IRegisterLogic
}

func (auth Auth) New() {
	auth.R.POST("authenticate", auth.Authenticate)
	auth.R.POST("register", auth.Register)
}

// @Summary Authenticate an user
// @Description It will authenticate an user from provided information
// @ID authenticate-user
// @Router /authenticate [post]
// @Accept json
// @Param username body string true "Username of the user"
// @Param password body string true "Password of the user"
// @Success 200 {boolean} boolean "The authentication is complete and access token provided in http only cookies"
// @Success 210 {boolean} boolean "Wrong authentication information provided"
func (auth Auth) Authenticate(c *gin.Context) {
	var authInfo authenticationInformation

	if err := c.ShouldBind(&authInfo); err != nil {
		//TODO logging need to be done
		panic(err)
	}

	auth.AuthenticationLogic.SetUserName(authInfo.Username)
	auth.AuthenticationLogic.SetPassword(authInfo.Password)

	authStatus, token := auth.AuthenticationLogic.Authenticate()

	statusCode := 0
	if authStatus == 1 {
		statusCode = 200
	} else if authStatus == 2 {
		statusCode = 210
	}

	c.SetCookie("auth-token", token, 60*60*24, "", "", true, true)
	c.Writer.WriteHeader(statusCode)
}

// @Summary Register an user
// @Description It will register an user with provided information
// @ID register-user
// @Router /register [post]
// @Accept json
// @Param username body string true "Username of the user"
// @Param password body string true "Password of the user"
// @Success 200 {boolean} boolean "The registration is completed"
// @Success 210 {boolean} boolean "The user is already registered"
func (auth Auth) Register(c *gin.Context) {
	var authInfo authenticationInformation

	if err := c.ShouldBind(&authInfo); err != nil {
		//TODO logging need to be done
		panic(err)
	}

	auth.RegisterLogic.SetUserName(authInfo.Username)
	auth.RegisterLogic.SetPassword(authInfo.Password)

	registerStatus := auth.RegisterLogic.Register()

	statusCode := 0
	if registerStatus == 1 {
		statusCode = 200
	} else if registerStatus == 2 {
		statusCode = 210
	}

	c.Writer.WriteHeader(statusCode)
}

//* Data Classes

//* Dealing authentication information
type authenticationInformation struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
