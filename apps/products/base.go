package products

import (
	"database/sql"
	"heintzz/apotekcare/internal/constants"
	"heintzz/apotekcare/internal/middleware"

	"github.com/go-chi/chi/v5"
)

func Run(router chi.Router, db *sql.DB) {
	repo := newRepository(db)
	svc := newService(repo)
	handler := newHandler(svc)
	
	router.Route("/v1/products", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(middleware.CheckToken)		
			r.Get("/", handler.getProductsHandler)
			r.Get("/id/{id}", handler.getDetailProductHandler)
		})
		r.Group(func(r chi.Router) {
			r.Use(middleware.CheckToken)		
			r.Use(middleware.VerifyRole(constants.ROLE_MERCHANT))
			r.Post("/", handler.addProductHandler)					
			r.Get("/id/{id}/merchant", handler.getDetailProductByMerchant)
			r.Get("/merchant", handler.getProductsByMerchantHandler)					
		})		
	})	

	router.Route("/v1/orders/checkout", func(r chi.Router) {
		r.Use(middleware.CheckToken)		
		r.Post("/", handler.checkoutProductHandler)		
	})	
}

