package productDTO

import (
	"github.com/google/uuid"
)

type Product struct {
	UUID  uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Price float64   `json:"price"`
}
