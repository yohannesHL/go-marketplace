package products

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	httputils "github.com/yohannesHL/go-marketplace/pkg/httputils"
)

// Product dto object for a product .
type Product struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

type contextKey string

var productContextKey = contextKey("product")

// RegisterRoutes applies the routes for the products service.
func RegisterRoutes(router chi.Router) http.Handler {
	router.Get("/", ListProducts)
	router.Post("/", CreateProducts)

	router.Route("/{productId}", func(r chi.Router) {
		r.Use(ProductContext)
		r.Get("/", GetProductItem)
		r.Patch("/", UpdateProductItem)
		r.Delete("/", DeleteProductItem)
	})

	return router
}

// ProductContext handles loading product item with given ID.
// Returns default Empty response if product not found.
func ProductContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		var product *Product
		if chi.URLParam(req, "productId") == "" {
			render.Respond(w, req, httputils.NotFoundResponse)
			return
		}

		product = &Product{
			ID:       1,
			Name:     "Toaster",
			Price:    19.99,
			Quantity: 100,
		}

		ctx := context.WithValue(req.Context(), productContextKey, product)
		next.ServeHTTP(w, req.WithContext(ctx))
	})
}

// ListProducts returns a list of product items.
func ListProducts(w http.ResponseWriter, req *http.Request) {

	render.JSON(w, req, httputils.NewAPIResponse(products))
}

// CreateProducts adds new product item/s.
func CreateProducts(w http.ResponseWriter, req *http.Request) {
	product, err := json.Marshal(Product{1, "Toaster", 19.99, 100})

	if err != nil {
		render.JSON(w, req, httputils.ServerError(err))
		return
	}

	render.JSON(w, req, httputils.NewAPIResponse(product))
}

// GetProductItem return a single product item matching ID.
func GetProductItem(w http.ResponseWriter, req *http.Request) {

	product := req.Context().Value(productContextKey).(*Product)

	render.JSON(w, req, httputils.NewAPIResponse(product))
}

// UpdateProductItem updates a product item (non idempotenent).
func UpdateProductItem(w http.ResponseWriter, req *http.Request) {

	product := req.Context().Value(productContextKey).(*Product)

	render.JSON(w, req, httputils.NewAPIResponse(product))
}

// DeleteProductItem deletes a product item matching given ID.
func DeleteProductItem(w http.ResponseWriter, req *http.Request) {
	product := req.Context().Value(productContextKey).(*Product)

	render.JSON(w, req, httputils.NewAPIResponse(product))
}

var products = []*Product{
	{1, "Toaster", 19.99, 70},
	{1, "Kettle", 29.99, 30},
	{1, "Chair", 99.99, 80},
	{1, "Lamp", 19.99, 50},
}
