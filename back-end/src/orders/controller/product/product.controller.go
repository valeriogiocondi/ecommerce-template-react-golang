package controllerProduct

import (
	uuid "github.com/google/uuid"

	Database "orders/database"
	BO "orders/model"
	DTO "orders/service/dto"
)

func GetProductListByOrderId(id uuid.UUID) ([]DTO.Product, error) {

	productBO, err := Database.GetProductListByOrderId(id)
	productDTO := make([]DTO.Product, len(productBO))

	if err != nil {
		return productDTO, err
	}

	// mapper
	for i, product := range productBO {

		productDTO[i].UUID = product.ProductId
		productDTO[i].OrderId = product.OrderId
		productDTO[i].OriginalPrice = product.OriginalPrice
		productDTO[i].DiscountPercentage = product.DiscountPercentage
		productDTO[i].Price = product.Price
	}

	return productDTO, nil
}

func CreateNewProductByOrder(productDTO DTO.Product, orderId uuid.UUID) (bool, error) {

	var productBO BO.Product

	productBO.UUID = uuid.New()
	productBO.OrderId = orderId
	productBO.ProductId = productDTO.UUID
	productBO.OriginalPrice = productDTO.OriginalPrice
	productBO.DiscountPercentage = productDTO.DiscountPercentage
	// Discount already apply in order
	productBO.Price = productDTO.Price

	return Database.CreateNewProductByOrder(productBO)
}

func DeleteProduct(productDTO DTO.Product) (bool, error) {

	var productBO BO.Product

	productBO.UUID = productDTO.UUID
	outcomeOrder, errOrder := Database.DeleteProduct(productBO)

	if outcomeOrder && errOrder != nil {

		return false, errOrder
	}

	return true, nil
}
