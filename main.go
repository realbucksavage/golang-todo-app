package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/realbucksavage/todos/api"
	"github.com/realbucksavage/todos/database"
)

func main() {
	// Initialize a database connection
	db := database.InitDb()

	// Create a default Gin engine
	r := gin.Default()

	// Apply the database middleware
	r.Use(database.Inject(db))

	// Let the api package register it's own routes.
	// Gopher talk on code organization - https://www.youtube.com/watch?v=oL6JBUk6tj0
	api.ApplyRoutes(r)

	if err := r.Run(":8080"); err != nil {
		fmt.Printf("Cannot start server: %v\n", err)
	}
}
