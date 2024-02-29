package learn_golang_database

import (
	"context"
	"fmt"
	"testing"
)

var email = "hutama@gmail.com"
var password = "secretPassword"

func TestSQLInjection(t *testing.T) {
	db := ConnectionDB()
	ctx := context.Background()
	ctxCan, cancel := context.WithCancel(ctx)
	defer cancel()
	defer db.Close()

	rows, err := db.QueryContext(ctxCan, "SELECT email FROM user WHERE email = '"+email+"' AND password = '"+password+"' LIMIT 1;")

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	if rows.Next() {
		var email string
		err2 := rows.Scan(&email)

		if err2 != nil {
			panic(err2)
		}

		fmt.Println("Success get data")
		fmt.Println("email", email)
	} else {
		fmt.Println("Failed get data")
	}
}

func TestSQLInjectionFix(t *testing.T) {
	db := ConnectionDB()
	ctx := context.Background()
	ctxCan, cancel := context.WithCancel(ctx)
	defer cancel()
	defer db.Close()

	rows, err := db.QueryContext(ctxCan, "SELECT email FROM user WHERE email = ? AND password = ? LIMIT 1;", email, password)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	if rows.Next() {
		var email string
		err2 := rows.Scan(&email)

		if err2 != nil {
			panic(err2)
		}

		fmt.Println("Success get data")
		fmt.Println("email", email)
	} else {
		fmt.Println("Failed get data")
	}
}
