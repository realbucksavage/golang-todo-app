package api

import (
	"github.com/gin-gonic/gin"
	"github.com/realbucksavage/todos/api/todos"
)

// ApplyRoutes registers the group `/api`
func ApplyRoutes(r *gin.Engine) {
	todosRoute := r.Group("/api")
	todos.ApplyRoutes(todosRoute)
}
