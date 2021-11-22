package order

import (
	"fmt"

	uuid "github.com/google/uuid"
	"gorm.io/gorm"

	BO "orders/model"
)

func GetOrderListByCustomerId(db *gorm.DB, customerId uuid.UUID, limit int, offset int) ([]BO.Order, error) {

	var orders []BO.Order
	result := db.Table("orders").Where("customer_id = ?", customerId).Limit(limit).Offset(offset).Find(&orders)

	return orders, errorGorm(result.Error)
}

func GetOrderById(db *gorm.DB, orderId uuid.UUID) (BO.Order, error) {

	var order BO.Order
	result := db.Table("orders").First(&order, "uuid = ?", orderId)

	return order, errorGorm(result.Error)
}

func CreateNewOrder(db *gorm.DB, newOrder BO.Order) (bool, error) {

	result := db.Table("orders").Create(&newOrder)

	return (result.Error == nil), errorGorm(result.Error)
}

func DeleteOrder(db *gorm.DB, order BO.Order) (bool, error) {

	result := db.Table("orders").Where("uuid = ?", order.UUID).Delete(&order)

	return (result.Error == nil), errorGorm(result.Error)
}

func errorGorm(err error) error {

	if err != nil {
		return fmt.Errorf("GORM - %s", err)
	}

	return nil
}
