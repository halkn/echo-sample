package controllers

import (
	"net/http"

	"github.com/halkn/echo-sample/interfaces/gateway"
	"github.com/halkn/echo-sample/usecase"
)

// TodoController ...
type TodoController struct {
	Interactor usecase.TodoInteractor
}

// NewTodoController create TodoController instance.
func NewTodoController(sqlHandler gateway.SQLHandler) *TodoController {
	return &TodoController{
		Interactor: usecase.TodoInteractor{
			TodoRepository: &gateway.TodoRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

// GetTodos returns all of todos as JSON object.
func (controller *TodoController) GetTodos(c Context) error {
	todos, err := controller.Interactor.GetAllTodos()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.JSON(http.StatusOK, todos)

}
