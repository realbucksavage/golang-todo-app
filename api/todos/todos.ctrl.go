package todos

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/realbucksavage/todos/database/models"
	"github.com/realbucksavage/todos/lib/common"
	"net/http"
)

// Alias to Todo
type Todo = models.Todo

// Alias to common.JSON
type JSON = common.JSON

func list(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var todos []Todo
	if err := db.Order("id desc").Find(&todos).Error; err != nil {
		fmt.Printf("Cannot list: %v\n", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	length := len(todos)
	response := make([]JSON, length, length)

	for i := 0; i < length; i++ {
		response[i] = todos[i].Serialize()
	}

	c.JSON(http.StatusOK, response)
}

func get(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")

	var todo Todo
	if err := db.First(&todo, id).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, todo.Serialize())
}

func create(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	type requestBody struct {
		Title     string `json:"title" binding:"required"`
		Completed bool   `json:"completed"`
	}
	var r requestBody

	if err := c.BindJSON(&r); err != nil {
		fmt.Printf("Cannot bind JSON: %v\n", err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	todo := Todo{Title: r.Title, Completed: r.Completed}
	db.NewRecord(todo)
	db.Create(&todo)

	c.JSON(http.StatusCreated, todo.Serialize())
}

func update(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")

	type requestBody struct {
		Title     string `json:"title"`
		Completed bool   `json:"completed"`
	}
	var r requestBody
	if err := c.BindJSON(&r); err != nil {
		fmt.Printf("Cannot bind JSON: %v\n", err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var t Todo
	if err := db.First(&t, id).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	t.Completed = r.Completed
	if r.Title != "" {
		t.Title = r.Title
	}

	db.Save(&t)
	c.JSON(http.StatusOK, t.Serialize())
}

func remove(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")

	var t Todo
	if err := db.First(&t, id).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	db.Delete(&t)
	c.Status(http.StatusNoContent)
}
