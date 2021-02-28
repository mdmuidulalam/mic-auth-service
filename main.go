package main

import (
	"log"
	"os"
	"runtime/debug"

	"github.com/gin-gonic/gin"

	"auth-service/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	data "auth-service/data"
	logics "auth-service/logics"
	routes "auth-service/routes"
)

func main() {
	r := gin.Default()
	r.Use(logger())
	routes.Auth{
		R: r,
		AuthenticationLogic: &logics.AuthenticationLogic{
			AuthenticationData: &data.AuthInformationData{
				MongoData: &data.MongoData{},
			},
		},
		RegisterLogic: &logics.RegisterLogic{
			RegisterData: &data.AuthInformationData{
				MongoData: &data.MongoData{},
			},
		},
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

func logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		file, err := os.OpenFile("panic.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			log.Fatal(err)
		}
		log.SetOutput(file)

		defer func() {
			if err := recover(); err != nil {
				log.Println(err)
				log.Println("stacktrace from panic: \n" + string(debug.Stack()))
				c.Writer.WriteHeader(500)
			}
		}()

		c.Next()
	}
}
