package mainTesting

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"

	app "customers/gateway"
	ModelCustomer "customers/model/customer"
)

type Customer = ModelCustomer.Customer

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/customer/{customer_id}", app.GetCustomerById).Methods("GET")
	return router
}

func TestGetCustomerById(t *testing.T) {

	// Create requwest
	request, _ := http.NewRequest("GET", "/customer/934b4b6e-703b-4966-af94-e1d759e7e1e8", nil)

	// Record HTTP request
	response := httptest.NewRecorder()

	// Assign HTTP handler function
	// Dispatch HTTP request
	Router().ServeHTTP(response, request)

	var getRes Customer
	json.NewDecoder(io.Reader(response.Body)).Decode(&getRes)

	// Assertion
	assert.Equal(t, 200, response.Code)

	expectedRes := []byte(`{ "Id":  1, "UUID":  "934b4b6e-703b-4966-af94-e1d759e7e1e8", "FirstName":  "John", "LastName":  "doe", "Email":  "john.doe@gmail.com", "Tel":  "12345", "Address":  "Main Street", "Num":  "1", "Cap":  "90210", "City": "Los Angeles", "State":  "CA", }`)
	assert.Equal(t, expectedRes, getRes)
}
