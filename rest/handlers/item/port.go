package item

import "ecommerce/domain"

type Service interface {
	Create(domain.Item) (*domain.Item, error)
	Delete(id int64, userID int64) error
	Get(id int64, userID int64) (*domain.Item, error)
	Count(userID int64) (int64, error)
	List(userID int64, page, limit int64) ([]*domain.Item, error)
	Update(domain.Item) (*domain.Item, error)
}
