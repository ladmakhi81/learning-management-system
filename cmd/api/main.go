package main

import (
	"fmt"

	basebootstrap "github.com/ladmakhi81/learning-management-system/internal/base/bootstrap"
)

func main() {
	bootstrap := basebootstrap.NewBootstrap()
	bootstrap.Apply()
	fmt.Println("main function invoked")
}
