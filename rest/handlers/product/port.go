package product

import "ecommerce/domain"

type Service interface {
	Create(domain.Product) (*domain.Product, error)
	Delete(id int) error
	Get(id int) (*domain.Product, error)
	Count() (int64, error)
	List(page, limit int64) ([]*domain.Product, error)
	Update(domain.Product) (*domain.Product, error)
}
