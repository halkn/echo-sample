package usecase

import "github.com/halkn/echo-sample/entity"

// TodoUsecase have usecase for todo.
type TodoUsecase interface {
	GetAllTodos() (entity.Todos, error)
	GetTodosByID(string) (entity.Todos, error)
	CreateTodo(entity.Todo) (int, error)
}
