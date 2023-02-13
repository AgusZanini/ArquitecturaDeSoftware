package services

import (
	oclient "mvc/clients/order"
	odclient "mvc/clients/order_detail"
	pclient "mvc/clients/product"
	"mvc/dto"
	"mvc/model"
	e "mvc/utils/errors"
	"time"

	log "github.com/sirupsen/logrus"
)

type orderService struct {
	orderClient       oclient.OrderClientInterface
	productClient     pclient.ProductClientInterface
	orderDetailClient odclient.OrderDetailClientInterface
}

type orderServiceInterface interface {
	InsertOrder(orderDto dto.OrderInsertDto) (dto.OrderResponseDto, e.ApiError)
	GetOrdersByUserId(id int) (dto.OrdersDto, e.ApiError)
}

var (
	OrderService orderServiceInterface
)

func initOrderService(orderClient oclient.OrderClientInterface, orderDetailClient odclient.OrderDetailClientInterface, productClient pclient.ProductClientInterface) orderServiceInterface {
	service := new(orderService)
	service.orderClient = orderClient
	service.productClient = productClient
	service.orderDetailClient = orderDetailClient
	return service
}

func init() {
	OrderService = initOrderService(
		oclient.OrderClient,
		odclient.OrderDetailClient,
		pclient.ProductClient)
}

func (s *orderService) GetOrdersByUserId(id int) (dto.OrdersDto, e.ApiError) {

	var orders model.Orders = s.orderClient.GetOrdersByUserId(id)
	var ordersDto dto.OrdersDto

	for _, order := range orders {
		var orderDto dto.OrderDto
		var details = s.orderDetailClient.GetOrderDetailsByOrderId(order.ID)
		orderDto.OrderId = order.ID
		orderDto.Date = order.Date
		orderDto.Total = order.Total
		orderDto.CurrencyId = order.CurrencyId
		for _, orderDetail := range details {
			var d dto.OrderDetailDto
			d.OrderDetailId = orderDetail.OrderDetailId
			d.ProductId = orderDetail.ProductId
			d.Quantity = orderDetail.Quantity
			d.Price = orderDetail.Price
			d.CurrencyId = orderDetail.CurrencyId
			d.Name = orderDetail.Name
			orderDto.OrderDetails = append(orderDto.OrderDetails, d)
		}

		ordersDto = append(ordersDto, orderDto)
	}
	return ordersDto, nil
}

func (s *orderService) InsertOrder(orderInsertDto dto.OrderInsertDto) (dto.OrderResponseDto, e.ApiError) {

	var order model.Order
	var total float32
	var orderResponseDto dto.OrderResponseDto

	total = 0
	order.UserId = orderInsertDto.UserId
	order.Date = time.Now().Format("2006.01.02 15:04:05")
	for i := 0; i < len(orderInsertDto.OrderDetails); i++ {
		detail := orderInsertDto.OrderDetails[i]
		product := s.productClient.GetProductById(detail.ProductId)
		if product.Stock < detail.Quantity {
			orderResponseDto.OrderId = 0
			return orderResponseDto, e.NewConflictApiError("Not enough stock on product: " + product.Name)
		}

		total += (detail.Price * float32(detail.Quantity))

	}

	for i := 0; i < len(orderInsertDto.OrderDetails); i++ {
		detail := orderInsertDto.OrderDetails[i]
		s.productClient.RemoveStock(detail.ProductId, detail.Quantity)
	}

	order.Total = total
	order.CurrencyId = "ARS"

	order = s.orderClient.InsertOrder(order)

	orderResponseDto.OrderId = order.ID

	log.Debug(order)
	return orderResponseDto, nil
}
