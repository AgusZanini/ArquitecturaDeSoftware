package services

import (
	client "github.com/AgusZanini/ArquitecturaDeSoftware/proyecto2.0/clients"
	"github.com/AgusZanini/ArquitecturaDeSoftware/proyecto2.0/dto"
	"github.com/AgusZanini/ArquitecturaDeSoftware/proyecto2.0/model"
	errors "github.com/AgusZanini/ArquitecturaDeSoftware/proyecto2.0/utils/errors"
)

type orderdetailservice struct{}

type orderdetailserviceInterface interface {
	GetOrderDetailsByOrder(id int) (dto.OrderDetailsDto, errors.ApiError)
}

var OrderdetailService orderdetailserviceInterface

func init() {
	OrderdetailService = &orderdetailservice{}
}

func (s *orderdetailservice) GetOrderDetailsByOrder(id int) (dto.OrderDetailsDto, errors.ApiError) {
	var orderdetails model.OrderDetails = client.GetOrderDetailsByOrder(id)
	var orderdetailsdto dto.OrderDetailsDto

	if orderdetails == nil {
		return orderdetailsdto, errors.NewBadRequestApiError("nil orderdetails")
	}

	for _, orderdetail := range orderdetails {
		var orderdetaildto dto.OrderDetailDto

		orderdetaildto.Quantity = orderdetail.Quantity
		orderdetaildto.Totalprice = orderdetail.Totalprice
		orderdetaildto.Unitaryprice = orderdetail.Unitaryprice

		orderdetailsdto = append(orderdetailsdto, orderdetaildto)
	}

	return orderdetailsdto, nil
}
