package productDTO

import (
	"github.com/google/uuid"
)

type Product struct {
	UUID               uuid.UUID `json:"id"`
	OrderId            uuid.UUID `json:"orderId"`
	OriginalPrice      float64   `json:"originalPrice"`
	DiscountPercentage uint16    `json:"discountPercentage"`
	Price              float64   `json:"price"`
}
