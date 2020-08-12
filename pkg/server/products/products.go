package products

import (
	"encoding/json"
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

// RegisterRoutes applies the routes for the products service.
func RegisterRoutes(router chi.Router) http.Handler {
	router.Get("/", GetProducts)
	router.Post("/", CreateProducts)
	router.Get("/{id}", GetProductItem)
	router.Patch("/{id}", UpdateProductItem)
	router.Delete("/{id}", DeleteProductItem)

	return router
}

// GetProducts returns a list of product items.
func GetProducts(w http.ResponseWriter, req *http.Request) {

	if chi.URLParam(req, "id") == "" {
		render.Respond(w, req, httputils.NotFoundResponse)
		return
	}

	product, err := json.Marshal(Product{1, "Toaster", 19.99, 100})

	if err != nil {
		render.JSON(w, req, httputils.ServerError(err))
		return
	}
	response := make(map[string]string)
	response["data"] = string(product)

	render.JSON(w, req, response)
}

// CreateProducts adds new product item/s.
func CreateProducts(w http.ResponseWriter, req *http.Request) {
	product, err := json.Marshal(Product{1, "Toaster", 19.99, 100})

	if err != nil {
		render.JSON(w, req, httputils.ServerError(err))
		return
	}
	response := make(map[string]string)
	response["data"] = string(product)

	render.JSON(w, req, response)
}

// GetProductItem return a single product item matching ID.
func GetProductItem(w http.ResponseWriter, req *http.Request) {
	product, err := json.Marshal(Product{1, "Toaster", 19.99, 100})

	if err != nil {
		render.JSON(w, req, httputils.ServerError(err))
		return
	}
	response := make(map[string]string)
	response["data"] = string(product)

	render.JSON(w, req, response)
}

// UpdateProductItem updates a product item (non idempotenent).
func UpdateProductItem(w http.ResponseWriter, req *http.Request) {
	product, err := json.Marshal(Product{1, "Toaster", 19.99, 100})

	if err != nil {
		render.JSON(w, req, httputils.ServerError(err))
		return
	}
	response := make(map[string]string)
	response["data"] = string(product)

	render.JSON(w, req, response)
}

// DeleteProductItem deletes a product item matching given ID.
func DeleteProductItem(w http.ResponseWriter, req *http.Request) {
	product, err := json.Marshal(Product{1, "Toaster", 19.99, 100})

	if err != nil {
		render.JSON(w, req, httputils.ServerError(err))
		return
	}
	response := make(map[string]string)
	response["data"] = string(product)

	render.JSON(w, req, response)
}
