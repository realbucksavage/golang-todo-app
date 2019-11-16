package todos

import "github.com/gin-gonic/gin"

// ApplyRoutes registers the group `/todos`
func ApplyRoutes(r *gin.RouterGroup) {
	api := r.Group("/todos")
	{
		api.POST("/", create)

		api.GET("/", list)
		api.GET("/:id", get)

		api.DELETE("/:id", remove)

		api.PATCH("/:id", update)
	}
}
