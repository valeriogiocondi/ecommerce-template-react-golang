package controllerProduct

import (
	"github.com/google/uuid"

	Database "products/database"
	DTO "products/service/dto"
)

func GetProductList(limit int, offset int) ([]DTO.Product, error) {

	productBO, err := Database.GetProductList(limit, offset)
	productDTO := make([]DTO.Product, len(productBO))

	if err != nil {
		return productDTO, err
	}

	// mapper
	for i, product := range productBO {

		productDTO[i].UUID = product.UUID
		productDTO[i].Name = product.Name
		productDTO[i].Price = product.Price
	}

	return productDTO, nil
}

func GetProductById(productId uuid.UUID) (DTO.Product, error) {

	var productDTO DTO.Product
	productBO, err := Database.GetProductById(productId)

	if err != nil {
		return productDTO, err
	}

	productDTO.UUID = productBO.UUID
	productDTO.Name = productBO.Name
	productDTO.Price = productBO.Price

	return productDTO, nil
}
