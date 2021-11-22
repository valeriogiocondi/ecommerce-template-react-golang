package customer

import (
	ModelCustomer "customers/model/customer"
	"fmt"

	uuid "github.com/google/uuid"
	"gorm.io/gorm"
)

type Customer = ModelCustomer.Customer

func GetCustomerList(db *gorm.DB, limit int, offset int) ([]Customer, error) {

	var customers []Customer
	result := db.Table("customers").Limit(limit).Offset(offset).Find(&customers)

	return customers, errorGorm(result.Error)
}

func GetCustomerById(db *gorm.DB, customerId uuid.UUID) (Customer, error) {

	var customer Customer
	result := db.Table("customers").First(&customer, "uuid = ?", customerId)

	return customer, errorGorm(result.Error)
}

func CreateNewCustomer(db *gorm.DB, newCustomer Customer) (bool, error) {

	result := db.Table("customers").Create(&newCustomer)

	return (result.Error == nil), errorGorm(result.Error)
}

func errorGorm(err error) error {

	if err != nil {
		return fmt.Errorf("GORM - %s", err)
	}

	return nil
}
