package services

import (
	orderdetailclient "proyectoarqsoft/clients/orderdetails"
	clients "proyectoarqsoft/clients/orders"
	productclients "proyectoarqsoft/clients/products"
	"proyectoarqsoft/dtos"
	"proyectoarqsoft/models"
	"proyectoarqsoft/utils/errors"
	"time"
)

type orderService struct {
	orderclient clients.OrderClientInterface
}

type orderServiceInterface interface {
	InsertOrder(orderrequest dtos.Orderrequestdto) errors.ApiError
	GetOrdersByUser(id int) (dtos.Ordersandetailsdto, errors.ApiError)
}

var (
	OrderService orderServiceInterface
)

func initOrderService(orderclient clients.OrderClientInterface) orderServiceInterface {
	service := new(orderService)
	service.orderclient = orderclient
	return service
}

func init() {
	OrderService = initOrderService(clients.OrderClient)
}

func (o *orderService) InsertOrder(orderrequest dtos.Orderrequestdto) errors.ApiError {
	var order models.Order
	var details models.Orderdetails
	detailsdto := orderrequest.Orderdetails
	var total float32 = 0.0

	order.Userid = orderrequest.Userid
	order.Date = time.Now().Format("2006.01.02 15:04:05")
	order.Address = orderrequest.Address

	for _, detaildto := range detailsdto {
		var detail models.Orderdetail

		product, err := productclients.ProductClient.GetproductById(detaildto.Productid)
		if err != nil {
			return errors.NewBadRequestApiError(err.Error())
		}

		if detaildto.Amount > product.Stock {
			return errors.NewBadRequestApiError("stock insuficiente en producto: " + product.Name)
		}

		detail.Amount = detaildto.Amount
		detail.Productid = detaildto.Productid
		detail.Total = product.Price * float32(detail.Amount)

		total += detail.Total
		details = append(details, detail)
	}

	order.Total = total
	er := o.orderclient.InsertOrder(order)
	if er != nil {
		return errors.NewConflictApiError("no se pudo insertar la orden")
	}

	neworder, er := o.orderclient.GetLastOrder()
	if er != nil {
		return errors.NewConflictApiError("error en nueva orden")
	}

	for _, detail := range details {
		detail.Orderid = neworder.Orderid

		err := orderdetailclient.OrderdetailClient.InserOrderDetail(detail)
		if err != nil {
			return errors.NewBadRequestApiError("no se pudo insertar detalles de la orden")
		}
		_, er := productclients.ProductClient.UpdateStock(detail.Productid, detail.Amount)
		if er != nil {
			return errors.NewBadRequestApiError("no se pudo actualizar el stock del producto")
		}
	}

	return nil
}

func (o *orderService) GetOrdersByUser(id int) (dtos.Ordersandetailsdto, errors.ApiError) {
	var ordersandetailsdto dtos.Ordersandetailsdto
	orders, err := o.orderclient.Getordersbyuser(id)
	if err != nil {
		return dtos.Ordersandetailsdto{}, errors.NewBadRequestApiError(err.Error())
	}

	for _, order := range orders {
		var orderdto dtos.Orderdto
		var orderandetailsdto dtos.Orderandetailsdto

		orderdto.Address = order.Address
		orderdto.Date = order.Date
		orderdto.Orderid = order.Orderid
		orderdto.Total = order.Total
		orderdto.Userid = order.Userid

		orderandetailsdto.Orderdto = orderdto

		orderdetails, err := orderdetailclient.OrderdetailClient.GetOrderDetailsByOrder(order.Orderid)
		if err != nil {
			return dtos.Ordersandetailsdto{}, errors.NewBadRequestApiError("no se pudo obtener los detalles de la orden: " + string(rune(order.Orderid)))
		}
		for _, orderdetail := range orderdetails {
			var orderdetaildto dtos.Orderdetaildto

			orderdetaildto.Amount = orderdetail.Amount
			orderdetaildto.Orderdetailid = orderdetail.Orderdetailid
			orderdetaildto.Orderid = orderdetail.Orderid
			orderdetaildto.Productid = orderdetail.Productid
			orderdetaildto.Total = orderdetail.Total

			orderandetailsdto.Orderdetailsdto = append(orderandetailsdto.Orderdetailsdto, orderdetaildto)
		}
		ordersandetailsdto = append(ordersandetailsdto, orderandetailsdto)
	}
	return ordersandetailsdto, nil
}
