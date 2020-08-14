package server

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/yohannesHL/go-marketplace/pkg/server/products"
)

// Start starts the server.
func Start(host string) {
	router := chi.NewRouter()

	router.Use(
		render.SetContentType(render.ContentTypeJSON), 
		middleware.AllowContentEncoding("application/json"),
		middleware.RequestID,
		middleware.Logger,
		middleware.RealIP,
		middleware.Recoverer,
	)

	router.Route("/products", func(r chi.Router) {
		products.RegisterRoutes(r)
	})

	http.ListenAndServe(host, router)
}
