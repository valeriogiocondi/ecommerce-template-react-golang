package product

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	Id        uint64    `gorm:"column:id; autoIncrement"`
	UUID      uuid.UUID `gorm:"column:uuid; type:uuid; primaryKey"`
	Name      string    `gorm:"column:name"`
	Price     float64   `gorm:"column:price"`
	CreatedAt time.Time `gorm:"column:date_create"`
}
