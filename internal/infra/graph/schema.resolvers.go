package graph

import (
	"context"
	"ordersystem/internal/infra/graph/model"
)

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) ListOrders(ctx context.Context) ([]*model.Order, error) {
	orders, err := r.ListOrdersUseCase.Execute(ctx)
	if err != nil {
		return nil, err
	}
	var result []*model.Order
	for _, o := range orders {
		result = append(result, &model.Order{
			ID:         o.ID.String(),
			Price:      o.Price,
			Tax:        o.Tax,
			FinalPrice: o.FinalPrice,
		})
	}
	return result, nil
}
