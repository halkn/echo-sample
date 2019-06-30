package controllers

import (
	"net/http"

	"github.com/halkn/echo-sample/entity"
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

// GetTodoByID return todos whoes ID mathces
func (controller *TodoController) GetTodoByID(c Context) error {
	todo, err := controller.TodoUsecase.GetTodosByID(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, todo)

}

// Create create a new todo.
func (controller *TodoController) Create(c Context) error {
	u := entity.Todo{}
	c.Bind(&u)

	todo, err := controller.TodoUsecase.CreateTodo(u)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, todo)

}
