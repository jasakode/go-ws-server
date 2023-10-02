package main

import (
	"fmt"
	"go-ws-server/src/config"
	_ "go-ws-server/src/docs"
	"go-ws-server/src/middleware"
	"go-ws-server/src/models"
	"go-ws-server/src/modules"
	"log"

	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	errLoadEnv := config.LoadConfig()
	if errLoadEnv != nil {
		log.Fatal(errLoadEnv)
	}

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(middleware.CorsMiddleware())
	errDb := models.DbConnect()
	if errDb != nil {
		log.Fatal(errDb.Error())
	}

	apiV1 := r.Group(config.Config.API_VERSION)
	modules.Routers(apiV1)

	// ...

	// DOCS
	url := ginSwagger.URL(config.Config.SERVER + config.Config.PORT + "/api/v1/docs/doc.json")
	apiV1.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	// red := color.New(color.FgRed)
	green := color.New(color.FgGreen)
	blue := color.New(color.FgBlue)
	runningText := fmt.Sprintf("%v Server Running Successfully...", green.Sprintf("[+]"))
	doctText := fmt.Sprintf("%v See Docs %v", green.Sprintf("[+]"), blue.Sprintf("%v%v%v/docs/index.html", config.Config.SERVER, config.Config.PORT, config.Config.API_VERSION))
	fmt.Println(runningText)
	fmt.Println(doctText)
	fmt.Println(fmt.Sprintf("%v %v%v", green.Sprintf("[+]"), config.Config.SERVER, config.Config.API_VERSION))
	log.Fatal(r.Run(fmt.Sprintf("%v", config.Config.PORT)))
}
