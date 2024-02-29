package repository

import "learn-golang-unit-test/entity"

type ProductRepository interface {
	FindById(id string) *entity.Product
	// FindAll(url string) []*entity.Product
}
