package user

import "ecommerce/domain"

type Service interface {
	Find(email string, passwordHash *string) (*domain.User, error)
	Create(domain.User) (*domain.User, error)
	FindByEmail(email string) (*domain.User, error)
}
