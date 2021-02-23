package main

import (
	"github.com/gin-gonic/gin"

	"auth-service/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	logics "auth-service/logics"
	routes "auth-service/routes"
)

func main() {
	r := gin.Default()
	routes.Auth{
		R:                   r,
		AuthenticationLogic: logics.AuthenticationLogic{},
	}.New()

	docs.SwaggerInfo.Title = "Musk Daily API Documentation"
	docs.SwaggerInfo.Description = "Simple API descriptions for musk daily API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = ""
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")
}
