package auth

import "database/sql"

type repository struct {
	db *sql.DB
}

func newRepository(db *sql.DB) repository {
	return repository{
		db: db,
	}
}

func (repo repository) registerUser(auth Auth) (err error) {
	query := `
		INSERT INTO auth (email, password, role, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)		
	`

	statement, err := repo.db.Prepare(query)
	if err != nil {
		return
	}

	defer statement.Close()

	_, err = statement.Exec(auth.Email, auth.Password, auth.Role, auth.CreatedAt, auth.UpdateAt)
	return
}

func (repo repository) insertToUsersTable(email, fullname string) error {
	query := `
		INSERT INTO users (email, full_name)
		VALUES ($1, $2)
	`
	_, err := repo.db.Exec(query, email, fullname)
	return err
}

func (repo repository) insertToMerchantsTable(email, name, address string) error {
	query := `
		INSERT INTO merchants (email, name, address)
		VALUES ($1, $2, $3)
	`
	_, err := repo.db.Exec(query, email, name, address)
	return err
}


func (repo repository) getByEmail(email string) (auth Auth, err error) {
	query := `
		SELECT 
			id, email, password, 
			role, created_at, updated_at
		FROM 
			auth
		WHERE 
			email = $1
	`

	statement, err := repo.db.Prepare(query)
	if err != nil {
		return
	}

	defer statement.Close()

	row := statement.QueryRow(email)
	err = row.Scan(
		&auth.Id, &auth.Email, &auth.Password,
		&auth.Role, &auth.CreatedAt, &auth.UpdateAt,
	)
	
	return
}