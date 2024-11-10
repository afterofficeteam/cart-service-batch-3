package cart

import (
	"cart-service/proto/cart"
	"database/sql"
)

type store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *store {
	return &store{db}
}

type Repository interface {
	Insert(req *cart.CartInsertRequest) (*string, error)
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
