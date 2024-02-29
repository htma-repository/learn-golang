package learn_golang_database

import (
	"context"
	"errors"
	"fmt"
	"testing"
)

func TestExecContextSQL(t *testing.T) {
	db := ConnectionDB()
	ctx := context.Background()
	ctxC, cancel := context.WithCancel(ctx)
	defer db.Close()
	defer cancel()

	query := "INSERT INTO customer(name, email, address, balance, birth_date, married, rating) VALUES (?, ?, ?, ?, ?, ?, ?)"

	_, err := db.ExecContext(ctxC, query, "Joko Tingkir", "joko@gmail.com", "Cipondoh", 5000000, nil, true, 90.50)

	if err != nil {
		fmt.Println(errors.New(err.Error()))
		panic(err)
	}

	fmt.Println("Success")
}
