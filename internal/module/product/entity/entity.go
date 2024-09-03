package entity

import (
	shop "github.com/hilmiikhsan/shopeefun-product-service/internal/module/shop/entity"
	"github.com/hilmiikhsan/shopeefun-product-service/pkg/types"
)

type CreateProductRequest struct {
	ShopId string `json:"shop_id" validate:"uuid" db:"shop_id"`

	Name        string  `json:"name" validate:"required" db:"name"`
	Description string  `json:"description" validate:"required,max=255" db:"description"`
	Category    string  `json:"category" validate:"required" db:"category"`
	Price       float64 `json:"price" validate:"required" db:"price"`
	Stock       int     `json:"stock" validate:"required" db:"stock"`
}

type CreateProductResponse struct {
	Id   string `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

type GetProductRequest struct {
	Id string `validate:"uuid" db:"id"`
}

type GetProductResponse struct {
	Id          string        `json:"id" db:"product_id"`
	Name        string        `json:"name" db:"product_name"`
	Description string        `json:"description" db:"description"`
	Category    string        `json:"category" db:"category"`
	Price       float64       `json:"price" db:"price"`
	Stock       int           `json:"stock" db:"stock"`
	ShopDetail  shop.ShopItem `json:"shop_detail"`
}

type ProductItem struct {
	Id          string  `json:"id" db:"id"`
	Name        string  `json:"name" db:"name"`
	Description string  `json:"description" db:"description"`
	Category    string  `json:"category" db:"category"`
	Price       float64 `json:"price" db:"price"`
	Stock       int     `json:"stock" db:"stock"`
	Rating      int     `json:"rating" db:"rating"`
}

type ProductRequest struct {
	Page     int    `query:"page" validate:"required,min=1"`
	Paginate int    `query:"paginate" validate:"required,min=1,max=100"`
	Category string `query:"category" validate:"omitempty,alpha"`
	MinPrice string `query:"min_price" validate:"omitempty,numeric"`
	MaxPrice string `query:"max_price" validate:"omitempty,numeric"`
	Brand    string `query:"brand" validate:"omitempty,alpha"`
	Rating   string `query:"rating" validate:"omitempty,numeric"`
	Name     string `query:"name" validate:"omitempty"`
}

type ProductsResponse struct {
	Items []ProductItem `json:"items"`
	Meta  types.Meta    `json:"meta"`
}

func (r *ProductRequest) SetDefault() {
	if r.Page < 1 {
		r.Page = 1
	}

	if r.Paginate < 1 {
		r.Paginate = 10
	}
}

type UpdateProductRequest struct {
	ShopId string `prop:"shop_id" validate:"uuid" db:"shop_id"`

	Id          string  `params:"id" validate:"uuid" db:"id"`
	Name        string  `json:"name" validate:"required" db:"name"`
	Description string  `json:"description" validate:"required" db:"description"`
	Category    string  `json:"category" validate:"required" db:"category"`
	Price       float64 `json:"price" validate:"required" db:"price"`
	Stock       int     `json:"stock" validate:"required" db:"stock"`
}

type UpdateProductResponse struct {
	Id   string `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

type DeleteProductRequest struct {
	ShopId string `prop:"shop_id" validate:"uuid" db:"shop_id"`

	Id string `params:"id" validate:"uuid" db:"id"`
}
