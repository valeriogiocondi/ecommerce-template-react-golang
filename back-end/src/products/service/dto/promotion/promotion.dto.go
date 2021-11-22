package promotionDTO

import (
	"github.com/google/uuid"
)

type Promotion struct {
	UUID               uuid.UUID `json:"id"`
	ProductId          uuid.UUID `json:"productId"`
	Name               string    `json:"name"`
	DiscountPercentage uint16    `json:"discountPercentage"`
}
