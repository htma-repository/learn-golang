package repository

import (
	"context"
	"learn-golang-database/entity"
)

type CustomerRepository interface {
	InsertUser(ctx context.Context, customer entity.Customer) (entity.Customer, error)
	FindUserById(ctx context.Context, id int32) (entity.Customer, error)
	FindAllUser(ctx context.Context) ([]entity.Customer, error)
}
