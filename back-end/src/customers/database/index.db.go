package database

import (
	uuid "github.com/google/uuid"
	"gorm.io/gorm"

	Mysql "customers/database/mysql"
	BO "customers/model"
)

var mysqlDB *gorm.DB

func Init() {

	mysqlDB = Mysql.Connection()
}

func GetMysqlConnection() *gorm.DB {
	return mysqlDB
}

func GetCustomerList(limit int, offset int) ([]BO.Customer, error) {
	// if mysqlDB == nil
	// 	panic
	return Mysql.Customer.GetCustomerList(mysqlDB, limit, offset)
}

func GetCustomerById(customerId uuid.UUID) (BO.Customer, error) {

	return Mysql.Customer.GetCustomerById(mysqlDB, customerId)
}

func CreateNewCustomer(newCustomer BO.Customer) (bool, error) {

	return Mysql.Customer.CreateNewCustomer(mysqlDB, newCustomer)
}
