package repository

import (
	"context"

	"github.com/hilmiikhsan/shopeefun-product-service/internal/module/shop/entity"
	"github.com/hilmiikhsan/shopeefun-product-service/internal/module/shop/ports"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

var _ ports.ShopRepository = &shopRepository{}

type shopRepository struct {
	db *sqlx.DB
}

func NewShopRepository(db *sqlx.DB) *shopRepository {
	return &shopRepository{
		db: db,
	}
}

func (r *shopRepository) CreateShop(ctx context.Context, req *entity.CreateShopRequest) (*entity.CreateShopResponse, error) {
	var resp = new(entity.CreateShopResponse)

	err := r.db.QueryRowContext(ctx, r.db.Rebind(queryInsertShop),
		req.UserId,
		req.Name,
		req.Description,
		req.Terms).Scan(&resp.Id)
	if err != nil {
		log.Error().Err(err).Any("payload", req).Msg("repository::CreateShop - Failed to create shop")
		return nil, err
	}

	return resp, nil
}

func (r *shopRepository) GetShop(ctx context.Context, req *entity.GetShopRequest) (*entity.GetShopResponse, error) {
	var resp = new(entity.GetShopResponse)

	err := r.db.QueryRowxContext(ctx, r.db.Rebind(queryGetShopById), req.Id).StructScan(resp)
	if err != nil {
		log.Error().Err(err).Any("payload", req).Msg("repository::GetShop - Failed to get shop")
		return nil, err
	}

	return resp, nil
}

func (r *shopRepository) DeleteShop(ctx context.Context, req *entity.DeleteShopRequest) error {
	_, err := r.db.ExecContext(ctx, r.db.Rebind(querySoftDeleteShop), req.Id, req.UserId)
	if err != nil {
		log.Error().Err(err).Any("payload", req).Msg("repository::DeleteShop - Failed to delete shop")
		return err
	}

	return nil
}

func (r *shopRepository) UpdateShop(ctx context.Context, req *entity.UpdateShopRequest) (*entity.UpdateShopResponse, error) {
	var resp = new(entity.UpdateShopResponse)

	err := r.db.QueryRowxContext(ctx, r.db.Rebind(queryUpdateShop),
		req.Name,
		req.Description,
		req.Terms,
		req.Id,
		req.UserId).Scan(&resp.Id)
	if err != nil {
		log.Error().Err(err).Any("payload", req).Msg("repository::UpdateShop - Failed to update shop")
		return nil, err
	}

	return resp, nil
}

func (r *shopRepository) GetShops(ctx context.Context, req *entity.ShopsRequest) (*entity.ShopsResponse, error) {
	type dao struct {
		TotalData int `db:"total_data"`
		entity.ShopItem
	}

	var (
		resp = new(entity.ShopsResponse)
		data = make([]dao, 0, req.Paginate)
	)
	resp.Items = make([]entity.ShopItem, 0, req.Paginate)

	err := r.db.SelectContext(ctx, &data, r.db.Rebind(queryGetAllShop),
		req.UserId,
		req.Paginate,
		req.Paginate*(req.Page-1),
	)
	if err != nil {
		log.Error().Err(err).Any("payload", req).Msg("repository::GetShops - Failed to get shops")
		return nil, err
	}

	if len(data) > 0 {
		resp.Meta.TotalData = data[0].TotalData
	}

	for _, d := range data {
		resp.Items = append(resp.Items, d.ShopItem)
	}

	resp.Meta.CountTotalPage(req.Page, req.Paginate, resp.Meta.TotalData)

	return resp, nil
}
