package main

import (
	"fmt"

	"github.com/halkn/echo-sample/infrastructures/api"
	"github.com/halkn/echo-sample/infrastructures/datastore"
	"github.com/halkn/echo-sample/interfaces/controllers"
)

func main() {

	conn := datastore.NewDB()
	defer func() {
		if err := conn.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	todoController := controllers.NewTodoController(datastore.NewSQLHandler(conn))

	api.NewRouter(todoController)
}
