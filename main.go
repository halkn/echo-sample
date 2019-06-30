package main

import (
	"fmt"

	"github.com/halkn/echo-sample/infrastructures/api"
	"github.com/halkn/echo-sample/registory"
)

func main() {

	// conn := datastore.NewDB()
	conn := registory.InjectDB()
	defer func() {
		if err := conn.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	todoController := registory.InjectTodoController()

	api.NewRouter(todoController)
}
