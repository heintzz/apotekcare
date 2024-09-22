package merchant

import (
	"context"
	"database/sql"
)

type repository struct {
	db *sql.DB
}

func newRepository(db *sql.DB) repository {
	return repository{
		db: db,
	}
}

func (r repository) updateMerchantProfile(ctx context.Context, merchant Merchant) (err error) {
	email := ctx.Value("AUTH_EMAIL")
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
	return 		
}