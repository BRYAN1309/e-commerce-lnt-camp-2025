package main

import (
	"backend/configs"
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")

	configs.SetUpDatabase()
}
