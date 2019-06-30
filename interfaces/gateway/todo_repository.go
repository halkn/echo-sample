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

// Store add a new todo recode.
func (repo *TodoRepository) Store(todo entity.Todo) (int, error) {

	_, err := repo.Exec("insert into todo (ID, Title, Content ,Status) values ($1,$2,$3,$4)", todo.ID, todo.Title, todo.Content, todo.Status)
	// _, err := repo.Exec("insert into todo (ID, Title, Content ,Status) values (4,'todo4','todo-sample','0')")
	if err != nil {
		return 1, err
	}

	return 0, nil
}
