package usecase

import "github.com/ryancarlos88/clean-arch/internal/entity"

type ListOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

type ListOrdersOutputDTO struct {
	Orders []OrderOutputDTO `json:"orders"`
}

func NewListOrdersUseCase(OrderRepository entity.OrderRepositoryInterface) *ListOrdersUseCase {
	return &ListOrdersUseCase{
		OrderRepository: OrderRepository,
	}
}

func (c *ListOrdersUseCase) Execute() (ListOrdersOutputDTO, error) {
	orders, err := c.OrderRepository.FindAll()
	if err != nil {
		return ListOrdersOutputDTO{}, err
	}

	var ordersDTO []OrderOutputDTO
	for _, order := range orders {
		dto := OrderOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.Price + order.Tax,
		}
		ordersDTO = append(ordersDTO, dto)
	}

	return ListOrdersOutputDTO{Orders: ordersDTO}, nil
}

