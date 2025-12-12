package product

import (
	"ecommerce/domain"
)

type service struct {
	prdRepo ProductRepo
}

func NewService(prdRepo ProductRepo) Service {
	return &service{
		prdRepo: prdRepo,
	}
}

func (svc *service) Create(p domain.Product) (*domain.Product, error) {
	prd, err := svc.prdRepo.Create(p)

	if err != nil {
		return nil, err
	}

	if prd == nil {
		return nil, nil
	}

	return prd, nil
}
func (svc *service) Delete(id int) error {
	err := svc.prdRepo.Delete(id)

	if err != nil {
		return err
	}

	return nil
}

func (svc *service) Get(id int) (*domain.Product, error) {
	prd, err := svc.prdRepo.Get(id)

	if err != nil {
		return nil, err
	}

	if prd == nil {
		return nil, nil
	}

	return prd, nil
}

func (svc *service) List(page, limit int64) ([]*domain.Product, error) {
	prds, err := svc.prdRepo.List(page, limit)

	if err != nil {
		return nil, err
	}

	if prds == nil {
		return nil, nil
	}

	return prds, nil
}

func (svc *service) Count() (int64, error) {
	prdCount, err := svc.prdRepo.Count()

	if err != nil {
		return 0, err
	}

	if prdCount == 0 {
		return 0, nil
	}

	return prdCount, nil
}

func (svc *service) Update(p domain.Product) (*domain.Product, error) {
	prd, err := svc.prdRepo.Update(p)

	if err != nil {
		return nil, err
	}

	if prd == nil {
		return nil, nil
	}

	return prd, nil
}
