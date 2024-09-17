package apps

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func RunServer(appPort string) {
	router := chi.NewRouter()

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})

	log.Println("Server running at port", appPort)
	if err := http.ListenAndServe(appPort, router); err != nil {
		panic(err)
	}
}