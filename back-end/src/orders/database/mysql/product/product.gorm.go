package order

import (
	"fmt"

	uuid "github.com/google/uuid"
	"gorm.io/gorm"

	BO "orders/model"
)

func GetProductListByOrderId(db *gorm.DB, id uuid.UUID) ([]BO.Product, error) {

	var products []BO.Product
	result := db.Table("orders_products").Where("order_id = ?", id).Find(&products)

	return products, errorGorm(result.Error)
}

func CreateNewProductByOrder(db *gorm.DB, newProduct BO.Product) (bool, error) {

	result := db.Table("orders_products").Create(&newProduct)

	return (result.Error == nil), errorGorm(result.Error)
}

func DeleteProduct(db *gorm.DB, product BO.Product) (bool, error) {

	result := db.Table("orders_products").Delete(&product)

	return (result.Error == nil), errorGorm(result.Error)
}

func errorGorm(err error) error {

	if err != nil {
		return fmt.Errorf("GORM - %s", err)
	}

	return nil
}
