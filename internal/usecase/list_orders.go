package usecase

import (
	"project/internal/domain"
)

type OrderRepository interface {
	FindAll() ([]domain.Order, error)
}

type ListOrdersUseCase struct {
	Repo OrderRepository
}

func (uc *ListOrdersUseCase) Execute() ([]domain.Order, error) {
	return uc.Repo.FindAll()
}
