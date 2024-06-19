package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sambasivareddy-ch/rest_api_go/db"
	"github.com/sambasivareddy-ch/rest_api_go/routes"
)

func main() {
	// Default() function creates an HTTP server instance
	httpServer := gin.Default()

	if err1 := db.InitDB(); err1 != nil {
		log.Fatal("Unable to Create the Database Instance")
	}

	// Basic GET request saying Hello
	httpServer.GET("/", rootRouteHandler)

	routes.RegisterRoutes(httpServer)

	// Listening at Port 8080
	if err := httpServer.Run(":8000"); err != nil {
		log.Fatalf(err.Error())
	}
}

func rootRouteHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Server saying Hello to everyone!!",
	})
}
