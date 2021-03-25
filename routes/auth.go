package routes

import (
	"auth-service/config"
	"auth-service/data"
	"auth-service/logics"

	routesinterface "auth-service/routes/interfaces"

	"github.com/gin-gonic/gin"
)

var singletonSiteGroupsConfig *config.SiteGroupsConfig
var singletonConfig *config.Config

func AuthRoutes(router *gin.Engine, siteGroupsConfig *config.SiteGroupsConfig, config *config.Config) {
	singletonSiteGroupsConfig = siteGroupsConfig
	singletonConfig = config
	router.POST("/authenticate", authenticate)
	router.POST("/register", register)
	router.GET("/authorize", authorize)
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
func authenticate(c *gin.Context) {
	var authInfo authenticationInformation
	if err := c.ShouldBind(&authInfo); err != nil {
		panic(err)
	}
	if _, exist := (*singletonSiteGroupsConfig.GetSiteGroupConfig())[authInfo.SiteGroup]; !exist {
		c.Writer.WriteHeader(211)
		return
	}

	var authenticationLogic routesinterface.IAuthenticationLogic
	authenticationLogic = &logics.AuthenticationLogic{
		AuthenticationData: &data.AuthInformationData{
			MongoData: &data.MongoData{
				SiteGroupsConfig: singletonSiteGroupsConfig,
			},
		},
		Config: singletonConfig,
	}

	authenticationLogic.SetUserName(authInfo.Username)
	authenticationLogic.SetPassword(authInfo.Password)
	authenticationLogic.SetSiteGroup(authInfo.SiteGroup)

	authStatus, token := authenticationLogic.Authenticate()

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
func register(c *gin.Context) {
	var authInfo authenticationInformation
	if err := c.ShouldBind(&authInfo); err != nil {
		panic(err)
	}
	if _, exist := (*singletonSiteGroupsConfig.GetSiteGroupConfig())[authInfo.SiteGroup]; !exist {
		c.Writer.WriteHeader(211)
		return
	}

	var registerLogic routesinterface.IRegisterLogic
	registerLogic = &logics.RegisterLogic{
		RegisterData: &data.AuthInformationData{
			MongoData: &data.MongoData{
				SiteGroupsConfig: singletonSiteGroupsConfig,
			},
		},
		Config: singletonConfig,
	}

	registerLogic.SetUserName(authInfo.Username)
	registerLogic.SetPassword(authInfo.Password)
	registerLogic.SetSiteGroup(authInfo.SiteGroup)

	registerStatus := registerLogic.Register()

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
func authorize(c *gin.Context) {
	authToken, err := c.Cookie("auth-token")
	if err != nil {
		panic(err)
	}

	var authorizationLogic routesinterface.IAuthorizationLogic
	authorizationLogic = &logics.AuthorizationLogic{
		Config: singletonConfig,
	}

	authorizationLogic.SetAuthToken(authToken)
	authorizationStatus := authorizationLogic.Authorize()

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
