package main

import (
	"github.com/AloisCRR/jwt-api-notes/config"
	routes "github.com/AloisCRR/jwt-api-notes/routes"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {

	// Database connection
	config.Connect()

	// Router
	router := gin.New()

	// Endpoints connection
	routes.Routes(router)

	log.Fatal(router.Run())
}
