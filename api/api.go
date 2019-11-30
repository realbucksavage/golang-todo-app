package api

import (
	"github.com/gin-gonic/gin"

	"github.com/realbucksavage/golang-todo-app/api/todos"
)

// ApplyRoutes registers the group `/api`
func ApplyRoutes(r *gin.Engine) {
	todosRoute := r.Group("/api")
	todos.ApplyRoutes(todosRoute)
}
