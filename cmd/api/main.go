package main

import (
	"fmt"
	"log"

	baseconfig "github.com/ladmakhi81/learning-management-system/internal/base/config"
	basestorage "github.com/ladmakhi81/learning-management-system/internal/base/storage"
	"github.com/spf13/viper"
)

func main() {
	viper.AutomaticEnv()

	config := baseconfig.NewConfig()

	storage := basestorage.NewStorage(config)

	if err := storage.Connect(); err != nil {
		log.Fatalf("database not connected : %v", err)
	}

	fmt.Println("main function invoked")
}
