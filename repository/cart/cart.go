package cart

import (
	"cart-service/proto/cart"
	"database/sql"
	"fmt"
)

type store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *store {
	return &store{db}
}

type Repository interface {
	Insert(req *cart.CartInsertRequest) (*string, error)
	GetDetails(req *cart.CartDetailRequest) (*cart.CartDetailResponse, error)
}

func (s *store) Insert(req *cart.CartInsertRequest) (*string, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	queryStatement := `
		INSERT INTO cart_items (user_id, product_id, qty) VALUES ($1, $2, $3)
	`

	if _, err := tx.Exec(queryStatement, req.UserId, req.ProductId, req.Qty); err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return nil, err
	}

	insertOK := "Insert cart success"
	return &insertOK, nil
}

func (s *store) GetDetails(req *cart.CartDetailRequest) (*cart.CartDetailResponse, error) {
	queryStatement := `
		SELECT * FROM cart_items WHERE user_id = $1 AND product_id = $2
	`

	var response cart.CartDetailResponse
	row, err := s.db.Query(queryStatement, req.Id, req.ProductId)
	if err != nil {
		return nil, err
	}

	for row.Next() {
		if err := row.Scan(
			&response.Id,
			&response.UserId,
			&response.ProductId,
			&response.Qty,
			&response.CreatedAt,
			&response.UpdatedAt,
			&response.DeletedAt,
		); err != nil {
			if err == sql.ErrNoRows {
				return nil, fmt.Errorf("no cart found")
			}

			return nil, fmt.Errorf("failed to fetch cart data")
		}
	}

	if err := row.Err(); err != nil {
		return nil, fmt.Errorf("failed to iterate over cart: %v", err)
	}

	return &response, nil
}
