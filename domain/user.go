package domain

import "time"

type User struct {
	ID           int64     `json:"id" db:"id"`
	Email        string    `json:"email" db:"email"`
	PasswordHash *string   `json:"password_hash,omitempty" db:"password_hash"` // NULL for Google-only users
	GoogleID     *string   `json:"google_id,omitempty" db:"google_id"`
	Name         string    `json:"name" db:"name"`
	AvatarURL    *string   `json:"avatar_url,omitempty" db:"avatar_url"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}
