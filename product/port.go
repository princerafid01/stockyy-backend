package product

import (
	"ecommerce/domain"
	prdHandler "ecommerce/rest/handlers/product"
)

// these depends on me
type Service interface {
	prdHandler.Service //embedding
}

// I depends on them
type ProductRepo interface {
	Create(p domain.Product) (*domain.Product, error)
	Get(productID int) (*domain.Product, error)
	List(page, limit int64) ([]*domain.Product, error)
	Count() (int64, error)
	Delete(productID int) error
	Update(product domain.Product) (*domain.Product, error)
}
