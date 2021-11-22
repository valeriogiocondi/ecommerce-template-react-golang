package database

import (
	uuid "github.com/google/uuid"
	"gorm.io/gorm"

	Mysql "orders/database/mysql"
	BO "orders/model"
)

var mysqlDB *gorm.DB

func Init() {

	mysqlDB = Mysql.Connection()
}

func GetMysqlConnection() *gorm.DB {
	return mysqlDB
}

// ORDER

func GetOrderListByCustomerId(customerId uuid.UUID, limit int, offset int) ([]BO.Order, error) {

	return Mysql.Order.GetOrderListByCustomerId(mysqlDB, customerId, limit, offset)
}

func GetOrderById(orderId uuid.UUID) (BO.Order, error) {

	return Mysql.Order.GetOrderById(mysqlDB, orderId)
}

func CreateNewOrder(newOrder BO.Order) (bool, error) {

	return Mysql.Order.CreateNewOrder(mysqlDB, newOrder)
}

func DeleteOrder(order BO.Order) (bool, error) {

	return Mysql.Order.DeleteOrder(mysqlDB, order)
}

// PRODUCT

func GetProductListByOrderId(id uuid.UUID) ([]BO.Product, error) {

	return Mysql.Product.GetProductListByOrderId(mysqlDB, id)
}

func CreateNewProductByOrder(newProduct BO.Product) (bool, error) {

	return Mysql.Product.CreateNewProductByOrder(mysqlDB, newProduct)
}

func DeleteProduct(product BO.Product) (bool, error) {

	return Mysql.Product.DeleteProduct(mysqlDB, product)
}
