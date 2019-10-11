package todos

import (
	"errors"
	"fmt"

	"github.com/realbucksavage/golang-todos-app/app/models"
)

var allTodos []models.Todo = generateSamples()

func generateSamples() []models.Todo {

	todos := []models.Todo{}

	for i := 1; i <= 10; i++ {
		t := models.Todo{
			TodoId:    i,
			Title:     fmt.Sprintf("Test Title %d", i),
			Completed: i%3 == 0,
		}
		todos = append(todos, t)
	}

	return todos
}

func GetTodoById(id int) (models.Todo, error) {
	for _, todo := range allTodos {
		if todo.TodoId == id {
			return todo, nil
		}
	}

	return models.Todo{}, errors.New("Todo not found")
}

func DeleteTodo(id int) {
	for i, todo := range allTodos {
		if todo.TodoId == id {
			allTodos = append(allTodos[:i], allTodos[i+1])
		}
	}
}

func CreateTodo(newTodo models.Todo) models.Todo {

	lastTodo := allTodos[len(allTodos)-1]

	todo := models.Todo{
		TodoId:    lastTodo.TodoId + 1,
		Title:     newTodo.Title,
		Completed: newTodo.Completed,
	}

	allTodos = append(allTodos, todo)

	return todo
}

func GetAllTodos() []models.Todo {
	return allTodos
}
