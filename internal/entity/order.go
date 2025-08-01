package entity

import (
	"context"
	"github.com/google/uuid"
)

type Order struct {
	ID         uuid.UUID
	Price      float64
	Tax        float64
	FinalPrice float64
}

type OrderRepository interface {
	Create(ctx context.Context, o Order) error
	FindAll(ctx context.Context) ([]*Order, error)
}
