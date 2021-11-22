package product

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	Id                 uint64    `gorm:"column:id; autoIncrement; primaryKey"`
	UUID               uuid.UUID `gorm:"column:uuid; type:uuid; unique"`
	OrderId            uuid.UUID `gorm:"column:order_id; type:uuid;"`
	ProductId          uuid.UUID `gorm:"column:product_id; type:uuid;"`
	OriginalPrice      float64   `gorm:"column:original_price"`
	DiscountPercentage uint16    `gorm:"column:discount_percentage"`
	Price              float64   `gorm:"column:price"`
	CreatedAt          time.Time `gorm:"column:date_create"`
}
