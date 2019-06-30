package api

import (
	"net/http"

	"github.com/halkn/echo-sample/interfaces/controllers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Run create a new instance of http server and route.
func Run(todoController *controllers.TodoController) {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/todos", func(c echo.Context) error {
		return todoController.GetTodos(c)
	})
	e.GET("/todos/:id", func(c echo.Context) error {
		return todoController.GetTodoByID(c)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
