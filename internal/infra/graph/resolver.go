package graph

import (
	"context"

	"github.com/VitorHMatzembacher/Desafio_Clean_Architecture_FullCycle/internal/infra/graph/model"
	"github.com/VitorHMatzembacher/Desafio_Clean_Architecture_FullCycle/internal/usecase"
)

type Resolver struct {
	ListOrdersUseCase *usecase.ListOrdersUseCase
}

func (r *Resolver) ListOrders(ctx context.Context) ([]*model.Order, error) {
	orders, err := r.ListOrdersUseCase.Execute(ctx)
	if err != nil {
		return nil, err
	}

	var resp []*model.Order
	for _, o := range orders {
		resp = append(resp, &model.Order{
			ID:         o.ID.String(),
			Price:      o.Price,
			Tax:        o.Tax,
			FinalPrice: o.FinalPrice,
		})
	}

	return resp, nil
}
