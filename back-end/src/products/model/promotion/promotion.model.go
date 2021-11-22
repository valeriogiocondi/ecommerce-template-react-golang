package promotion

import (
	"time"

	"github.com/google/uuid"
)

type Promotion struct {
	Id                 uint64    `gorm:"column:id; autoIncrement"`
	UUID               uuid.UUID `gorm:"column:uuid; type:uuid; primaryKey"`
	ProductId          uuid.UUID `gorm:"column:product_id; "type:uuid;"`
	Name               string    `gorm:"column:name"`
	DiscountPercentage uint16    `gorm:"column:discount_percentage"`
	CreatedAt          time.Time `gorm:"column:date_create"`
}
