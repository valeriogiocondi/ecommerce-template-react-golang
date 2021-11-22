package mysql

import (
	uuid "github.com/google/uuid"
	"gorm.io/gorm"

	ConnectionMysql "customers/database/mysql/_connection"
	CustomerMysql "customers/database/mysql/customer"
	BO "customers/model"
)

// Definition
type getCustomerList_type func(db *gorm.DB, limit int, offset int) ([]BO.Customer, error)
type getCustomerById_type func(db *gorm.DB, customerId uuid.UUID) (BO.Customer, error)
type createNewCustomer_type func(db *gorm.DB, newCustomer BO.Customer) (bool, error)

type CustomerFactory struct {
	GetCustomerList   getCustomerList_type
	GetCustomerById   getCustomerById_type
	CreateNewCustomer createNewCustomer_type
}

// Export
var Connection = ConnectionMysql.Init
var Customer = CustomerFactory{
	GetCustomerList:   CustomerMysql.GetCustomerList,
	GetCustomerById:   CustomerMysql.GetCustomerById,
	CreateNewCustomer: CustomerMysql.CreateNewCustomer,
}
