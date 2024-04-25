package service

import (
	"context"
	"github.com/rodrigoachilles/clean-architecture/internal/infra/grpc/pb"
	"github.com/rodrigoachilles/clean-architecture/internal/usecase"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	CreateOrderUseCase usecase.CreateOrderUseCase
	ListOrdersUseCase  usecase.ListOrdersUseCase
}

func NewOrderService(createOrderUseCase usecase.CreateOrderUseCase, listOrdersUseCase usecase.ListOrdersUseCase) *OrderService {
	return &OrderService{
		CreateOrderUseCase: createOrderUseCase,
		ListOrdersUseCase:  listOrdersUseCase,
	}
}

func (s *OrderService) CreateOrder(_ context.Context, in *pb.CreateOrderRequest) (*pb.Order, error) {
	dto := usecase.OrderInputDTO{
		ProductName: in.ProductName,
		Price:       float64(in.Price),
		Tax:         float64(in.Tax),
	}
	output, err := s.CreateOrderUseCase.Execute(dto)
	if err != nil {
		return nil, err
	}
	return &pb.Order{
		Id:          output.ID.String(),
		ProductName: output.ProductName,
		Price:       float32(output.Price),
		Tax:         float32(output.Tax),
		FinalPrice:  float32(output.FinalPrice),
	}, nil
}

func (s *OrderService) ListOrders(_ context.Context, in *pb.ListOrdersRequest) (*pb.ListOrdersResponse, error) {
	dto := usecase.ListOrderInputDTO{
		Page:  int(in.Page),
		Limit: int(in.Limit),
		Sort:  in.Sort,
	}
	output, err := s.ListOrdersUseCase.Execute(dto)
	if err != nil {
		return nil, err
	}

	var listOrdersResponse []*pb.Order
	for _, order := range *output.Orders {
		orderResponse := &pb.Order{
			Id:          order.ID.String(),
			ProductName: order.ProductName,
			Price:       float32(order.Price),
			Tax:         float32(order.Tax),
			FinalPrice:  float32(order.FinalPrice),
		}

		listOrdersResponse = append(listOrdersResponse, orderResponse)
	}

	return &pb.ListOrdersResponse{Orders: listOrdersResponse}, nil
}
