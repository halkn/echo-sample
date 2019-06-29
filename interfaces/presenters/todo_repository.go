package presenters

import (
	"github.com/halkn/echo-sample/entity"
	"github.com/halkn/echo-sample/interfaces/gateway"
)

// TodoRepository is struct of todo.
type TodoRepository struct {
	SqlHandler gateway.SQLHandler
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
