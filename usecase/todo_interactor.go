package usecase

import (
	"github.com/halkn/echo-sample/entity"
)

// TodoInteractor is an interactor for Todo entity.
type TodoInteractor struct {
	TodoRepository
}

// NewTodoInteractor is constructor that creates TodoInteractor
func NewTodoInteractor(repo TodoRepository) *TodoInteractor {
	return &TodoInteractor{repo}
}

// GetAllTodos returns All of todos.
func (interactor *TodoInteractor) GetAllTodos() (entity.Todos, error) {
	return interactor.FindAll()
}

// GetTodosByID returns todo whoes that ID mathces.
func (interactor *TodoInteractor) GetTodosByID(id string) (entity.Todos, error) {
	return interactor.FindByID(id)
}

// CreateTodo create a new todo.
func (interactor *TodoInteractor) CreateTodo(todo entity.Todo) (int, error) {
	return interactor.Store(todo)
}
