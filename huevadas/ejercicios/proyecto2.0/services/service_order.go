package services

import (
	client "github.com/AgusZanini/ArquitecturaDeSoftware/proyecto2.0/clients"
	"github.com/AgusZanini/ArquitecturaDeSoftware/proyecto2.0/dto"
	"github.com/AgusZanini/ArquitecturaDeSoftware/proyecto2.0/model"
	errors "github.com/AgusZanini/ArquitecturaDeSoftware/proyecto2.0/utils/errors"
)

type orderservice struct{}

type orderserviceInterface interface {
	GetOrdersByUser(id int) (dto.OrdersDto, errors.ApiError)
	GetOrders() (dto.OrdersDto, errors.ApiError)
	InsertOrder(orderdto dto.OrderDto) (dto.OrderDto, errors.ApiError)
}

var (
	OrderService orderserviceInterface
)

func init() {
	OrderService = &orderservice{}
}

func (s *orderservice) GetOrdersByUser(id int) (dto.OrdersDto, errors.ApiError) {
	var orders model.Orders = client.GetOrdersByUser(id)
	var ordersdto dto.OrdersDto

	if orders == nil {
		return ordersdto, errors.NewBadRequestApiError("nil orders")
	}

	for _, order := range orders {
		var orderdto dto.OrderDto

		orderdto.IDorder = order.IDorder
		orderdto.State = order.State
		orderdto.Totalprice = order.Totalprice
		orderdto.IDuser = order.IDuser

		ordersdto = append(ordersdto, orderdto)
	}

	return ordersdto, nil
}

func (s *orderservice) GetOrders() (dto.OrdersDto, errors.ApiError) {
	var orders model.Orders = client.GetOrders()
	var ordersdto dto.OrdersDto

	if orders == nil {
		return ordersdto, errors.NewBadRequestApiError("nil orders")
	}

	for _, order := range orders {
		var orderdto dto.OrderDto

		orderdto.IDorder = order.IDorder
		orderdto.State = order.State
		orderdto.Totalprice = order.Totalprice
		orderdto.IDuser = order.IDuser

		ordersdto = append(ordersdto, orderdto)
	}

	return ordersdto, nil
}

func (s *orderservice) InsertOrder(orderdto dto.OrderDto) (dto.OrderDto, errors.ApiError) {
	var order model.Order

	order.IDuser = orderdto.IDuser

	order = client.InsertOrder(order)

	var totalprice float32
	var orderdetails model.OrderDetails

	for _, orderdetaildto := range orderdto.OrderDetailsDto {
		var orderdetail model.OrderDetail

		orderdetail.IDproduct = orderdetaildto.IDproduct

		var product model.Product = client.GetProductById(orderdetail.IDproduct)

		orderdetail.Quantity = orderdetaildto.Quantity
		orderdetail.Unitaryprice = product.Baseprice
		orderdetail.Totalprice = orderdetail.Unitaryprice * float32(orderdetail.Quantity)
		orderdetail.IDorder = orderdetaildto.IDorder

		totalprice += orderdetail.Totalprice

		client.UpdateStock(orderdetail.IDproduct, int(orderdetaildto.Quantity))

		orderdetails = append(orderdetails, orderdetail)
	}

	client.UpdateTotalPrice(order.IDorder, totalprice)

	client.InsertOrderDetails(orderdetails)

	return orderdto, nil
}
