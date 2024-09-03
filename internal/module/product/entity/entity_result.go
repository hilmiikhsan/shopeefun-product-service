package entity

type GetProductResult struct {
	Id          string  `db:"product_id"`
	Name        string  `db:"product_name"`
	Description string  `db:"description"`
	Category    string  `db:"category"`
	Price       float64 `db:"price"`
	Stock       int     `db:"stock"`
	ShopId      string  `db:"shop_id"`
	ShopName    string  `db:"shop_name"`
	ShopRating  int     `db:"shop_rating"`
}
