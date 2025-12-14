package repo

import (
	"database/sql"
	"ecommerce/domain"

	"github.com/jmoiron/sqlx"
)

type ItemRepo interface {
	Create(item domain.Item) (*domain.Item, error)
	Get(itemID int64, userID int64) (*domain.Item, error)
	List(userID int64, page, limit int64) ([]*domain.Item, error)
	Count(userID int64) (int64, error)
	Delete(itemID int64, userID int64) error
	Update(item domain.Item) (*domain.Item, error)
}

type itemRepo struct {
	db *sqlx.DB
}

// Constructor function
func NewItemRepo(db *sqlx.DB) ItemRepo {
	return &itemRepo{
		db: db,
	}
}

func (r *itemRepo) Create(item domain.Item) (*domain.Item, error) {
	query := `
		INSERT INTO items (
			user_id,
			name,
			description,
			quantity,
			low_stock_threshold,
			price,
			cost_price,
			sku,
			barcode,
			category_id,
			location_id,
			image_urls,
			notes
		) VALUES (
			:user_id,
			:name,
			:description,
			:quantity,
			:low_stock_threshold,
			:price,
			:cost_price,
			:sku,
			:barcode,
			:category_id,
			:location_id,
			:image_urls,
			:notes
		)
		RETURNING id, created_at, updated_at
	`

	rows, err := r.db.NamedQuery(query, item)
	if err != nil {
		return nil, err
	}

	if rows.Next() {
		err = rows.Scan(&item.ID, &item.CreatedAt, &item.UpdatedAt)
		if err != nil {
			return nil, err
		}
	}

	return &item, nil
}

func (r *itemRepo) Get(itemID int64, userID int64) (*domain.Item, error) {
	var item domain.Item

	query := `
		SELECT id, user_id, name, description, quantity, low_stock_threshold,
		       price, cost_price, sku, barcode, category_id, location_id,
		       image_urls, notes, created_at, updated_at
		FROM items
		WHERE id = $1 AND user_id = $2
	`

	err := r.db.Get(&item, query, itemID, userID)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &item, nil
}

func (r *itemRepo) List(userID int64, page, limit int64) ([]*domain.Item, error) {
	var itemList []*domain.Item

	offset := (page - 1) * limit

	query := `
		SELECT id, user_id, name, description, quantity, low_stock_threshold,
		       price, cost_price, sku, barcode, category_id, location_id,
		       image_urls, notes, created_at, updated_at
		FROM items
		WHERE user_id = $1
		ORDER BY id
		LIMIT $2
		OFFSET $3
	`

	err := r.db.Select(&itemList, query, userID, limit, offset)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return itemList, nil
}

func (r *itemRepo) Count(userID int64) (int64, error) {
	var itemCount int64

	query := `
		SELECT COUNT(*) FROM items WHERE user_id = $1
	`

	err := r.db.Get(&itemCount, query, userID)

	if err != nil {
		return 0, err
	}

	return itemCount, nil
}

func (r *itemRepo) Delete(itemID int64, userID int64) error {
	query := `
		DELETE FROM items WHERE id = $1 AND user_id = $2
	`
	_, err := r.db.Exec(query, itemID, userID)
	if err != nil {
		return err
	}

	return nil
}

func (r *itemRepo) Update(item domain.Item) (*domain.Item, error) {
	query := `
		UPDATE items
		SET name = :name,
		    description = :description,
		    quantity = :quantity,
		    low_stock_threshold = :low_stock_threshold,
		    price = :price,
		    cost_price = :cost_price,
		    sku = :sku,
		    barcode = :barcode,
		    category_id = :category_id,
		    location_id = :location_id,
		    image_urls = :image_urls,
		    notes = :notes,
		    updated_at = NOW()
		WHERE id = :id AND user_id = :user_id
		RETURNING updated_at
	`

	rows, err := r.db.NamedQuery(query, item)
	if err != nil {
		return nil, err
	}

	if rows.Next() {
		err = rows.Scan(&item.UpdatedAt)
		if err != nil {
			return nil, err
		}
	}

	return &item, nil
}
