package apps

import (
	"database/sql"
	"heintzz/apotekcare/apps/auth"
	"heintzz/apotekcare/apps/categories"
	"heintzz/apotekcare/apps/merchant"
	"heintzz/apotekcare/apps/products"
	"heintzz/apotekcare/apps/users"
	"heintzz/apotekcare/internal/middleware"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func registerRoute(router chi.Router, db *sql.DB) {
	auth.Run(router, db)
	users.Run(router, db)
	merchant.Run(router, db)
	categories.Run(router, db)
	products.Run(router, db)
}


func RunServer(appPort string, db *sql.DB) {
	router := chi.NewRouter()

	router.Use(middleware.Tracer)
	registerRoute(router, db)

	log.Println("Server running at port", appPort)
	if err := http.ListenAndServe(appPort, router); err != nil {
		panic(err)
	}
}