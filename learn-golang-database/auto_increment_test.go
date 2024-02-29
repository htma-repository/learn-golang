package learn_golang_database

import (
	"context"
	"errors"
	"fmt"
	"testing"
)

func TestAutoIncrement(t *testing.T) {
	db := ConnectionDB()
	ctx := context.Background()
	ctxC, cancel := context.WithCancel(ctx)
	defer db.Close()
	defer cancel()

	query := "INSERT INTO customer(name, email, address, balance, birth_date, married, rating) VALUES (?, ?, ?, ?, ?, ?, ?)"

	result, err := db.ExecContext(ctxC, query, "John Doe", "john@gmail.com", "United State", 5000000, nil, true, 90.50)

	if err != nil {
		fmt.Println(errors.New(err.Error()))
		panic(err)
	}

	id, err2 := result.LastInsertId()

	if err2 != nil {
		fmt.Println(errors.New(err.Error()))
		panic(err2)
	}

	fmt.Println("new customer id, ", id)

	fmt.Println("Success")
}
