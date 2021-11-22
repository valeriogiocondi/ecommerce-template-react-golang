package gateway

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	uuid "github.com/google/uuid"
	"github.com/gorilla/mux"

	ControllerCustomer "customers/controller/customer"
	DTO "customers/service/dto"
	REST "customers/utils/rest"
)

var Restify = REST.Restify

func Init() {

	r := mux.NewRouter()
	r.
		Methods(http.MethodGet).
		Path("/").
		HandlerFunc(HandlerHome)

	r.
		Methods(http.MethodGet).
		Path("/customer").
		Queries("offset", "{offset:[0-9]+}", "limit", "{limit:[0-9]+}").
		HandlerFunc(getCustomerList)

	r.
		Methods(http.MethodGet).
		Path("/customer/{customer_id}").
		HandlerFunc(GetCustomerById)

	r.
		Methods(http.MethodPost).
		Path("/customer").
		HandlerFunc(createNewCustomer)

	r.NotFoundHandler = http.HandlerFunc(handler404)

	http.Handle("/", r)
}

func HandlerHome(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "300 Erbe - Customer List")
}

func getCustomerList(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()

	vars := mux.Vars(r)
	limit, _ := strconv.Atoi(vars["limit"])
	offset, _ := strconv.Atoi(vars["offset"])

	// call controller
	payload, err := ControllerCustomer.GetCustomerList(limit, offset)
	Restify(w, payload, err)
}

func GetCustomerById(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()

	vars := mux.Vars(r)
	customerId, _ := uuid.Parse(vars["customer_id"])

	// call controller
	payload, err := ControllerCustomer.GetCustomerById(customerId)
	Restify(w, payload, err)
}

func createNewCustomer(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()

	var customerDTO DTO.Customer

	err := json.NewDecoder(r.Body).Decode(&customerDTO)
	if err != nil {
		Restify(w, nil, err)
	}

	// call controller
	payload, err := ControllerCustomer.CreateNewCustomer(customerDTO)
	Restify(w, payload, err)
}

func handler404(w http.ResponseWriter, r *http.Request) {

	Restify(w, nil, errors.New(http.StatusText(http.StatusNotFound)))
}
