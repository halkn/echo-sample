package usecase

import "github.com/halkn/echo-sample/entity"

// TodoUsecase have usecase for todo.
type TodoUsecase interface {
	GetAllTodos() (entity.Todos, error)
}
