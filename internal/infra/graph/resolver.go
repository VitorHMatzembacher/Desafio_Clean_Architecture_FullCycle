package graph

import (
	"context"
	"github.com/graph-gophers/graphql-go"

	"github.com/VitorHMatzembacher/Desafio_Clean_Architecture_FullCycle/internal/infra/graph/model"
	"github.com/VitorHMatzembacher/Desafio_Clean_Architecture_FullCycle/internal/usecase"
)

type Resolver struct {
	ListOrdersUseCase *usecase.ListOrdersUseCase
}

type orderResolver struct {
	o *model.Order
}

func (r *Resolver) ListOrders(ctx context.Context) ([]*orderResolver, error) {
	orders, err := r.ListOrdersUseCase.Execute(ctx)
	if err != nil {
		return nil, err
	}

	resp := make([]*orderResolver, len(orders))
	for i, o := range orders {
		m := &model.Order{
			ID:         graphql.ID(o.ID.String()),
			Price:      o.Price,
			Tax:        o.Tax,
			FinalPrice: o.FinalPrice,
		}
		resp[i] = &orderResolver{o: m}
	}
	return resp, nil
}

func (r *orderResolver) Id() graphql.ID      { return r.o.ID }
func (r *orderResolver) Price() float64      { return r.o.Price }
func (r *orderResolver) Tax() float64        { return r.o.Tax }
func (r *orderResolver) FinalPrice() float64 { return r.o.FinalPrice }
