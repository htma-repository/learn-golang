package learn_golang_database

import (
	"context"
	"errors"
	"fmt"
	"testing"
)

func TestQuery(t *testing.T) {
	db := ConnectionDB()
	ctx := context.Background()
	ctxC, cancel := context.WithCancel(ctx)
	defer db.Close()
	defer cancel()

	query := "SELECT id, name FROM customer"

	rows, err := db.QueryContext(ctxC, query)

	if err != nil {
		fmt.Println(errors.New(err.Error()))
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id uint32
		var name string
		err2 := rows.Scan(&id, &name)

		if err2 != nil {
			fmt.Println(errors.New(err.Error()))
			panic(err)
		}

		fmt.Println("id", id)
		fmt.Println("name", name)

	}

	fmt.Println("Success")
}
