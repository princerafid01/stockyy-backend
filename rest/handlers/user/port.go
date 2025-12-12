package user

import "ecommerce/domain"

type Service interface {
	Find(email string, password string) (*domain.User, error)
	Create(domain.User) (*domain.User, error)
}
