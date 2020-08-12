package server

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/yohannesHL/go-marketplace/pkg/server/products"
)

// Start starts the server.
func Start(host string) {
	router := chi.NewRouter()

	router.Route("/products", func(r chi.Router) {
		products.RegisterRoutes(r)
	})

	http.ListenAndServe(host, router)
}
