package ports

import (
	"context"

	"github.com/hilmiikhsan/shopeefun-product-service/internal/module/product/entity"
)

type ProductRepository interface {
	CreateProduct(ctx context.Context, req *entity.CreateProductRequest) (*entity.CreateProductResponse, error)
	GetProduct(ctx context.Context, req *entity.GetProductRequest) (*entity.GetProductResult, error)
	GetProducts(ctx context.Context, req *entity.ProductRequest) (*entity.ProductsResponse, error)
	UpdateProduct(ctx context.Context, req *entity.UpdateProductRequest) (*entity.UpdateProductResponse, error)
	DeleteProduct(ctx context.Context, req *entity.DeleteProductRequest) error
}

type ProductService interface {
	CreateProduct(ctx context.Context, req *entity.CreateProductRequest) (*entity.CreateProductResponse, error)
	GetProduct(ctx context.Context, req *entity.GetProductRequest) (*entity.GetProductResponse, error)
	GetProducts(ctx context.Context, req *entity.ProductRequest) (*entity.ProductsResponse, error)
	UpdateProduct(ctx context.Context, req *entity.UpdateProductRequest) (*entity.UpdateProductResponse, error)
	DeleteProduct(ctx context.Context, req *entity.DeleteProductRequest) error
}
