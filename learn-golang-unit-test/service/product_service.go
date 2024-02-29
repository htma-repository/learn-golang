package service

import (
	"errors"
	"learn-golang-unit-test/entity"
	"learn-golang-unit-test/repository"
)

type ProductServiceInterface interface {
	GetById(id string) (*entity.Product, error)
	// GetAll(url string) ([]*entity.Product, error)
}

type ProductService struct {
	Repository repository.ProductRepository
}

func (service *ProductService) GetById(id string) (*entity.Product, error) {
	product := service.Repository.FindById(id)

	if product == nil {
		return nil, errors.New("product not found")
	} else {
		return product, nil
	}
}

// func (service *ProductService) GetAll(url string) ([]*entity.Product, error) {
// 	products := service.Repository.FindAll(url)

// 	if products == nil {
// 		return nil, errors.New("products not found")
// 	} else {
// 		return products, nil
// 	}
// }
