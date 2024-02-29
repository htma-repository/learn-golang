package learn_golang_database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"testing"
	"time"
)

func TestQueryComplex(t *testing.T) {
	db := ConnectionDB()
	ctx := context.Background()
	ctxC, cancel := context.WithCancel(ctx)
	defer db.Close()
	defer cancel()

	query := "SELECT id, name, email, address, balance, birth_date, married, rating, created_at FROM customer"

	rows, err := db.QueryContext(ctxC, query)

	if err != nil {
		fmt.Println(errors.New(err.Error()))
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, balance int32
		var name, email, address string
		var rating float64
		var married bool
		// use sql.Null type if db can store NULL data
		var birthDate sql.NullTime
		var createdAt time.Time

		err2 := rows.Scan(&id, &name, &email, &address, &balance, &birthDate, &married, &rating, &createdAt)

		if err2 != nil {
			fmt.Println("Error occurred")
			fmt.Println(errors.New(err.Error()))
			panic(err)
		}

		fmt.Println("============================")
		fmt.Println("id", id)
		fmt.Println("name", name)
		fmt.Println("email", email)
		fmt.Println("address", address)
		fmt.Println("balance", balance)
		// Check from db data if db can store NULL data
		if birthDate.Valid {
			fmt.Println("birth_date", birthDate.Time)
		} else {
			fmt.Println("birth_date", nil)
		}
		fmt.Println("married", married)
		fmt.Println("rating", rating)
		fmt.Println("created_at", createdAt)

	}

	fmt.Println("Success")
}
