package products

import (
	"context"
	"database/sql"
	"heintzz/ecommerce/internal/constants"
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