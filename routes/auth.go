package routes

import (
	"auth-service/config"
	routesinterface "auth-service/routes/interfaces"

	"github.com/gin-gonic/gin"
)

type Auth struct {
	R                   *gin.Engine
	AuthenticationLogic routesinterface.IAuthenticationLogic
	RegisterLogic       routesinterface.IRegisterLogic
	AuthorizationLogic  routesinterface.IAuthorizationLogic
	SiteGroupsConfig    *config.SiteGroupsConfig
}

func (auth Auth) New() {
	auth.R.POST("authenticate", auth.Authenticate)
	auth.R.POST("register", auth.Register)
	auth.R.GET("authorize", auth.Authorize)
}

// @Summary Authenticate an user
// @Description It will authenticate an user from provided information
// @ID authenticate-user
// @Router /authenticate [post]
// @Accept json
// @Param username body string true "Username of the user"
// @Param password body string true "Password of the user"
// @Success 200 {boolean} boolean "The authentication is complete and access token provided in http only cookies with key 'auth-token'"
// @Success 210 {boolean} boolean "Wrong authentication information provided"
// @Success 211 {boolean} boolean "The siteGroup doesn't exists"
func (auth *Auth) Authenticate(c *gin.Context) {
	var authInfo authenticationInformation
	if err := c.ShouldBind(&authInfo); err != nil {
		panic(err)
	}
	if _, exist := (*auth.SiteGroupsConfig.GetSiteGroupConfig())[authInfo.SiteGroup]; !exist {
		c.Writer.WriteHeader(211)
		return
	}

	auth.AuthenticationLogic.SetUserName(authInfo.Username)
	auth.AuthenticationLogic.SetPassword(authInfo.Password)
	auth.AuthenticationLogic.SetSiteGroup(authInfo.SiteGroup)

	authStatus, token := auth.AuthenticationLogic.Authenticate()

	statusCode := 0
	if authStatus == 1 {
		statusCode = 200
	} else if authStatus == 2 {
		statusCode = 210
	}

	c.SetCookie("auth-token", token, 60*60*24*30, "", "", false, true)
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
// @Success 211 {boolean} boolean "The siteGroup doesn't exists"
func (auth *Auth) Register(c *gin.Context) {
	var authInfo authenticationInformation
	if err := c.ShouldBind(&authInfo); err != nil {
		panic(err)
	}
	if _, exist := (*auth.SiteGroupsConfig.GetSiteGroupConfig())[authInfo.SiteGroup]; !exist {
		c.Writer.WriteHeader(211)
		return
	}

	auth.RegisterLogic.SetUserName(authInfo.Username)
	auth.RegisterLogic.SetPassword(authInfo.Password)
	auth.RegisterLogic.SetSiteGroup(authInfo.SiteGroup)

	registerStatus := auth.RegisterLogic.Register()

	statusCode := 0
	if registerStatus == 1 {
		statusCode = 200
	} else if registerStatus == 2 {
		statusCode = 210
	}

	c.Writer.WriteHeader(statusCode)
}

// @Summary Authorize an user
// @Description It will authorize an user with jwt token from http only cookies with key name 'auth-token'
// @ID authorize-user
// @Router /authorize [get]
// @Accept json
// @Success 200 {boolean} boolean "The user is authorized"
// @Success 210 {boolean} boolean "Wrong jwt token"
func (auth *Auth) Authorize(c *gin.Context) {
	authToken, err := c.Cookie("auth-token")
	if err != nil {
		panic(err)
	}

	auth.AuthorizationLogic.SetAuthToken(authToken)

	authorizationStatus := auth.AuthorizationLogic.Authorize()
	if authorizationStatus == 1 {
		c.Writer.WriteHeader(200)
	} else {
		c.Writer.WriteHeader(210)
	}
}

//* Data Classes

//* Dealing authentication information
type authenticationInformation struct {
	Username  string `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required"`
	SiteGroup string `json:"siteGroup" binding:"required"`
}
