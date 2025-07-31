package graph

import "github.com/VitorHMatzembacher/Desafio_Clean_Architecture_FullCycle/internal/usecase"

type Resolver struct {
	ListOrdersUseCase *usecase.ListOrdersUseCase
}
