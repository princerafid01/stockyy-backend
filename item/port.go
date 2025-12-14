package item

import (
	"ecommerce/domain"
	itemHandler "ecommerce/rest/handlers/item"
)

// these depends on me
type Service interface {
	itemHandler.Service //embedding
}

// I depends on them
type ItemRepo interface {
	Create(item domain.Item) (*domain.Item, error)
	Get(itemID int64, userID int64) (*domain.Item, error)
	List(userID int64, page, limit int64) ([]*domain.Item, error)
	Count(userID int64) (int64, error)
	Delete(itemID int64, userID int64) error
	Update(item domain.Item) (*domain.Item, error)
}
