package database

import (
	uuid "github.com/google/uuid"
	"gorm.io/gorm"

	Mysql "products/database/mysql"
	BO "products/model"
)

var mysqlDB *gorm.DB

func Init() {

	mysqlDB = Mysql.Connection()
}

func GetMysqlConnection() *gorm.DB {
	return mysqlDB
}

func GetProductList(limit int, offset int) ([]BO.Product, error) {
	// if mysqlDB == nil
	// 	panic
	return Mysql.Product.GetProductList(mysqlDB, limit, offset)
}

func GetProductById(productId uuid.UUID) (BO.Product, error) {

	return Mysql.Product.GetProductById(mysqlDB, productId)
}

func CreateNewProduct(newProduct BO.Product) (bool, error) {

	return Mysql.Product.CreateNewProduct(mysqlDB, newProduct)
}
