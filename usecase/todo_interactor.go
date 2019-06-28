package usecase

import (
	"github.com/halkn/echo-sample/entity"
)

// TodoInteractor is an interactor for Todo entity.
type TodoInteractor struct {
	TodoRepository TodoRepository
}

// FindAll returns All of todos.
func (interactor *TodoInteractor) FindAll() (entity.Todos, error) {
	return interactor.TodoRepository.FindAll()
}
