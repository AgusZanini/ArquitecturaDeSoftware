package services

import (
	orderClient "github.com/AgusZanini/ArquitecturaDeSoftware/proyecto/clients/order"

	"github.com/AgusZanini/ArquitecturaDeSoftware/proyecto/dto"
	"github.com/AgusZanini/ArquitecturaDeSoftware/proyecto/model"
	errors "github.com/AgusZanini/ArquitecturaDeSoftware/proyecto/utils/errors"
)

type orderservice struct {
}

type orderserviceInterface interface {
	GetOrderById(id int) (dto.OrderDto, errors.ApiError)
	GetOrders() (dto.OrdersDto, errors.ApiError)
	InsertOrder(orderDto dto.OrderDto) (dto.OrderDto, errors.ApiError)
}

var (
	OrderService orderserviceInterface
)

func init() {
	OrderService = &orderservice{}
}

func (s *orderservice) GetOrderById(id int) (dto.OrderDto, errors.ApiError) {
	var order model.Order = orderClient.GetOrderById(id)
	var orderDto dto.OrderDto

	if order.ID_order == 0 {
		return orderDto, errors.NewBadRequestApiError("orden no encontrada")
	}

	orderDto.ID_order = order.ID_order
	//orderDto.Client.ID_client = order.Client.ID_client
	orderDto.ID_client = order.ID_client
	orderDto.Discount = order.Discount
	orderDto.Total_price = order.Total_price
	orderDto.State = order.State

	return orderDto, nil
}

func (s *orderservice) GetOrders() (dto.OrdersDto, errors.ApiError) {
	var orders model.Orders = orderClient.GetOrders()
	var ordersDto dto.OrdersDto

	for _, order := range orders {
		var orderDto dto.OrderDto
		orderDto.ID_order = order.ID_order
		//orderDto.Client.ID_client = order.Client.ID_client
		orderDto.ID_client = order.ID_client
		orderDto.Discount = order.Discount
		orderDto.Total_price = order.Total_price
		orderDto.State = order.State

		ordersDto = append(ordersDto, orderDto)
	}

	return ordersDto, nil
}

func (s *orderservice) InsertOrder(orderDto dto.OrderDto) (dto.OrderDto, errors.ApiError) {
	var order model.Order

	//order.Client.ID_client = orderDto.Client.ID_client
	order.ID_client = orderDto.ID_client
	order.Discount = orderDto.Discount
	order.Total_price = orderDto.Total_price
	order.State = orderDto.State

	order = orderClient.InsertOrder(order)
	order.ID_order = orderDto.ID_order

	return orderDto, nil
}
