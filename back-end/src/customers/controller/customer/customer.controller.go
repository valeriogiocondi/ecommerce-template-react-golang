package controllerCustomer

import (
	uuid "github.com/google/uuid"

	Database "customers/database"
	BO "customers/model/customer"
	DTO "customers/service/dto"
)

func GetCustomerList(limit int, offset int) ([]DTO.Customer, error) {

	customerBO, err := Database.GetCustomerList(limit, offset)
	customerDTO := make([]DTO.Customer, len(customerBO))

	if err != nil {
		return customerDTO, err
	}

	// mapper
	for i, customer := range customerBO {

		customerDTO[i].UUID = customer.UUID
		customerDTO[i].FirebaseId = customer.FirebaseId
		customerDTO[i].FirstName = customer.FirstName
		customerDTO[i].LastName = customer.LastName
		customerDTO[i].Email = customer.Email
		customerDTO[i].Tel = customer.Tel
		customerDTO[i].Address = customer.Address
		customerDTO[i].Num = customer.Num
		customerDTO[i].Cap = customer.Cap
		customerDTO[i].City = customer.City
		customerDTO[i].State = customer.State
	}

	return customerDTO, nil
}

func GetCustomerById(customerId uuid.UUID) (DTO.Customer, error) {

	var customerDTO DTO.Customer
	customerBO, err := Database.GetCustomerById(customerId)

	if err != nil {
		return customerDTO, err
	}

	// mapper
	customerDTO.UUID = customerBO.UUID
	customerDTO.FirebaseId = customerBO.FirebaseId
	customerDTO.FirstName = customerBO.FirstName
	customerDTO.LastName = customerBO.LastName
	customerDTO.Email = customerBO.Email
	customerDTO.Tel = customerBO.Tel
	customerDTO.Address = customerBO.Address
	customerDTO.Num = customerBO.Num
	customerDTO.Cap = customerBO.Cap
	customerDTO.City = customerBO.City
	customerDTO.State = customerBO.State

	return customerDTO, nil
}

func CreateNewCustomer(customerDTO DTO.Customer) (DTO.Customer, error) {

	customerDTO.UUID = uuid.New()

	var customerBO BO.Customer
	customerBO.UUID = customerDTO.UUID
	customerBO.FirebaseId = customerDTO.FirebaseId
	customerBO.FirstName = customerDTO.FirstName
	customerBO.LastName = customerDTO.LastName
	customerBO.Email = customerDTO.Email
	customerBO.Tel = customerDTO.Tel
	customerBO.Address = customerDTO.Address
	customerBO.Num = customerDTO.Num
	customerBO.Cap = customerDTO.Cap
	customerBO.City = customerDTO.City
	customerBO.State = customerDTO.State

	// Create new customer
	outcomeCustomer, errCustomer := Database.CreateNewCustomer(customerBO)

	if !outcomeCustomer && errCustomer != nil {
		return DTO.Customer{}, errCustomer
	}

	return customerDTO, nil
}
