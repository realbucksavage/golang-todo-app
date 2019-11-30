package main

import (
	"log"
	"net/http"

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

	// Handle missing routes.
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code":     "RESOURCE_NOT_FOUND",
			"resource": c.Request.RequestURI,
		})
	})

	log.Fatal(r.Run(":8080"))
}
