package service

import (
	"context"

	pb "github.com/VitorHMatzembacher/Desafio_Clean_Architecture_FullCycle/internal/infra/grpc/pb"
	"github.com/VitorHMatzembacher/Desafio_Clean_Architecture_FullCycle/internal/usecase"
)

type OrderServiceServer struct {
	pb.UnimplementedOrderServiceServer
	ListOrdersUseCase *usecase.ListOrdersUseCase
}

func NewOrderServiceServer(uc *usecase.ListOrdersUseCase) *OrderServiceServer {
	return &OrderServiceServer{ListOrdersUseCase: uc}
}

func (s *OrderServiceServer) ListOrders(ctx context.Context, _ *pb.Empty) (*pb.OrderList, error) {
	orders, err := s.ListOrdersUseCase.Execute(ctx)
	if err != nil {
		return nil, err
	}

	resp := &pb.OrderList{}
	for _, o := range orders {
		resp.Orders = append(resp.Orders, &pb.Order{
			Id:         o.ID.String(),
			Price:      o.Price,
			Tax:        o.Tax,
			FinalPrice: o.FinalPrice,
		})
	}
	return resp, nil
}
