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

	fmt.Println("Runnig Migration...")
	err := models.DbConnect()
	if err != nil {
		fmt.Println(err)
	}

	errMigrate := models.DB.AutoMigrate(&models.Users{})
	if errMigrate != nil {
		fmt.Println(errMigrate.Error())
	}
	errMigrate = models.DB.AutoMigrate(&models.Webhooks{})
	if errMigrate != nil {
		fmt.Println(errMigrate.Error())
	}
	errMigrate = models.DB.AutoMigrate(&models.User_historys{})
	if errMigrate != nil {
		fmt.Println(errMigrate.Error())
	}
}
