package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mfuadfakhruzaki/Jadwalin/config"
	"github.com/mfuadfakhruzaki/Jadwalin/routes"
)

func main() {
	config.LoadConfig()

	// Initialize Redis
	config.InitRedis()

	// Initialize the database
	config.InitDB()

	// Perform database migration
	config.MigrateDB()

	// Set up Gin router
	r := gin.Default()

	// Initialize routes
	routes.SetupRoutes(r)

	// Run the server
	port := ":8080" // Port where the app will run
	fmt.Println("Server running on port", port)
	if err := r.Run(port); err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}
