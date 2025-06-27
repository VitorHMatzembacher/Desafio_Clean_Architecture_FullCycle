package repository

import (
	"database/sql"
	"project/internal/domain"
)

type OrderPostgresRepository struct {
	DB *sql.DB
}

func NewOrderPostgresRepository(db *sql.DB) *OrderPostgresRepository {
	return &OrderPostgresRepository{DB: db}
}

func (r *OrderPostgresRepository) FindAll() ([]domain.Order, error) {
	rows, err := r.DB.Query("SELECT id, total FROM orders")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []domain.Order
	for rows.Next() {
		var o domain.Order
		err := rows.Scan(&o.ID, &o.Total)
		if err != nil {
			return nil, err
		}
		orders = append(orders, o)
	}
	return orders, nil
}
