package learn_golang_database

import (
	"context"
	"errors"
	"fmt"
	"testing"
)

func TestDBTransaction(t *testing.T) {
	db := ConnectionDB()
	ctx := context.Background()
	ctxC, cancel := context.WithCancel(ctx)
	defer db.Close()
	defer cancel()

	tx, err := db.BeginTx(ctxC, nil)

	if err != nil {
		fmt.Println(errors.New(err.Error()))
		panic(err)
	}

	defer tx.Rollback()

	query := "INSERT INTO customer(name, email, address, balance, birth_date, married, rating) VALUES (?, ?, ?, ?, ?, ?, ?)"

	result, err := tx.ExecContext(ctxC, query, "Budi", "budi@gmail.com", "Tangerang", 1000000, "1980-05-17", true, 90.2)

	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	id, err := result.LastInsertId()

	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	fmt.Println("new customer id, ", id)

	if err = tx.Commit(); err != nil {
		fmt.Println(err.Error())
		panic(err)
	} else {
		fmt.Println("Commit tx")
	}

	fmt.Println("Success")
}
