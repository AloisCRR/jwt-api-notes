package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"projects/mongodb-notes-api/config"
	"projects/mongodb-notes-api/routes"
)

func main() {

	// Database connection
	config.Connect()

	// Router
	router := gin.Default()

	// Endpoints connection
	routes.Routes(router)

	log.Fatal(router.Run())
}
