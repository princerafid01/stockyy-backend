package item

import (
	"ecommerce/domain"
)

type service struct {
	itemRepo ItemRepo
}

func NewService(itemRepo ItemRepo) Service {
	return &service{
		itemRepo: itemRepo,
	}
}

func (svc *service) Create(item domain.Item) (*domain.Item, error) {
	itm, err := svc.itemRepo.Create(item)

	if err != nil {
		return nil, err
	}

	if itm == nil {
		return nil, nil
	}

	return itm, nil
}

func (svc *service) Delete(id int64, userID int64) error {
	err := svc.itemRepo.Delete(id, userID)

	if err != nil {
		return err
	}

	return nil
}

func (svc *service) Get(id int64, userID int64) (*domain.Item, error) {
	itm, err := svc.itemRepo.Get(id, userID)

	if err != nil {
		return nil, err
	}

	if itm == nil {
		return nil, nil
	}

	return itm, nil
}

func (svc *service) List(userID int64, page, limit int64) ([]*domain.Item, error) {
	items, err := svc.itemRepo.List(userID, page, limit)

	if err != nil {
		return nil, err
	}

	if items == nil {
		return nil, nil
	}

	return items, nil
}

func (svc *service) Count(userID int64) (int64, error) {
	itemCount, err := svc.itemRepo.Count(userID)

	if err != nil {
		return 0, err
	}

	if itemCount == 0 {
		return 0, nil
	}

	return itemCount, nil
}

func (svc *service) Update(item domain.Item) (*domain.Item, error) {
	itm, err := svc.itemRepo.Update(item)

	if err != nil {
		return nil, err
	}

	if itm == nil {
		return nil, nil
	}

	return itm, nil
}
