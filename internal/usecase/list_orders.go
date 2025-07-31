package usecase

import (
	"context"

	"github.com/VitorHMatzembacher/Desafio_Clean_Architecture_FullCycle/internal/entity"
)

type OrderRepository interface {
	FindAll(ctx context.Context) ([]*entity.Order, error)
}

type ListOrdersUseCase struct {
	Repository OrderRepository
}

func NewListOrdersUseCase(repo OrderRepository) *ListOrdersUseCase {
	return &ListOrdersUseCase{
		Repository: repo,
	}
}

func (uc *ListOrdersUseCase) Execute(ctx context.Context) ([]*entity.Order, error) {
	return uc.Repository.FindAll(ctx)
}
