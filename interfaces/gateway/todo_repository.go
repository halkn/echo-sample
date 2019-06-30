package gateway

import (
	"github.com/halkn/echo-sample/entity"
)

// TodoRepository is struct of todo.
type TodoRepository struct {
	SQLHandler
}

// NewTodoRepository is constructor that creates TodoRepository
func NewTodoRepository(handler SQLHandler) *TodoRepository {
	return &TodoRepository{handler}
}

// FindAll will return all recode of todo table.
func (repo *TodoRepository) FindAll() (todos entity.Todos, err error) {

	rows, err := repo.Query("select * from todo")
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

// FindByID will return todo whoes ID mathces
func (repo *TodoRepository) FindByID(id string) (todos entity.Todos, err error) {

	rows, err := repo.Query("select * from todo where ID = $1", id)
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
