package repository

const (
	queryInsertShop = `
		INSERT INTO shops (
			user_id, 
			name, 
			description, 
			terms
		) VALUES (?, ?, ?, ?) RETURNING id
	`

	queryGetShopById = `
		SELECT 
			name, 
			description, 
			terms
		FROM shops
		WHERE id = ?
	`

	querySoftDeleteShop = `
		UPDATE shops
		SET 
			deleted_at = NOW()
		WHERE id = ? AND user_id = ?
	`

	queryUpdateShop = `
		UPDATE shops
		SET 
			name = ?, 
			description = ?, 
			terms = ?, 
			updated_at = NOW()
		WHERE id = ? AND user_id = ?
		RETURNING id
	`

	queryGetAllShop = `
		SELECT
			COUNT(id) OVER() as total_data,
			id,
			name
		FROM shops
		WHERE
			deleted_at IS NULL
			AND user_id = ?
		LIMIT ? OFFSET ?
	`
)
