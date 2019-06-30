package registory

import (
	"database/sql"

	"github.com/halkn/echo-sample/infrastructures/datastore"
	"github.com/halkn/echo-sample/interfaces/controllers"
	"github.com/halkn/echo-sample/interfaces/gateway"
	"github.com/halkn/echo-sample/usecase"
)

func InjectDB() *sql.DB {
	return datastore.NewDB()
}

func InjectSqlHandler() *datastore.SQLHandler {
	return datastore.NewSQLHandler(InjectDB())
}

func InjectTodoRepository() *gateway.TodoRepository {
	return gateway.NewTodoRepository(InjectSqlHandler())
}

func InjectTodoInteractor() *usecase.TodoInteractor {
	return usecase.NewTodoInteractor(InjectTodoRepository())
}

func InjectTodoController() *controllers.TodoController {
	return controllers.NewTodoController(InjectTodoInteractor())
}
