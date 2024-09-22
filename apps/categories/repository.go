package categories

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

func (r repository) addNewCategory(ctx context.Context, category Category) (err error) {
	query := `
		INSERT INTO categories (name, created_at, updated_at)
		VALUES ($1, $2, $3)
	`

	statement, err := r.db.Prepare(query)
	if err != nil {
		return
	}

	_, err = statement.Exec(category.Name, category.CreatedAt, category.UpdatedAt)
	return 
}