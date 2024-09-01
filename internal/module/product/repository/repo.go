package repository

import (
	"context"

	"github.com/hilmiikhsan/shopeefun-product-service/internal/module/product/entity"
	"github.com/hilmiikhsan/shopeefun-product-service/internal/module/product/ports"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

var _ ports.ProductRepository = &productRepository{}

type productRepository struct {
	db *sqlx.DB
}

func NewProductRepository(db *sqlx.DB) *productRepository {
	return &productRepository{
		db: db,
	}
}

func (r *productRepository) CreateProduct(ctx context.Context, req *entity.CreateProductRequest) (*entity.CreateProductResponse, error) {
	var resp = new(entity.CreateProductResponse)

	err := r.db.QueryRowContext(ctx, r.db.Rebind(queryInsertProduct),
		req.ShopId,
		req.Name,
		req.Description,
		req.Category,
		req.Price,
		req.Stock,
	).Scan(&resp.Id, &resp.Name)
	if err != nil {
		log.Error().Err(err).Any("payload", req).Msg("repository::CreateProduct - Failed to create product")
		return nil, err
	}

	return resp, nil
}

func (r *productRepository) GetProduct(ctx context.Context, req *entity.GetProductRequest) (*entity.GetProductResult, error) {
	var resp = new(entity.GetProductResult)

	err := r.db.QueryRowxContext(ctx, r.db.Rebind(queryGetProductById), req.Id).StructScan(resp)
	if err != nil {
		log.Error().Err(err).Any("payload", req).Msg("repository::GetProduct - Failed to get product")
		return nil, err
	}

	return resp, nil
}

func (r *productRepository) GetProducts(ctx context.Context, req *entity.ProductRequest) (*entity.ProductsResponse, error) {
	type dao struct {
		TotalData int `db:"total_data"`
		entity.ProductItem
	}

	var (
		resp = new(entity.ProductsResponse)
		data = make([]dao, 0, req.Paginate)
	)
	resp.Items = make([]entity.ProductItem, 0, req.Paginate)

	err := r.db.SelectContext(ctx, &data, r.db.Rebind(queryGetProducts),
		req.Paginate,
		req.Paginate*(req.Page-1),
	)
	if err != nil {
		log.Error().Err(err).Any("payload", req).Msg("repository::GetProducts - Failed to get products")
		return nil, err
	}

	if len(data) > 0 {
		resp.Meta.TotalData = data[0].TotalData
	}

	for _, d := range data {
		resp.Items = append(resp.Items, d.ProductItem)
	}

	resp.Meta.CountTotalPage(req.Page, req.Paginate, resp.Meta.TotalData)

	return resp, nil
}

func (r *productRepository) UpdateProduct(ctx context.Context, req *entity.UpdateProductRequest) (*entity.UpdateProductResponse, error) {
	var resp = new(entity.UpdateProductResponse)

	err := r.db.QueryRowxContext(ctx, r.db.Rebind(queryUpdateProduct),
		req.Name,
		req.Description,
		req.Category,
		req.Price,
		req.Stock,
	).Scan(&resp.Id)
	if err != nil {
		log.Error().Err(err).Any("payload", req).Msg("repository::UpdateProduct - Failed to update product")
		return nil, err
	}

	return resp, nil
}

func (r *productRepository) DeleteProduct(ctx context.Context, req *entity.DeleteProductRequest) error {
	_, err := r.db.ExecContext(ctx, r.db.Rebind(queryDeleteProduct), req.Id)
	if err != nil {
		log.Error().Err(err).Any("payload", req).Msg("repository::DeleteProduct - Failed to delete product")
		return err
	}

	return nil
}
