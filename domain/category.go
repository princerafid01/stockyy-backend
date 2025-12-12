package domain

import "time"

type Category struct {
	ID        int64     `json:"id" db:"id"`
	UserID    int64     `json:"user_id" db:"user_id"`
	Name      string    `json:"name" db:"name"`
	Color     string    `json:"color" db:"color"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
