package products

import (
	"context"
	"database/sql"
	"errors"
	"heintzz/apotekcare/internal/constants"
)

type repository struct {
	db *sql.DB
}

func newRepository(db *sql.DB) repository {
	return repository{
		db: db,
	}
}

func (r repository) addNewProduct(ctx context.Context, product Product) (err error) {
	email, ok := ctx.Value(constants.AUTH_EMAIL).(string)
  if !ok || email == "" {
    return 
  }

	var merchantID int
	query := `SELECT id FROM merchants WHERE email = $1`
	err = r.db.QueryRow(query, email).Scan(&merchantID)
	if err != nil {
		if err == sql.ErrNoRows {
			return 
		}
		return err 
	}

	product.MerchantId = merchantID	
	
	insertQuery := `
		INSERT INTO products (name, image_url, price, stock, description, category_id, merchant_id) 
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`
	_, err = r.db.Exec(
		insertQuery, product.Name, product.ImageUrl, product.Price, product.Stock,
		product.Description, product.CategoryId, product.MerchantId, 
	)
	if err != nil {
		return 
	}

	return 
}

func (r repository) getProducts(ctx context.Context, queryParams string) (products []ProductResponse, err error) {
	query := `
		SELECT p.id, p.name, p.image_url, p.price,
		c.id category_id, c.name category_name,
		m.id merchant_id, m.name merchant_name, m.city merchant_city
		FROM products p 
		JOIN categories c
		ON p.category_id = c.id
		JOIN merchants m
		ON p.merchant_id = m.id
	`

	if queryParams != "" {
		query += `WHERE LOWER(p.name) LIKE LOWER($1) OR LOWER(c.name) LIKE LOWER($1)`
		queryParams = "%" + queryParams + "%"

		rows, err := r.db.QueryContext(ctx, query, queryParams)
		if err != nil {
			return nil, err
		}

		defer rows.Close()

		for rows.Next() {
			var product ProductResponse
			if err := rows.Scan(
				&product.Id, &product.Name, &product.ImageUrl, &product.Price, 
				&product.Category.Id, &product.Category.Name,
				&product.Merchant.Id, &product.Merchant.Name, &product.Merchant.City,
			); err != nil {
				return nil, err
			}
			products = append(products, product)
		}
	} else {	
		rows, err := r.db.QueryContext(ctx, query)
		if err != nil {
			return nil, err
		}
		
		defer rows.Close()

		for rows.Next() {
			var product ProductResponse
			if err := rows.Scan(
				&product.Id, &product.Name, &product.Price, &product.ImageUrl,
				&product.Category.Id, &product.Category.Name,
				&product.Merchant.Id, &product.Merchant.Name, &product.Merchant.City,
			); err != nil {
				return nil, err
			}
			products = append(products, product)
		}
	}

	return products, nil
}

func (r repository) getProductsByMerchant(ctx context.Context, queryParams string) (products []ProductResponse, err error) {
	email := ctx.Value(constants.AUTH_EMAIL)
	query := `
		SELECT p.id, p.name, p.image_url, p.price,
		c.id category_id, c.name category_name,
		m.id merchant_id, m.name merchant_name, m.city merchant_city
		FROM products p 
		JOIN categories c
		ON p.category_id = c.id
		JOIN merchants m
		ON p.merchant_id = m.id
		WHERE m.email = $1
	`

	if queryParams != "" {
		query += `AND LOWER(p.name) LIKE LOWER($2)`
		queryParams = "%" + queryParams + "%"

		rows, err := r.db.QueryContext(ctx, query, email, queryParams)
		if err != nil {
			return nil, err
		}

		defer rows.Close()

		for rows.Next() {
			var product ProductResponse
			if err := rows.Scan(
				&product.Id, &product.Name,  &product.ImageUrl, &product.Price, 
				&product.Category.Id, &product.Category.Name,
				&product.Merchant.Id, &product.Merchant.Name, &product.Merchant.City,
			); err != nil {
				return nil, err
			}
			products = append(products, product)
		}
	} else {	
		rows, err := r.db.QueryContext(ctx, query, email)
		if err != nil {
			return nil, err
		}
		
		defer rows.Close()

		for rows.Next() {
			var product ProductResponse
			if err := rows.Scan(
				&product.Id, &product.Name, &product.Price, &product.ImageUrl,
				&product.Category.Id, &product.Category.Name,
				&product.Merchant.Id, &product.Merchant.Name, &product.Merchant.City,
			); err != nil {
				return nil, err
			}
			products = append(products, product)
		}
	}

	return products, nil
}

func (r repository) getDetailProduct(ctx context.Context, productId string) (product DetailProduct, err error) {
	query := `
		SELECT 
		 p.id, p.name product_name, p.image_url, p.price, p.stock, p.description,
		 c.id category_id, c.name category_name,
		 m.id merchant_id, m.name merchant_name, m.city merchant_city
		FROM products p
		JOIN categories c
		ON p.category_id = c.id
		JOIN merchants m
		ON p.merchant_id = m.id
		WHERE p.id = $1
	`

	row := r.db.QueryRowContext(ctx, query, productId)
	err = row.Scan(
		&product.Id, &product.Name, &product.ImageUrl, &product.Price, &product.Stock, &product.Description,
		&product.Category.Id, &product.Category.Name, 
		&product.Merchant.Id, &product.Merchant.Name, &product.Merchant.City,
	)
	if err != nil {
		return
	}

	return product, nil
}

func (r repository) getDetailProductByMerchant(ctx context.Context, productId string) (product DetailProduct, err error) {
	email := ctx.Value(constants.AUTH_EMAIL)
	query := `
		SELECT 
		 p.id, p.name product_name, p.image_url, p.price, p.stock, p.description,
		 c.id category_id, c.name category_name,
		 m.id merchant_id, m.name merchant_name, m.city merchant_city
		FROM products p
		JOIN categories c
		ON p.category_id = c.id
		JOIN merchants m
		ON p.merchant_id = m.id
		WHERE p.id = $1 AND m.email = $2
	`

	row := r.db.QueryRowContext(ctx, query, productId, email)
	err = row.Scan(
		&product.Id, &product.Name, &product.ImageUrl, &product.Price, &product.Stock, &product.Description,
		&product.Category.Id, &product.Category.Name, 
		&product.Merchant.Id, &product.Merchant.Name, &product.Merchant.City,
	)
	if err != nil {
		return
	}

	return product, nil
}

func (r repository) getProductStock(ctx context.Context, req checkoutProductRequest) (stock int, err error) {	
	query := `
		SELECT stock 
		FROM products
		WHERE id = $1
	`
	
	err = r.db.QueryRowContext(ctx, query, req.ProductId).Scan(&stock)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, errors.New("product not found") 
		}
		return 0, err 
	}

	return stock, nil 
}

func (r repository) checkoutProduct(ctx context.Context, req checkoutProductRequest) (err error) {
	query := `
		UPDATE products
		SET stock = stock - $1
		WHERE id = $2
	`

	_, err = r.db.ExecContext(ctx, query, req.Quantity, req.ProductId)
	if err != nil {
		return err 
	}

	return
}