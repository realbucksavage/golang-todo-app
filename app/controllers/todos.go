package controllers

import (
	"github.com/realbucksavage/golang-todos-app/app/models"
	"github.com/revel/revel"

	"fmt"
)

var allTodos []models.Todo = getSampleTodos()

type Todos struct {
	*revel.Controller
}

func (c Todos) Index() revel.Result {
	return c.RenderJSON(allTodos)
}

func getSampleTodos() []models.Todo {

	todos := []models.Todo{}

	for i := 0; i < 25; i++ {
		t := models.Todo{}
		t.TodoId = i
		t.Title = fmt.Sprintf("Test Title %d", i)
		t.Completed = i%3 == 0

		todos = append(todos, t)
	}

	return todos
}
