package repository

import (
	"context"
	"strconv"

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
		resp  = new(entity.ProductsResponse)
		data  = make([]dao, 0, req.Paginate)
		query = queryGetProducts
	)

	resp.Items = make([]entity.ProductItem, 0, req.Paginate)
	var minPrice, maxPrice float64
	var rating int64
	var err error

	if req.MinPrice != "" {
		minPrice, err = strconv.ParseFloat(req.MinPrice, 64)
		if err != nil {
			log.Error().Err(err).Any("payload", req).Msg("repository::GetProducts - Failed to parse min price")
			return nil, err
		}
	}

	if req.MaxPrice != "" {
		maxPrice, err = strconv.ParseFloat(req.MaxPrice, 64)
		if err != nil {
			log.Error().Err(err).Any("payload", req).Msg("repository::GetProducts - Failed to parse max price")
			return nil, err
		}
	}

	if req.Rating != "" {
		rating, err = strconv.ParseInt(req.Rating, 10, 64)
		if err != nil {
			log.Error().Err(err).Any("payload", req).Msg("repository::GetProducts - Failed to parse rating")
			return nil, err
		}
	}

	if req.Category != "" {
		query += " AND category = :category"
	}

	if minPrice > 0 {
		query += " AND price >= :min_price"
	}

	if maxPrice > 0 {
		query += " AND price <= :max_price"
	}

	if req.Brand != "" {
		query += " AND brand = :brand"
	}

	if rating > 0 {
		query += " AND rating >= :rating"
	}

	if req.Name != "" {
		query += " AND name ILIKE '%' || :name || '%'"
	}

	query += " LIMIT :limit OFFSET :offset"

	query, args, err := sqlx.Named(query, map[string]interface{}{
		"limit":     req.Paginate,
		"offset":    req.Paginate * (req.Page - 1),
		"category":  req.Category,
		"min_price": req.MinPrice,
		"max_price": req.MaxPrice,
		"brand":     req.Brand,
		"rating":    req.Rating,
		"name":      req.Name,
	})
	if err != nil {
		log.Error().Err(err).Msg("repository::GetProducts - Failed to bind named query")
		return nil, err
	}

	query = r.db.Rebind(query)

	err = r.db.SelectContext(ctx, &data, query, args...)
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
