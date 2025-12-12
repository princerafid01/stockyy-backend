package repo

import (
	"database/sql"
	"ecommerce/domain"

	"github.com/jmoiron/sqlx"
)

type ProductRepo interface {
	Create(p domain.Product) (*domain.Product, error)
	Get(productID int) (*domain.Product, error)
	List(page, limit int64) ([]*domain.Product, error)
	Count() (int64, error)
	Delete(productID int) error
	Update(product domain.Product) (*domain.Product, error)
}

type productRepo struct {
	// productList []*domain.Product
	db *sqlx.DB
}

// Constructor function
func NewProductRepo(db *sqlx.DB) ProductRepo {
	return &productRepo{
		db: db,
	}
}

func (r *productRepo) Create(p domain.Product) (*domain.Product, error) {
	query := `
		INSERT INTO products (
			title,
			description,
			price,
			img_url
		) VALUES (
			$1,
			$2,
			$3,
			$4
		)
		RETURNING id
	`

	row := r.db.QueryRow(query, p.Title, p.Description, p.Price, p.ImgUrl)

	err := row.Scan(&p.ID)

	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *productRepo) Get(productID int) (*domain.Product, error) {
	var prd domain.Product

	query := `SELECT id, title, description, price, img_url from products WHERE id = $1`

	err := r.db.Get(&prd, query, productID)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &prd, nil
}

func (r *productRepo) List(page, limit int64) ([]*domain.Product, error) {
	var prdList []*domain.Product

	offset := (page - 1) * limit

	query := `
		SELECT id, title, description, price, img_url  from products
		ORDER BY id
		LIMIT $1
		OFFSET $2
	`

	err := r.db.Select(&prdList, query, limit, offset)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return prdList, nil
}

func (r *productRepo) Count() (int64, error) {
	var productCount int64

	query := `
		SELECT COUNT(*) from products
	`

	err := r.db.Get(&productCount, query)

	if err != nil {
		return 0, err
	}

	return productCount, nil
}

func (r *productRepo) Delete(productID int) error {
	query := `
		DELETE from products where id= $1
	`
	_, err := r.db.Exec(query, productID)
	if err != nil {
		return err
	}

	return nil
}

func (r *productRepo) Update(p domain.Product) (*domain.Product, error) {
	query := `
		UPDATE products
		SET title=$1, description=$2, price=$3 , img_url = $4
	 	WHERE id=$5
	`

	row := r.db.QueryRow(query, p.Title, p.Description, p.Price, p.ImgUrl, p.ID)
	err := row.Err()

	if err != nil {
		return nil, err
	}

	return &p, nil
}
