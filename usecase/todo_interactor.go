package usecase

import (
	"github.com/halkn/echo-sample/entity"
)

// TodoInteractor is an interactor for Todo entity.
type TodoInteractor struct {
	TodoRepository
}

// GetAllTodos returns All of todos.
func (interactor *TodoInteractor) GetAllTodos() (entity.Todos, error) {
	return interactor.TodoRepository.FindAll()
}
