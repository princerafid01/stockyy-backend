package user

import (
	"ecommerce/config"
)

type Handler struct {
	svc Service
	cnf *config.Config
}

func NewHandler(cnf *config.Config, svc Service) *Handler {
	return &Handler{
		svc: svc,
		cnf: cnf,
	}
}
