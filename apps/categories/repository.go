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

func (r repository) getCategories(ctx context.Context) (categories []Category, err error) {
	query := `
		SELECT id, name
		FROM categories
	`	
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var category Category
		err := rows.Scan(&category.Id, &category.Name)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return categories, nil
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