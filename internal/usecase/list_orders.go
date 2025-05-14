package usecase

import (
	"context"

	"github.com/leonardogomesdossantos/clean-architecture-challenge/internal/entity"
)

type ListOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewListOrdersUseCase(orderRepository entity.OrderRepositoryInterface) *ListOrdersUseCase {
	return &ListOrdersUseCase{
		OrderRepository: orderRepository,
	}
}

func (uc *ListOrdersUseCase) Execute(ctx context.Context) ([]entity.Order, error) {
	return uc.OrderRepository.List(ctx)
}
