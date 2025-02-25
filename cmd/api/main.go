package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	basebootstrap "github.com/ladmakhi81/learning-management-system/internal/base/bootstrap"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {

	bootstrap := basebootstrap.NewBootstrap()
	if err := bootstrap.Apply(); err != nil {
		log.Fatalln(err)
	}

	server := gin.Default()
	apiServer := server.Group("/api")

	server.StaticFile("/swagger.yaml", "./docs/swagger.yaml")

	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL("/swagger.yaml")))

	container := bootstrap.GetContainer()
	container.Provide(func() *gin.RouterGroup {
		return apiServer
	})

	config := bootstrap.GetConfig()
	port := config.ServerConfig.ServerPort

	bootstrap.LoadModules()

	fmt.Printf("the server running at port %d \n", port)

	log.Fatalln(server.Run(fmt.Sprintf(":%d", port)))

}
