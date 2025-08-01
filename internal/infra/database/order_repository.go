package database

import (
	"context"
	"database/sql"

	"github.com/VitorHMatzembacher/Desafio_Clean_Architecture_FullCycle/internal/entity"
)

type OrderRepository struct {
	DB *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{DB: db}
}

func (r *OrderRepository) FindAll(ctx context.Context) ([]*entity.Order, error) {
	rows, err := r.DB.QueryContext(ctx,
		"SELECT id, price, tax, final_price FROM orders",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	orders := make([]*entity.Order, 0, 10)
	for rows.Next() {
		var o entity.Order
		if err := rows.Scan(&o.ID, &o.Price, &o.Tax, &o.FinalPrice); err != nil {
			return nil, err
		}
		orders = append(orders, &o)
	}
	return orders, nil
}

func (r *OrderRepository) Create(ctx context.Context, o entity.Order) error {
	_, err := r.DB.ExecContext(ctx,
		`INSERT INTO orders (id, price, tax, final_price) VALUES ($1, $2, $3, $4)`,
		o.ID, o.Price, o.Tax, o.FinalPrice,
	)
	return err
}
