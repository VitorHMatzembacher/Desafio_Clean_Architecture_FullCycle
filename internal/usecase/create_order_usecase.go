package usecase

import (
	"context"

	"github.com/VitorHMatzembacher/Desafio_Clean_Architecture_FullCycle/internal/entity"
	"github.com/google/uuid"
)

// CreateOrderUseCase encapsula a lógica de criação de pedidos.
type CreateOrderUseCase struct {
	Repo entity.OrderRepository
}

// NewCreateOrderUseCase constrói o use case com o repositório fornecido.
func NewCreateOrderUseCase(r entity.OrderRepository) *CreateOrderUseCase {
	return &CreateOrderUseCase{Repo: r}
}

// Execute gera um novo Order e o persiste via repositório.
func (uc *CreateOrderUseCase) Execute(ctx context.Context, price, tax float64) (entity.Order, error) {
	id := uuid.New()
	order := entity.Order{
		ID:         id,
		Price:      price,
		Tax:        tax,
		FinalPrice: price + tax,
	}

	if err := uc.Repo.Create(ctx, order); err != nil {
		return entity.Order{}, err
	}
	return order, nil
}
