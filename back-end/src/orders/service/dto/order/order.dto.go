package orderDTO

import (
	"github.com/google/uuid"

	productDTO "orders/service/dto/product"
)

type ProductDTO = productDTO.Product

type Order struct {
	UUID               uuid.UUID    `json:"id"`
	CustomerId         uuid.UUID    `json:"customerId"`
	PromotionId        uuid.UUID    `json:"promotionId"`
	OriginalPrice      float64      `json:"originalPrice"`
	DiscountPercentage uint16       `json:"discountPercentage"`
	TotalPrice         float64      `json:"totalPrice"`
	Products           []ProductDTO `json:"products"`
}
