package merchant

import (
	"context"
	"database/sql"
	"fmt"
	"heintzz/apotekcare/internal/constants"
	"heintzz/apotekcare/internal/helper"
)

type repository struct {
	db *sql.DB
}

func newRepository(db *sql.DB) repository {
	return repository{
		db: db,
	}
}

func (r repository) getMerchantProfile(ctx context.Context) (merchant Merchant, err error) {
	email := ctx.Value(constants.AUTH_EMAIL)
	fmt.Println(email)
	query := `
		SELECT 
			id, email, name, img_url, city, address
		FROM 
			merchants
		WHERE email = $1
	`

	row := r.db.QueryRowContext(ctx, query, email)
	
	err = row.Scan(
		&merchant.Id,
		&merchant.Email,
		&merchant.Name,
		&merchant.ImageUrl,
		&merchant.City,
		&merchant.Address,
	)

	if err != nil {		
		if err == sql.ErrNoRows {
			err = helper.ErrMerchantNotFound
		}
		return
	}

	return
}

func (r repository) updateMerchantProfile(ctx context.Context, merchant Merchant) (err error) {
	email := ctx.Value(constants.AUTH_EMAIL)
	query := `
		UPDATE merchants
		SET 
			name = $2,
			img_url = $3,
			city = $4,
			address = $5
		WHERE email = $1
	`

	_, err = r.db.Exec(query, email, merchant.Name, merchant.ImageUrl, merchant.City, merchant.Address)
	if err != nil {		
		if err == sql.ErrNoRows {
			err = helper.ErrMerchantNotFound
		}
		return
	}


	return 		
}