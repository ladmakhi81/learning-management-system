package main

import (
	"fmt"
	"log"

	basestorage "github.com/ladmakhi81/learning-management-system/internal/base/storage"
)

func main() {
	storage := basestorage.NewStorage()
	if err := storage.Connect(); err != nil {
		log.Fatalf("database not connected : %v", err)
	}

	fmt.Println("main function invoked")
}
