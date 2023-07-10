package main

import (
	"github.com/spades0/go-url-shortener/models"
	"github.com/spades0/go-url-shortener/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// Create a new router
	router := gin.Default()

	// Setup the router
	routes.SetupRouter(router)

	// Defer the closing of the database connection
	defer models.CloseDB()

	// Run the server
	router.Run()
}
