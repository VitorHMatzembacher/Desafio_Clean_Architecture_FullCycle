package graph

import (
	"context"
	"ordersystem/internal/infra/graph/model"
)

type QueryResolver interface {
	ListOrders(ctx context.Context) ([]*model.Order, error)
}
