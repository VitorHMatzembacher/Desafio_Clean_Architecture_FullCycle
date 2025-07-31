package graph

import "ordersystem/internal/usecase"

type Resolver struct {
	ListOrdersUseCase *usecase.ListOrdersUseCase
}
