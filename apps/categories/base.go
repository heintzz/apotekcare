package categories

import (
	"database/sql"

	"github.com/go-chi/chi/v5"
)

func Run(router chi.Router, db *sql.DB) {
	repo := newRepository(db)
	svc := newService(repo)
	handler := newHandler(svc)

	router.Route("/v1/categories", func(r chi.Router) {
		r.Post("/", handler.addCategoryHandler)	
	})
}

