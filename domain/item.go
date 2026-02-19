package domain

import (
	"time"

	"github.com/lib/pq"
)

// Item represents an inventory item
type Item struct {
	ID                int64          `json:"id" db:"id"`
	UserID            int64          `json:"user_id" db:"user_id"`
	Name              string         `json:"name" db:"name"`
	Description       *string        `json:"description,omitempty" db:"description"`
	Quantity          int            `json:"quantity" db:"quantity"`
	LowStockThreshold int            `json:"low_stock_threshold" db:"low_stock_threshold"`
	Price             *float64       `json:"price,omitempty" db:"price"`           // DECIMAL(12,2)
	CostPrice         *float64       `json:"cost_price,omitempty" db:"cost_price"` // DECIMAL(12,2)
	SKU               *string        `json:"sku,omitempty" db:"sku"`
	Barcode           *string        `json:"barcode,omitempty" db:"barcode"`
	CategoryID        *int64         `json:"category_id,omitempty" db:"category_id"`
	LocationID        *int64         `json:"location_id,omitempty" db:"location_id"`
	ImageURLs         pq.StringArray `json:"image_urls" db:"image_urls"` // TEXT[]
	Notes             *string        `json:"notes,omitempty" db:"notes"`
	CreatedAt         time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt         time.Time      `json:"updated_at" db:"updated_at"`
}
