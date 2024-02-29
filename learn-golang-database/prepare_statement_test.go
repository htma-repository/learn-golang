package learn_golang_database

import (
	"context"
	"fmt"
	"strconv"
	"testing"
)

func TestPrepareStatement(t *testing.T) {
	db := ConnectionDB()
	ctx := context.Background()
	ctxC, cancel := context.WithCancel(ctx)
	defer db.Close()
	defer cancel()

	query := "INSERT INTO customer(name, email, address, balance, birth_date, married, rating) VALUES (?, ?, ?, ?, ?, ?, ?)"

	// If you want query sql with same query but different data to store, use prepare statement
	statement, err := db.PrepareContext(ctxC, query)

	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	defer statement.Close()

	for i := 0; i < 10; i++ {
		result, err2 := statement.ExecContext(ctxC, "John-"+strconv.Itoa(i), "john-"+strconv.Itoa(i)+"@gmail.com", "UnitedState", 1000000, "1994-12-19", true, 90.50)

		if err2 != nil {
			fmt.Println(err.Error())
			panic(err)
		}

		resultId, err3 := result.LastInsertId()

		if err3 != nil {
			fmt.Println(err.Error())
			panic(err)
		}

		fmt.Println("New Customer id, ", resultId)
	}
}
