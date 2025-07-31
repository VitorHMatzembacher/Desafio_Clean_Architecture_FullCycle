package graph

import (
	"context"

	"github.com/VitorHMatzembacher/Desafio_Clean_Architecture_FullCycle/internal/infra/graph/model"
)

type QueryResolver interface {
	ListOrders(ctx context.Context) ([]*model.Order, error)
}
