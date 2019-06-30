package gateway

import (
	"github.com/halkn/echo-sample/entity"
)

// TodoRepository is struct of todo.
type TodoRepository struct {
	SqlHandler SQLHandler
}

// FindAll will return all recode of todo table.
func (repo *TodoRepository) FindAll() (todos entity.Todos, err error) {

	rows, err := repo.SqlHandler.Query("select * from todo")
	defer rows.Close()
	if err != nil {
		return
	}

	for rows.Next() {
		var todo entity.Todo
		rows.Scan(&todo.ID, &todo.Title, &todo.Content, &todo.Status)
		todos = append(todos, todo)
	}

	return
}
