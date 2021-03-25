package main

import (
	"log"
	"os"
	"runtime/debug"

	"github.com/gin-gonic/gin"

	"auth-service/config"
	"auth-service/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	routes "auth-service/routes"
)

func main() {

	siteGroupsConfig := &config.SiteGroupsConfig{}
	config := &config.Config{}

	configInstance := *config.GetConfig()

	r := gin.Default()

	if configInstance.Environment.Type != "local" {
		r.Use(logger())
	}

	setupApiDoc(r, configInstance.Environment.Port)

	routes.AuthRoutes(r, siteGroupsConfig, config)

	r.Run(":" + configInstance.Environment.Port)
}

func setupApiDoc(r *gin.Engine, port string) {
	docs.SwaggerInfo.Title = "Auth Service API Documentation"
	docs.SwaggerInfo.Description = "Simple API descriptions for Auth Service"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:" + port
	docs.SwaggerInfo.BasePath = ""
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
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
