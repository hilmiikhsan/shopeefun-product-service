package repository

const (
	queryInsertProduct = `
		INSERT INTO products (
			shop_id, 
			name, 
			description, 
			category,
			price,
			stock
		) VALUES (?, ?, ?, ?, ?, ?) RETURNING id, name
	`

	queryGetProductById = `
		SELECT
			p.id as product_id,
			p.name as product_name,
			p.description,
			p.category,
			p.price,
			p.stock,
			s.id as shop_id,
			s.name as shop_name,
			s.rating as shop_rating
		FROM products p
		JOIN shops s ON p.shop_id = s.id
		WHERE p.id = ?
	`

	queryGetProducts = `
	SELECT
		COUNT(id) OVER() as total_data,
		id,
		name,
		description,
		category,
		price,
		stock
	FROM products
	WHERE
		deleted_at IS NULL
	LIMIT ? OFFSET ?
	`

	queryUpdateProduct = `
		UPDATE products
		SET
			name = ?,
			description = ?,
			category = ?,
			price = ?,
			stock = ?
		WHERE id = ? AND shop_id = ?
		RETURNING id
	`

	queryDeleteProduct = `
		UPDATE products
		SET
			delete_at = NOW()
		WHERE id = ? AND shop_id = ?
	`
)
