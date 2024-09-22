package users

import (
	"database/sql"
	"heintzz/ecommerce/internal/middleware"

	"github.com/go-chi/chi/v5"
)

func Run(router chi.Router, db *sql.DB) {
	repo := newRepository(db)
	svc := newService(repo)
	handler := newHandler(svc)

	router.Route("/v1/user", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(middleware.CheckToken)
			r.Use(middleware.VerifyRole("user"))
			r.Get("/profile", handler.getProfileHandler)
			r.Put("/profile", handler.updateProfileHandler)
		})		
	})
}

