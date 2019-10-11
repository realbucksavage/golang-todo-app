package models

import "github.com/revel/revel"

type Todo struct {
	TodoId    int
	Title     string
	Completed bool
}

func (todo *Todo) Validate(v *revel.Validation) {

	v.Required(todo.Title)
	v.MaxSize(todo.Title, 100)
}
