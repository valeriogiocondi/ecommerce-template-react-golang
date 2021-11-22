package order

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	Id                 uint64    `gorm:"column:id; autoIncrement"`
	UUID               uuid.UUID `gorm:"column:uuid; type:uuid; primaryKey"`
	CustomerId         uuid.UUID `gorm:"column:customer_id; type:uuid;"`
	PromotionId        uuid.UUID `gorm:"column:promotion_id; type:uuid;"`
	OriginalPrice      float64   `gorm:"column:original_price"`
	DiscountPercentage uint16    `gorm:"column:discount_percentage"`
	TotalPrice         float64   `gorm:"column:total_price"`
	CreatedAt          time.Time `gorm:"column:date_create"`
}
