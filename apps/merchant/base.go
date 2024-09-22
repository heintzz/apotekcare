package merchant

import (
	"database/sql"
	"heintzz/ecommerce/internal/middleware"

	"github.com/go-chi/chi/v5"
)

func Run(router chi.Router, db *sql.DB) {
	repo := newRepository(db)
	svc := newService(repo)
	handler := newHandler(svc)

	router.Route("/v1/merchant", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(middleware.CheckToken)
			r.Use(middleware.VerifyMerchantRole)
			r.Put("/profile", handler.editMerchantHandler)		
		})
	})
}

