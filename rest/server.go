package rest

import (
	"ecommerce/config"
	"ecommerce/rest/handlers/product"
	"ecommerce/rest/handlers/user"
	"ecommerce/rest/middlewares"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

type Server struct {
	cnf            *config.Config
	productHandler *product.Handler
	userHandler    *user.Handler
}

func NewServer(cnf *config.Config, productHandler *product.Handler, userHandler *user.Handler) *Server {
	return &Server{
		cnf,
		productHandler,
		userHandler,
	}
}

func (server *Server) Start() {
	manager := middleware.NewManger()
	manager.Use(middleware.Preflight, middleware.Cors, middleware.Logger)

	// router
	mux := http.NewServeMux()
	wrappedMux := manager.WrapMux(mux)

	// initRoutes(mux, manager)
	server.productHandler.RegisterRoutes(mux, manager)
	server.userHandler.RegisterRoutes(mux, manager)

	addr := ":" + strconv.Itoa(server.cnf.HttpPort)
	fmt.Println("Server running on port ", addr)
	err := http.ListenAndServe(addr, wrappedMux) // "Failed to start server"
	if err != nil {
		fmt.Println("Error starting server ", err)
		os.Exit(1)
	}
}
