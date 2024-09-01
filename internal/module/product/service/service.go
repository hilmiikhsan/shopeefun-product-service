package service

import (
	"context"

	"github.com/rs/zerolog/log"

	"github.com/hilmiikhsan/shopeefun-product-service/internal/module/product/entity"
	"github.com/hilmiikhsan/shopeefun-product-service/internal/module/product/ports"
	shopEntity "github.com/hilmiikhsan/shopeefun-product-service/internal/module/shop/entity"
)

var _ ports.ProductService = &productService{}

type productService struct {
	repo ports.ProductRepository
}

func NewProductService(repo ports.ProductRepository) *productService {
	return &productService{
		repo: repo,
	}
}

func (s *productService) CreateProduct(ctx context.Context, req *entity.CreateProductRequest) (*entity.CreateProductResponse, error) {
	return s.repo.CreateProduct(ctx, req)
}

func (s *productService) GetProduct(ctx context.Context, req *entity.GetProductRequest) (*entity.GetProductResponse, error) {
	result, err := s.repo.GetProduct(ctx, req)
	if err != nil {
		log.Error().Err(err).Any("payload", req).Msg("service::GetProduct - Failed to get product")
		return nil, err
	}

	return &entity.GetProductResponse{
		Id:          result.Id,
		Name:        result.Name,
		Description: result.Description,
		Category:    result.Category,
		Price:       result.Price,
		Stock:       result.Stock,
		ShopDetail: shopEntity.ShopItem{
			Id:     result.ShopId,
			Name:   result.ShopName,
			Rating: result.ShopRating,
		},
	}, nil
}

func (s *productService) GetProducts(ctx context.Context, req *entity.ProductRequest) (*entity.ProductsResponse, error) {
	return s.repo.GetProducts(ctx, req)
}

func (s *productService) UpdateProduct(ctx context.Context, req *entity.UpdateProductRequest) (*entity.UpdateProductResponse, error) {
	return s.repo.UpdateProduct(ctx, req)
}

func (s *productService) DeleteProduct(ctx context.Context, req *entity.DeleteProductRequest) error {
	return s.repo.DeleteProduct(ctx, req)
}
