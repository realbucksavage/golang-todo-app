package database

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// Inject is a middleware function that injects the database reference to Gin's request context.
func Inject(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	}
}
