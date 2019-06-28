package infrastructure

import (
	"net/http"

	"github.com/halkn/echo-sample/interfaces/controllers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// NewRouter create a new instance of http server and route.
func NewRouter() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	todoController := controllers.NewTodoController(NewSQLHandler())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/todos", func(c echo.Context) error {
		return todoController.Show(c)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
