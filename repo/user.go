package repo

import (
	"database/sql"
	"ecommerce/domain"
	"ecommerce/user"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type UserRepo interface {
	user.UserRepo
}

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) UserRepo {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) Create(user domain.User) (*domain.User, error) {
	query := `
		INSERT INTO users (
		    email,
		    password_hash,
		    google_id,
		    name,
		    avatar_url
		) VALUES (
		    :email,
		    :password_hash,
		    :google_id,
		    :name,
		    :avatar_url
		)
		RETURNING id
	`
	var userID int64

	rows, err := r.db.NamedQuery(query, user)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if rows.Next() {
		rows.Scan(&userID)
	}

	user.ID = userID

	return &user, nil

}

func (r *userRepo) Find(email, pass string) (*domain.User, error) {
	var user domain.User

	query := `
		SELECT id, email, password_hash, google_id, name, avatar_url, created_at, updated_at
		FROM users
		WHERE email = $1 AND password_hash = $2
		LIMIT 1;
	`

	err := r.db.Get(&user, query, email, pass)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, err
}
