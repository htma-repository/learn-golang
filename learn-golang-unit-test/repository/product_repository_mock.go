package repository

import (
	"learn-golang-unit-test/entity"

	"github.com/stretchr/testify/mock"
)

type ProductRepositoryMock struct {
	Mock mock.Mock
}

func (repository *ProductRepositoryMock) FindById(id string) *entity.Product {
	args := repository.Mock.Called(id)

	if product, ok := args.Get(0).(entity.Product); ok {
		return &product
	}
	return nil
}

// func (repository *ProductRepositoryMock) FindAll(url string) []*entity.Product {
// 	args := repository.Mock.Called(url)

// 	if products, ok := args.Get(0).([]*entity.Product); ok {
// 		return products
// 	}

// 	return nil
// }
