package repository

import (
	"context"
	"fmt"
	learn_golang_database "learn-golang-database"
	"learn-golang-database/entity"
	"testing"
	"time"
)

func TestCustomerInsert(t *testing.T) {
	db := learn_golang_database.ConnectionDB()
	customerRepository := NewCustomerRepository(db)
	ctx := context.Background()

	layout := "2006-01-02 15:04:05"

	t2, _ := time.Parse(layout, "1994-12-19 00:00:00")

	customer := entity.Customer{
		Name:      "John",
		Email:     "john@gmail.com",
		Address:   "United State",
		Balance:   1000000,
		BirthDate: t2,
		Married:   true,
		Rating:    90.6,
	}

	result, err := customerRepository.InsertUser(ctx, customer)

	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	fmt.Println(result)
}

func TestCustomerFIndById(t *testing.T) {
	db := learn_golang_database.ConnectionDB()
	customerRepository := NewCustomerRepository(db)
	ctx := context.Background()

	customer, err := customerRepository.FindUserById(ctx, 28)

	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	fmt.Println(customer)
}

func TestCustomerFIndAll(t *testing.T) {
	db := learn_golang_database.ConnectionDB()
	customerRepository := NewCustomerRepository(db)
	ctx := context.Background()

	customers, err := customerRepository.FindAllUser(ctx)

	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	for _, customer := range customers {
		fmt.Println(customer)
	}

}
