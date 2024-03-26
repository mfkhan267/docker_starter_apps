package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

var serviceName string

func main() {
	log.SetOutput(os.Stderr)
	serviceName = "silly-demo"
	if os.Getenv("SERVICE_NAME") != "" {
		serviceName = os.Getenv("SERVICE_NAME")
	}

	// Server
	log.Println("Starting server...")
	router := gin.New()
	router.GET("/", rootHandler)
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	router.Run(fmt.Sprintf(":%s", port))
}
