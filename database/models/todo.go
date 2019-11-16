package models

import (
	"github.com/jinzhu/gorm"
	"github.com/realbucksavage/todos/lib/common"
)

// Struct representation of Todo
type Todo struct {
	gorm.Model

	Title     string
	Completed bool
}

// Serialize to JSON
func (t *Todo) Serialize() common.JSON {
	return common.JSON{
		"id":        t.ID,
		"content":   t.Title,
		"completed": t.Completed,
	}
}

// Deserialize from JSON
func (t *Todo) Deserialize(m common.JSON) {
	t.ID = uint(m["id"].(float64))
	t.Title = m["content"].(string)
	t.Completed = m["completed"].(bool)
}
