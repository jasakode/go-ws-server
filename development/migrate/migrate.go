package main

import (
	"fmt"
	"go-ws-server/src/config"
	"go-ws-server/src/models"
	"log"
)

func main() {

	errLoadEnv := config.LoadConfig()
	if errLoadEnv != nil {
		log.Fatal(errLoadEnv)
	}

	fmt.Println("init development")
	err := models.DbConnect()
	if err != nil {
		fmt.Println(err)
	}
}
