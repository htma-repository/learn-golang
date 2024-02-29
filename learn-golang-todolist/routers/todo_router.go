package routers

import (
	"net/http"

	"learn-golang-todolist/controllers"
)

func TodoRouter() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", controllers.InsertTodoHandler)

	// mux.HandleFunc("/add", controllers.InsertTodoHandler)

	mux.HandleFunc("/todos", controllers.AllTodoHandler)

	mux.HandleFunc("/todo", controllers.TodoByIdHandler)

	mux.HandleFunc("/todo-delete", controllers.DeleteTodoHandler)

	mux.HandleFunc("/todo-edit", controllers.UpdateTodoHandler)

	server := http.Server{Addr: "localhost:8080", Handler: mux}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
