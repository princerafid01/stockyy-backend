package middleware

import "net/http"

type Middleware func(http.Handler) http.Handler

type Manager struct {
	globalMiddlewares []Middleware
}

func NewManger() *Manager {
	return &Manager{
		globalMiddlewares: make([]Middleware, 0),
	}
}

func (mngr *Manager) Use(middlewares ...Middleware) *Manager {
	mngr.globalMiddlewares = append(mngr.globalMiddlewares, middlewares...)
	return mngr
}

func (mngr *Manager) With(handler http.Handler, middlewares ...Middleware) http.Handler {
	h := handler

	for _, middleware := range middlewares {
		h = middleware(h)
	}

	return h
}

func (mngr *Manager) WrapMux(handler http.Handler) http.Handler {
	// [Preflight, Cors, Logger]
	// logger(cors(preflight(mux)))
	h := handler

	for _, middleware := range mngr.globalMiddlewares {
		h = middleware(h)
	}
	return h
}
