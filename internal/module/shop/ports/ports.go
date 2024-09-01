package ports

import (
	"context"

	"github.com/hilmiikhsan/shopeefun-product-service/internal/module/shop/entity"
)

type ShopRepository interface {
	CreateShop(ctx context.Context, req *entity.CreateShopRequest) (*entity.CreateShopResponse, error)
	GetShop(ctx context.Context, req *entity.GetShopRequest) (*entity.GetShopResponse, error)
	DeleteShop(ctx context.Context, req *entity.DeleteShopRequest) error
	UpdateShop(ctx context.Context, req *entity.UpdateShopRequest) (*entity.UpdateShopResponse, error)
	GetShops(ctx context.Context, req *entity.ShopsRequest) (*entity.ShopsResponse, error)
}

type ShopService interface {
	CreateShop(ctx context.Context, req *entity.CreateShopRequest) (*entity.CreateShopResponse, error)
	GetShop(ctx context.Context, req *entity.GetShopRequest) (*entity.GetShopResponse, error)
	DeleteShop(ctx context.Context, req *entity.DeleteShopRequest) error
	UpdateShop(ctx context.Context, req *entity.UpdateShopRequest) (*entity.UpdateShopResponse, error)
	GetShops(ctx context.Context, req *entity.ShopsRequest) (*entity.ShopsResponse, error)
}
