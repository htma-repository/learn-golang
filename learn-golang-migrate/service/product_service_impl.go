package service

import (
	"context"
	"database/sql"
	"learn-golang-migrate/exception"
	"learn-golang-migrate/helper"
	"learn-golang-migrate/model"
	"learn-golang-migrate/model/web"
	"learn-golang-migrate/repository"

	"github.com/go-playground/validator/v10"
)

type ProductServiceImpl struct {
	ProductRepository repository.ProductRepository
	DB                *sql.DB
	validator         *validator.Validate
}

func NewProductService(repository repository.ProductRepository, db *sql.DB, validator *validator.Validate) ProductService {
	return &ProductServiceImpl{
		ProductRepository: repository,
		DB:                db,
		validator:         validator,
	}
}

func (service *ProductServiceImpl) CreateProduct(ctx context.Context, request web.ProductCreateRequest) web.ProductResponse {
	err := service.validator.Struct(request)
	helper.PanicError(err)

	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.TxRollbackCommit(tx)

	product := model.Product{
		Name:        request.Name,
		Description: request.Description,
	}

	product = service.ProductRepository.Save(ctx, tx, product)

	return web.ProductResponse(product)
}

func (service *ProductServiceImpl) FindAllProduct(ctx context.Context) []web.ProductResponse {
	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.TxRollbackCommit(tx)

	products := service.ProductRepository.FindAll(ctx, tx)

	var productsResponse []web.ProductResponse
	for _, product := range products {
		productsResponse = append(productsResponse, web.ProductResponse(product))
	}

	return productsResponse
}

func (service *ProductServiceImpl) FindProductById(ctx context.Context, productId int) web.ProductResponse {
	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.TxRollbackCommit(tx)

	product, err := service.ProductRepository.FindById(ctx, tx, productId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return web.ProductResponse(product)
}

func (service *ProductServiceImpl) UpdateProduct(ctx context.Context, request web.ProductUpdateRequest) web.ProductResponse {
	err := service.validator.Struct(request)
	helper.PanicError(err)

	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.TxRollbackCommit(tx)

	product, err := service.ProductRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	product.Name = request.Name
	product.Description = request.Description

	product = service.ProductRepository.Update(ctx, tx, product)

	return web.ProductResponse(product)
}

func (service *ProductServiceImpl) DeleteProduct(ctx context.Context, productId int) {
	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.TxRollbackCommit(tx)

	product, err := service.ProductRepository.FindById(ctx, tx, productId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.ProductRepository.Delete(ctx, tx, product.Id)
}
