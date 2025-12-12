package product

import (
	"ecommerce/rest/middlewares"
	"net/http"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle("GET /api/products", manager.With(http.HandlerFunc(h.GetProducts)))
	mux.Handle(
		"POST /api/products",
		manager.With(
			http.HandlerFunc(h.CreateProducts),
			h.middlewares.AuthenticateJWT,
		),
	)
	mux.Handle("GET /api/products/{id}", manager.With(http.HandlerFunc(h.GetProduct)))
	mux.Handle(
		"PUT /api/products/{id}",
		manager.With(
			http.HandlerFunc(h.UpdateProduct),
			h.middlewares.AuthenticateJWT,
		),
	)
	mux.Handle(
		"DELETE /api/products/{id}",
		manager.With(
			http.HandlerFunc(h.DeleteProduct),
			h.middlewares.AuthenticateJWT,
		),
	)

}
