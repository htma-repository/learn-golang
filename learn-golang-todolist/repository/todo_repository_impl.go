package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"learn-golang-todolist/model"
)

type TodoRepositoryImpl struct {
	DB *sql.DB
}

func NewTodoRepository(db *sql.DB) TodoRepository {
	return &TodoRepositoryImpl{DB: db}
}

func (db *TodoRepositoryImpl) InsertTodo(ctx context.Context, todo model.Todo) (model.Todo, error) {
	ctxCancel, cancel := context.WithCancel(ctx)
	defer db.DB.Close()
	defer cancel()

	query := "INSERT INTO task(name, description) VALUES(?, ?)"

	result, err := db.DB.ExecContext(ctxCancel, query, todo.Name, todo.Description)

	if err != nil {
		fmt.Println(err.Error())
		return todo, err
	}

	lastId, err := result.LastInsertId()

	if err != nil {
		fmt.Println(err.Error())
		return todo, err
	}

	todo.Id = int32(lastId)
	return todo, nil
}

func (db *TodoRepositoryImpl) FindTodoById(ctx context.Context, id int32) (model.Todo, error) {

	ctxCancel, cancel := context.WithCancel(ctx)
	defer db.DB.Close()
	defer cancel()

	query := "SELECT name, description FROM task WHERE id = ? LIMIT 1"

	rows, err := db.DB.QueryContext(ctxCancel, query, id)
	todo := model.Todo{}

	if err != nil {
		fmt.Println(err.Error())
		return todo, err
	}

	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&todo.Name, &todo.Description)

		if err != nil {
			fmt.Println(err.Error())
			return todo, err
		}

		return todo, nil
	} else {
		return todo, errors.New("todo not found")
	}
}

func (db *TodoRepositoryImpl) FindAllTodo(ctx context.Context) ([]model.Todo, error) {
	ctxCancel, cancel := context.WithCancel(ctx)
	defer db.DB.Close()
	defer cancel()

	query := "SELECT id, name, description FROM task"

	rows, err := db.DB.QueryContext(ctxCancel, query)
	var todos []model.Todo

	if err != nil {
		fmt.Println(err.Error())
		return todos, err
	}

	defer rows.Close()

	for rows.Next() {
		todo := model.Todo{}

		err := rows.Scan(&todo.Id, &todo.Name, &todo.Description)

		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}

		todos = append(todos, todo)
	}
	return todos, nil
}

func (db *TodoRepositoryImpl) UpdateTodo(ctx context.Context, todo model.Todo) (model.Todo, error) {
	ctxCancel, cancel := context.WithCancel(ctx)
	defer db.DB.Close()
	defer cancel()

	query := "UPDATE task SET name = ?, description = ? WHERE id = ?"

	_, err := db.DB.ExecContext(ctxCancel, query, todo.Name, todo.Description, todo.Id)

	if err != nil {
		fmt.Println(err.Error())
		return todo, err
	}
	return todo, nil
}

func (db *TodoRepositoryImpl) DeleteTodo(ctx context.Context, id int32) error {
	ctxCancel, cancel := context.WithCancel(ctx)
	defer db.DB.Close()
	defer cancel()

	query := "DELETE FROM task WHERE id = ?"

	_, err := db.DB.ExecContext(ctxCancel, query, id)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}
