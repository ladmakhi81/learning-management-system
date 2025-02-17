package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	basebootstrap "github.com/ladmakhi81/learning-management-system/internal/base/bootstrap"
)

func main() {

	bootstrap := basebootstrap.NewBootstrap()
	if err := bootstrap.Apply(); err != nil {
		log.Fatalln(err)
	}

	server := gin.Default()
	apiServer := server.Group("/api")

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
