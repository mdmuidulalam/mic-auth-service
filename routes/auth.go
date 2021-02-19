package routes

import "github.com/gin-gonic/gin"

type Auth struct {
	R *gin.Engine
}

func (this Auth) New() {
	this.R.POST("authenticate", this.Authenticate)
}

// @Summary Authenticate an user
// @Description It will authenticate an user from provided information
// @ID authenticate-user
// @Router /authenticate [post]
// @Accept json
// @Param username body string true "Username of the user"
// @Param password body string true "Password of the user"
// @Success 200 {boolean} boolean "The authentication is complete and access token provided in http only cookies"
func (this Auth) Authenticate(c *gin.Context) {
	c.Writer.WriteHeader(200)
}
