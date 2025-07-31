package entity

import "github.com/google/uuid"

type Order struct {
	ID         uuid.UUID
	Price      float64
	Tax        float64
	FinalPrice float64
}

func NewOrder(price, tax float64) *Order {
	finalPrice := price + tax
	return &Order{
		ID:         uuid.New(),
		Price:      price,
		Tax:        tax,
		FinalPrice: finalPrice,
	}
}
