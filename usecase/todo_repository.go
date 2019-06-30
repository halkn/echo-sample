package usecase

import "github.com/halkn/echo-sample/entity"

// TodoRepository ...
type TodoRepository interface {
	FindAll() (entity.Todos, error)
	FindByID(string) (entity.Todos, error)
}
