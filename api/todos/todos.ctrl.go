package todos

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"github.com/realbucksavage/todos/database/models"
	"github.com/realbucksavage/todos/lib/common"
)

// Alias to Todo
type Todo = models.Todo

// Alias to common.JSON
type JSON = common.JSON

func list(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var todos []Todo
	if err := db.Order("id desc").Find(&todos).Error; err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
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

	type RequestBody struct {
		Title     string `json:"title" binding:"required"`
		Completed bool   `json:"completed"`
	}
	var r RequestBody

	if err := c.BindJSON(&r); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	todo := Todo{Title: r.Title, Completed: r.Completed}

	if err := db.Create(&todo).Error; err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, todo.Serialize())
}

func update(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")

	type RequestBody struct {
		Title     string `json:"title"`
		Completed bool   `json:"completed"`
	}
	var r RequestBody
	if err := c.BindJSON(&r); err != nil {
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

	if err := db.Save(&t).Error; err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

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

	if err := db.Delete(&t).Error; err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusNoContent)
}
