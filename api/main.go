package main

import (
	"api/db"
	"api/handler"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	middleware "github.com/oapi-codegen/gin-middleware"
)

func newGinServer(gameApi *handler.GameApi) *gin.Engine {
	swagger, err := handler.GetSwagger()
	if err != nil {
		log.Fatal("Error getting swagger")
	}
	swagger.Servers = nil // Remove the server information from Swagger

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	handler.RegisterHandlers(r, gameApi)
	r.Use(middleware.OapiRequestValidator(swagger))

	return r
}

func main() {
	if err := db.Init(); err != nil {
		log.Fatalf("Could not establish connection to database: %v", err)
	}

	gameApi := handler.NewApiHandler()
	server := newGinServer(gameApi)

	fmt.Println("listening on port: 8080")
	server.Run(":8080")
}
