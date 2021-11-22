package controllerOrder

import (
	uuid "github.com/google/uuid"

	ControllerProduct "orders/controller/product"
	Database "orders/database"
	BO "orders/model"
	DTO "orders/service/dto"
)

func GetOrderListByCustomerId(customerId uuid.UUID, limit int, offset int) ([]DTO.Order, error) {

	orderBO, err := Database.GetOrderListByCustomerId(customerId, limit, offset)
	orderDTO := make([]DTO.Order, len(orderBO))

	if err != nil {
		return orderDTO, err
	}

	// mapper
	for i, order := range orderBO {

		orderDTO[i].UUID = order.UUID
		orderDTO[i].CustomerId = order.CustomerId
		orderDTO[i].PromotionId = order.PromotionId
		orderDTO[i].OriginalPrice = order.OriginalPrice
		orderDTO[i].DiscountPercentage = order.DiscountPercentage
		orderDTO[i].TotalPrice = order.TotalPrice

		// TODO
		// decide to include products
		orderDTO[i].Products, _ = ControllerProduct.GetProductListByOrderId(order.UUID)
	}

	return orderDTO, nil
}

func GetOrderById(orderId uuid.UUID) (DTO.Order, error) {

	var orderDTO DTO.Order
	orderBO, err := Database.GetOrderById(orderId)

	if err != nil {
		return orderDTO, err
	}

	// mapper
	orderDTO.UUID = orderBO.UUID
	orderDTO.CustomerId = orderBO.CustomerId
	orderDTO.PromotionId = orderBO.PromotionId
	orderDTO.OriginalPrice = orderBO.OriginalPrice
	orderDTO.DiscountPercentage = orderBO.DiscountPercentage
	orderDTO.TotalPrice = orderBO.TotalPrice
	orderDTO.Products, _ = ControllerProduct.GetProductListByOrderId(orderBO.UUID)

	return orderDTO, nil
}

func CreateNewOrder(orderDTO DTO.Order) (DTO.Order, error) {

	orderDTO.UUID = uuid.New()

	var orderBO BO.Order
	orderBO.UUID = orderDTO.UUID
	orderBO.CustomerId = orderDTO.CustomerId
	orderBO.PromotionId = orderDTO.PromotionId
	orderBO.DiscountPercentage = orderDTO.DiscountPercentage

	// orderBO.TotalPrice
	for i, product := range orderDTO.Products {

		// Calculate product totalPrice
		orderDTO.Products[i].Price = product.OriginalPrice
		if product.DiscountPercentage > 0 {

			orderDTO.Products[i].Price = product.OriginalPrice - (product.OriginalPrice*float64(product.DiscountPercentage))/100
		}

		// Calculate order totalPrice
		orderBO.TotalPrice += orderDTO.Products[i].Price

		// Calculate order OriginalPrice
		orderBO.OriginalPrice += product.OriginalPrice
	}

	// Discount order discounted price
	if orderDTO.DiscountPercentage > 0 {

		orderBO.TotalPrice = orderBO.TotalPrice - (orderBO.TotalPrice*float64(orderDTO.DiscountPercentage))/100
	}

	// Create new order
	outcomeOrder, errOrder := Database.CreateNewOrder(orderBO)

	if !outcomeOrder && errOrder != nil {

		return DTO.Order{}, errOrder
	}

	// create new products
	for _, product := range orderDTO.Products {

		outcomeProduct, errProduct := ControllerProduct.CreateNewProductByOrder(product, orderBO.UUID)

		// If one product has a problem => rollback
		if !outcomeProduct && errProduct != nil {

			orderDTO.UUID = orderBO.UUID
			DeleteOrder(orderDTO)
		}
	}

	return orderDTO, nil
}

func DeleteOrder(orderDTO DTO.Order) (bool, error) {

	var orderBO BO.Order

	orderBO.UUID = orderDTO.UUID
	outcomeOrder, errOrder := Database.DeleteOrder(orderBO)

	if !outcomeOrder && errOrder != nil {

		return false, errOrder
	}
	// products deleted by mysql cascade

	return true, nil
}
