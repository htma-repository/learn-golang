package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"learn-golang-database/entity"
	"time"
)

type CustomerRepositoryImpl struct {
	DB *sql.DB
}

func NewCustomerRepository(db *sql.DB) CustomerRepository {
	return &CustomerRepositoryImpl{DB: db}
}

func (db *CustomerRepositoryImpl) InsertUser(ctx context.Context, customer entity.Customer) (entity.Customer, error) {
	ctxC, cancel := context.WithCancel(ctx)
	defer db.DB.Close()
	defer cancel()

	query := "INSERT INTO customer(name, email, address, balance, birth_date, married, rating) VALUES (?, ?, ?, ?, ?, ?, ?)"

	result, err := db.DB.ExecContext(ctxC, query, customer.Name, customer.Email, customer.Address, customer.Balance, customer.BirthDate, customer.Married, customer.Rating)

	if err != nil {
		fmt.Println(err.Error())
		return customer, err
	}

	lastId, err := result.LastInsertId()

	if err != nil {
		fmt.Println(err.Error())
		return customer, err
	}

	customer.Id = int32(lastId)
	return customer, nil
}

func (db *CustomerRepositoryImpl) FindUserById(ctx context.Context, id int32) (entity.Customer, error) {
	ctxC, cancel := context.WithCancel(ctx)
	defer db.DB.Close()
	defer cancel()

	query := "SELECT id, name, email, address, balance, birth_date, married, rating FROM customer WHERE id = ? LIMIT 1"

	rows, err := db.DB.QueryContext(ctxC, query, id)
	customer := entity.Customer{}

	if err != nil {
		fmt.Println(err.Error())
		return customer, err
	}

	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&customer.Id, &customer.Name, &customer.Email, &customer.Address, &customer.Balance, &customer.BirthDate, &customer.Married, &customer.Rating)
		if err != nil {
			fmt.Println(err.Error())
			return customer, err
		}
		return customer, nil
	} else {
		return customer, errors.New("customer not found")
	}
}

func (db *CustomerRepositoryImpl) FindAllUser(ctx context.Context) ([]entity.Customer, error) {
	ctxC, cancel := context.WithCancel(ctx)
	defer db.DB.Close()
	defer cancel()

	query := "SELECT id, name, email, address, balance, birth_date, married, rating FROM customer"

	rows, err := db.DB.QueryContext(ctxC, query)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	defer rows.Close()

	var customers []entity.Customer

	for rows.Next() {
		customer := entity.Customer{}

		var birthDateNullable sql.NullTime

		err := rows.Scan(&customer.Id, &customer.Name, &customer.Email, &customer.Address, &customer.Balance, &birthDateNullable, &customer.Married, &customer.Rating)

		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}

		if birthDateNullable.Valid {
			customer.BirthDate = birthDateNullable.Time
		} else {
			customer.BirthDate = time.Time{} // Set to default zero time for nil values
		}

		customers = append(customers, customer)

	}

	return customers, nil
}
