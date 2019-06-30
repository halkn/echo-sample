package registory

import (
	"database/sql"

	"github.com/halkn/echo-sample/infrastructures/datastore"
	"github.com/halkn/echo-sample/interfaces/controllers"
	"github.com/halkn/echo-sample/interfaces/gateway"
	"github.com/halkn/echo-sample/usecase"
)

// InjectDB is Dependency injector for *sql.DB
func InjectDB() *sql.DB {
	return datastore.NewDB()
}

// InjectSQLHandler is Dependency injector for SQLHandler
func InjectSQLHandler() *datastore.SQLHandler {
	return datastore.NewSQLHandler(InjectDB())
}

// InjectTodoRepository is Dependency injector for TodoRepository
func InjectTodoRepository() *gateway.TodoRepository {
	return gateway.NewTodoRepository(InjectSQLHandler())
}

// InjectTodoInteractor is Dependency injector for TodoInteractor
func InjectTodoInteractor() *usecase.TodoInteractor {
	return usecase.NewTodoInteractor(InjectTodoRepository())
}

// InjectTodoController is Dependency injector for TodoController
func InjectTodoController() *controllers.TodoController {
	return controllers.NewTodoController(InjectTodoInteractor())
}
