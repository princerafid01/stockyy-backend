package user

import (
	"ecommerce/domain"
	userHandler "ecommerce/rest/handlers/user"
)

// these depends on me
type Service interface {
	userHandler.Service // embedding
}

// I depends on them
type UserRepo interface {
	Create(user domain.User) (*domain.User, error)
	Find(email string, passwordHash *string) (*domain.User, error)
	FindByEmail(email string) (*domain.User, error)
	// List() ([]*User, error)
	// Delete(userID int) error
	// Update(user User) (*User, error)
}
