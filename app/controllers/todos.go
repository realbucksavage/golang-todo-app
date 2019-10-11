package controllers

import (
	"encoding/json"

	"github.com/revel/revel"

	"github.com/realbucksavage/golang-todos-app/app/models"
	"github.com/realbucksavage/golang-todos-app/app/services/todos"
)

type Todos struct {
	*revel.Controller
}

func (c Todos) GetAll() revel.Result {
	allTodos := todos.GetAllTodos()
	return c.RenderJSON(allTodos)
}

func (c Todos) Create() revel.Result {
	if todoItem, err := c.parseTodo(); err != nil {
		return c.RenderError(err)
	} else {
		todoItem.Validate(c.Validation)

		if c.Validation.HasErrors() {
			return c.RenderText("Invalid todo")
		}

		todoItem := todos.CreateTodo(todoItem)

		return c.RenderJSON(todoItem)
	}
}

func (c Todos) GetById(id int) revel.Result {
	todo, err := todos.GetTodoById(id)

	if err != nil {
		return c.NotFound("For ID %d", id)
	}

	return c.RenderJSON(todo)
}

func (c Todos) Delete(id int) revel.Result {
	_, err := todos.GetTodoById(id)

	if err != nil {
		return c.NotFound("For ID %d", id)
	}

	todos.DeleteTodo(id)

	return nil
}

func (c Todos) parseTodo() (models.Todo, error) {
	todoItem := models.Todo{}
	err := json.Unmarshal(c.Params.JSON, &todoItem)
	return todoItem, err
}
