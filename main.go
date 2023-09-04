package main

import (
	"fmt"
	_ "go-ws-server/src/docs"
	"go-ws-server/src/middleware"
	"go-ws-server/src/modules"
	"log"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(middleware.CorsMiddleware())

	r.GET("/ws", func(c *gin.Context) { middleware.HandleWebSocket(c.Writer, c.Request) })

	apiV1 := r.Group("/api/v1")
	modules.Routers(apiV1)

	// DOCS
	url := ginSwagger.URL("http://localhost:8080/api/v1/docs/doc.json") // The url pointing to API definition
	apiV1.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	fmt.Println("Server Running Successfully...")
	fmt.Println("See Docs http://localhost:8080/api/v1/docs/index.html")
	log.Fatal(r.Run(":8080"))
}
