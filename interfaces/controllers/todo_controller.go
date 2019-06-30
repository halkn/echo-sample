package controllers

import (
	"net/http"

	"github.com/halkn/echo-sample/usecase"
)

// TodoController ...
type TodoController struct {
	usecase.TodoUsecase
}

// NewTodoController create TodoController instance.
func NewTodoController(usecase usecase.TodoUsecase) *TodoController {
	return &TodoController{usecase}
}

// GetTodos returns all of todos as JSON object.
func (controller *TodoController) GetTodos(c Context) error {
	todos, err := controller.TodoUsecase.GetAllTodos()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.JSON(http.StatusOK, todos)

}
