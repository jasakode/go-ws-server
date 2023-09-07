package main

import (
	"fmt"
	"go-ws-server/development/seed/seeds"
	"go-ws-server/src/config"
	"go-ws-server/src/models"
	"log"
)

func main() {
	errLoadEnv := config.LoadConfig()
	if errLoadEnv != nil {
		log.Fatal(errLoadEnv)
	}
	errC := models.DbConnect()
	if errC != nil {
		fmt.Println(errC)
	}

	fmt.Println("Running Seeder...")
	var err error
	err = seeds.Users(models.DB)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = seeds.Webhook(models.DB)
	if err != nil {
		fmt.Println(err)
		return
	}
}
