package service

import (
	"context"
	"learn-golang-restful-api/model/web"
)

type ProductService interface {
	CreateProduct(ctx context.Context, request web.ProductCreateRequest) web.ProductResponse
	FindAllProduct(ctx context.Context) []web.ProductResponse
	FindProductById(ctx context.Context, productId int) web.ProductResponse
	UpdateProduct(ctx context.Context, request web.ProductUpdateRequest) web.ProductResponse
	DeleteProduct(ctx context.Context, productId int)
}
