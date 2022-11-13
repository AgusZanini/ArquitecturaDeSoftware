package services

import (
	orderdetailClient "github.com/AgusZanini/ArquitecturaDeSoftware/proyecto/clients/orderdetail"

	"github.com/AgusZanini/ArquitecturaDeSoftware/proyecto/dto"
	"github.com/AgusZanini/ArquitecturaDeSoftware/proyecto/model"
	errors "github.com/AgusZanini/ArquitecturaDeSoftware/proyecto/utils/errors"
)

type orderdetailservice struct{}

type orderdetailserviceInterface interface {
	GetOrderDetailById(id int) (dto.OrderdetailDto, errors.ApiError)
	GetOrderDetails() (dto.OrderdetailsDto, errors.ApiError)
	InsertOrderDetail(orderdetailDto dto.OrderdetailDto) (dto.OrderdetailDto, errors.ApiError)
}

var (
	Orderdetailservice orderdetailserviceInterface
)

func init() {
	Orderdetailservice = &orderdetailservice{}
}

func (s *orderdetailservice) GetOrderDetailById(id int) (dto.OrderdetailDto, errors.ApiError) {
	var orderdetail model.OrderDetail = orderdetailClient.GetOrderDetailById(id)
	var orderdetailDto dto.OrderdetailDto

	if orderdetail.ID_orderdetail == 0 {
		return orderdetailDto, errors.NewBadRequestApiError("detalle de orden no encontrada")
	}

	orderdetailDto.ID_orderdetail = orderdetail.ID_orderdetail
	//orderdetailDto.Order.ID_order = orderdetail.Order.ID_order
	orderdetailDto.ID_order = orderdetail.ID_order
	//orderdetailDto.Product.ID_product = orderdetail.Product.ID_product
	orderdetailDto.ID_product = orderdetail.ID_product
	orderdetailDto.Quantity = orderdetail.Quantity

	return orderdetailDto, nil
}

func (s *orderdetailservice) GetOrderDetails() (dto.OrderdetailsDto, errors.ApiError) {
	var orderdetails model.OrderDetails = orderdetailClient.GetOrderDetails()
	var orderdetailsDto dto.OrderdetailsDto

	for _, orderdetail := range orderdetails {
		var orderdetailDto dto.OrderdetailDto
		orderdetailDto.ID_orderdetail = orderdetail.ID_orderdetail
		//orderdetailDto.Order.ID_order = orderdetail.Order.ID_order
		//orderdetailDto.Product.ID_product = orderdetail.Product.ID_product
		orderdetailDto.ID_order = orderdetail.ID_order
		orderdetailDto.ID_product = orderdetail.ID_product
		orderdetailDto.Quantity = orderdetail.Quantity

		orderdetailsDto = append(orderdetailsDto, orderdetailDto)
	}

	return orderdetailsDto, nil
}

func (s *orderdetailservice) InsertOrderDetail(orderdetailDto dto.OrderdetailDto) (dto.OrderdetailDto, errors.ApiError) {
	var orderdetail model.OrderDetail

	//orderdetail.Order.ID_order = orderdetailDto.Order.ID_order
	//orderdetail.Product.ID_product = orderdetailDto.Product.ID_product
	orderdetail.ID_order = orderdetailDto.ID_order
	orderdetail.ID_product = orderdetailDto.ID_product
	orderdetail.Quantity = orderdetailDto.Quantity

	orderdetail = orderdetailClient.InsertOrderDetail(orderdetail)

	orderdetail.ID_orderdetail = orderdetailDto.ID_orderdetail

	return orderdetailDto, nil

}
