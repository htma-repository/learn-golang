package controllers

import (
	"context"
	"fmt"
	"learn-golang-todolist/model"
	"learn-golang-todolist/repository"
	"learn-golang-todolist/views"
	"strconv"

	"net/http"
)

func InsertTodoHandler(writer http.ResponseWriter, req *http.Request) {
	db := ConnectionDB()
	todoRepository := repository.NewTodoRepository(db)
	ctx := context.Background()

	name := req.PostFormValue("name")
	description := req.PostFormValue("description")

	todo := model.Todo{
		Name:        name,
		Description: description,
	}

	var resultTodo model.Todo
	var err error

	switch req.Method {
	case http.MethodGet:
		err = views.TemplateGlobal.ExecuteTemplate(writer, "form.html", nil)
	case http.MethodPost:
		resultTodo, err = todoRepository.InsertTodo(ctx, todo)
	default:
		http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
	}

	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	if len(resultTodo.Name) > 0 {
		http.Redirect(writer, req, "/todos", http.StatusPermanentRedirect)
	}
}

func AllTodoHandler(writer http.ResponseWriter, request *http.Request) {
	db := ConnectionDB()
	todoRepository := repository.NewTodoRepository(db)
	ctx := context.Background()

	todos, err := todoRepository.FindAllTodo(ctx)

	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	err = views.TemplateGlobal.ExecuteTemplate(writer, "list-todo.html", map[string]any{
		"Todos": todos,
	})

	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
}

func TodoByIdHandler(writer http.ResponseWriter, request *http.Request) {
	db := ConnectionDB()
	todoRepository := repository.NewTodoRepository(db)
	ctx := context.Background()

	id := request.URL.Query().Get("id")

	fmt.Println("ID", id)

	i, err := strconv.Atoi(id)

	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	todo, err := todoRepository.FindTodoById(ctx, int32(i))

	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	if request.Method == http.MethodPost {
		err = views.TemplateGlobal.ExecuteTemplate(writer, "todo.html", map[string]any{
			"Name":        todo.Name,
			"Description": todo.Description,
		})
	}

	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
}

func DeleteTodoHandler(writer http.ResponseWriter, request *http.Request) {
	db := ConnectionDB()
	todoRepository := repository.NewTodoRepository(db)
	ctx := context.Background()

	id := request.URL.Query().Get("id")

	i, err := strconv.Atoi(id)

	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	err = todoRepository.DeleteTodo(ctx, int32(i))

	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	http.Redirect(writer, request, "/todos", http.StatusPermanentRedirect)
}

func UpdateTodoHandler(writer http.ResponseWriter, request *http.Request) {
	db := ConnectionDB()
	todoRepository := repository.NewTodoRepository(db)
	ctx := context.Background()

	id := request.URL.Query().Get("id")
	name := request.URL.Query().Get("name")
	desc := request.URL.Query().Get("desc")

	i, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	err = views.TemplateGlobal.ExecuteTemplate(writer, "update-form.html", map[string]any{
		"Name":        name,
		"Description": desc,
	})

	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	updatedName := request.PostFormValue("name")
	updatedDesc := request.PostFormValue("description")

	if len(updatedName) <= 0 || len(updatedDesc) <= 0 {
		fmt.Println("Empty Field")
		return
	}

	todo := model.Todo{
		Id:          int32(i),
		Name:        updatedName,
		Description: updatedDesc,
	}

	updatedTodo, err := todoRepository.UpdateTodo(ctx, todo)

	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	if len(updatedTodo.Name) > 0 {
		http.Redirect(writer, request, "/todos", http.StatusPermanentRedirect)
	}
}
