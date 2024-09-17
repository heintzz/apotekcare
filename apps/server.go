package apps

import (
	"database/sql"
	"heintzz/ecommerce/apps/auth"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func registerRoute(router chi.Router, db *sql.DB) {
	auth.Run(router, db)
}


func RunServer(appPort string, db *sql.DB) {
	router := chi.NewRouter()

	registerRoute(router, db)

	log.Println("Server running at port", appPort)
	if err := http.ListenAndServe(appPort, router); err != nil {
		panic(err)
	}
}