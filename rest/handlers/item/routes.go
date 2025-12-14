package item

import (
	middleware "ecommerce/rest/middlewares"
	"net/http"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle("GET /api/items", manager.With(
		http.HandlerFunc(h.GetItems),
		h.middlewares.AuthenticateJWT,
	))
	mux.Handle(
		"POST /api/items",
		manager.With(
			http.HandlerFunc(h.CreateItem),
			h.middlewares.AuthenticateJWT,
		),
	)
	mux.Handle("GET /api/items/{id}", manager.With(
		http.HandlerFunc(h.GetItem),
		h.middlewares.AuthenticateJWT,
	))
	mux.Handle(
		"PUT /api/items/{id}",
		manager.With(
			http.HandlerFunc(h.UpdateItem),
			h.middlewares.AuthenticateJWT,
		),
	)
	mux.Handle(
		"DELETE /api/items/{id}",
		manager.With(
			http.HandlerFunc(h.DeleteItem),
			h.middlewares.AuthenticateJWT,
		),
	)
}
