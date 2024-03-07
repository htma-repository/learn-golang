package repository

import (
	"context"
	"database/sql"
	"errors"
	"learn-golang-dependency-injection/helper"
	"learn-golang-dependency-injection/model"
)

type ProductRepositoryImpl struct {
}

func NewProductRepository() *ProductRepositoryImpl {
	return &ProductRepositoryImpl{}
}

func (p *ProductRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, product model.Product) model.Product {
	ctxC, cancel := context.WithCancel(ctx)
	defer cancel()

	query := "INSERT INTO product(name, description) VALUES(?, ?)"

	result, err := tx.ExecContext(ctxC, query, product.Name, product.Description)

	helper.PanicError(err)

	id, err := result.LastInsertId()

	helper.PanicError(err)

	product.Id = int(id)

	return product

}

func (p *ProductRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []model.Product {
	ctxC, cancel := context.WithCancel(ctx)
	defer cancel()

	query := "SELECT id, name, description FROM product"

	rows, err := tx.QueryContext(ctxC, query)

	helper.PanicError(err)

	defer rows.Close()

	products := []model.Product{}

	for rows.Next() {
		product := model.Product{}

		err := rows.Scan(&product.Id, &product.Name, &product.Description)

		helper.PanicError(err)

		products = append(products, product)
	}

	return products
}

func (p *ProductRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, productId int) (model.Product, error) {
	ctxC, cancel := context.WithCancel(ctx)
	defer cancel()

	query := "SELECT id, name, description FROM product WHERE id = ?"

	rows, err := tx.QueryContext(ctxC, query, productId)

	helper.PanicError(err)

	defer rows.Close()

	product := model.Product{}

	if rows.Next() {
		err := rows.Scan(&product.Id, &product.Name, &product.Description)
		helper.PanicError(err)
		return product, nil
	} else {
		return product, errors.New("product not found")
	}
}

func (p *ProductRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, product model.Product) model.Product {
	ctxC, cancel := context.WithCancel(ctx)
	defer cancel()

	query := "UPDATE product SET name = ?, description = ? WHERE id = ?"

	_, err := tx.ExecContext(ctxC, query, product.Name, product.Description, product.Id)

	helper.PanicError(err)

	return product
}

func (p *ProductRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, productId int) {
	ctxC, cancel := context.WithCancel(ctx)
	defer cancel()

	query := "DELETE FROM product WHERE id = ?"

	_, err := tx.ExecContext(ctxC, query, productId)

	helper.PanicError(err)
}
