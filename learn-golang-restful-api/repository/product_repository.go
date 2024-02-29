package repository

import (
	"context"
	"database/sql"
	"learn-golang-restful-api/model"
)

type ProductRepository interface {
	Save(ctx context.Context, tx *sql.Tx, product model.Product) model.Product
	FindAll(ctx context.Context, tx *sql.Tx) []model.Product
	FindById(ctx context.Context, tx *sql.Tx, productId int) (model.Product, error)
	Update(ctx context.Context, tx *sql.Tx, product model.Product) model.Product
	Delete(ctx context.Context, tx *sql.Tx, productId int)
}
