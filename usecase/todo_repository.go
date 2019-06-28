package usecase

import "github.com/halkn/echo-sample/entity"

// TodoRepository is an repository for Todo entity.
type TodoRepository interface {
	FindAll() (entity.Todos, error)
}
