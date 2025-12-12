package domain

import "time"

type StockMovement struct {
	ID             int64     `json:"id" db:"id"`
	ItemID         int64     `json:"item_id" db:"item_id"`
	UserID         int64     `json:"user_id" db:"user_id"`
	QuantityChange int       `json:"quantity_change" db:"quantity_change"` // positive = in, negative = out
	NewQuantity    int       `json:"new_quantity" db:"new_quantity"`       // quantity after this movement
	Reason         *string   `json:"reason,omitempty" db:"reason"`         // "Sold", "Received", "Damage", "Manual adjust"…
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
}
