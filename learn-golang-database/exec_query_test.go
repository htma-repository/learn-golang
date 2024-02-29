package learn_golang_database

import (
	"context"
	"errors"
	"fmt"
	"testing"
)

func TestExecQuerySQL(t *testing.T) {
	db := ConnectionDB()
	ctx := context.Background()
	ctxC, cancel := context.WithCancel(ctx)
	defer db.Close()
	defer cancel()

	query := "INSERT INTO customers(name) VALUES (?)"

	_, err := db.ExecContext(ctxC, query, "Tama")

	if err != nil {
		fmt.Println(errors.New(err.Error()))
		panic(err)
	}

	fmt.Println("Success")
}
