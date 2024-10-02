package users

import (
	"context"
	"database/sql"
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

func (r repository) fetchProfile(ctx context.Context) (user User, err error) {
	email := ctx.Value(constants.AUTH_EMAIL)
	query := `
		SELECT id, email, full_name, gender, address, phone_number, created_at, updated_at
		FROM
			users
		WHERE
			email = $1
	`
	row := r.db.QueryRowContext(ctx, query, email)

	err = row.Scan(
		&user.Id, &user.Email, &user.FullName, &user.Gender, 
		&user.Address, &user.PhoneNumber, &user.CreatedAt, &user.UpdateAt,
	)
	if err != nil {
		return
	}
	return
}

func (r repository) updateProfile(ctx context.Context, user User) (err error) {
	email := ctx.Value(constants.AUTH_EMAIL)
	query := `
		UPDATE users
		SET 
			full_name = $2,
			gender = $3,
			address = $4,
			phone_number = $5
		WHERE email = $1
	`

	_, err = r.db.Exec(query, email, user.FullName, user.Gender, user.Address, user.PhoneNumber)
	return 	
}