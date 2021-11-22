package gateway

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	uuid "github.com/google/uuid"
	"github.com/gorilla/mux"

	ControllerOrder "orders/controller/order"
	ControllerProduct "orders/controller/product"
	DTO "orders/service/dto"
	REST "orders/utils/rest"
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
		Path("/order/customer/{customer_id}/offset/{offset:[0-9]+}/limit/{limit:[0-9]+}").
		HandlerFunc(getOrderListByCustomerId)

	r.
		Methods(http.MethodGet).
		Path("/order/{order_id}").
		HandlerFunc(getOrderById)

	r.
		Methods(http.MethodGet).
		Path("/order/{order_id}/products").
		HandlerFunc(getProductListByOrderId)

	r.
		Methods(http.MethodPost).
		Path("/order").
		HandlerFunc(createNewOrder)

	r.NotFoundHandler = http.HandlerFunc(handler404)

	http.Handle("/", r)
}

func HandlerHome(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "300 Erbe - Order List")
}

func getOrderListByCustomerId(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()

	vars := mux.Vars(r)
	customerId, _ := uuid.Parse(vars["customer_id"])
	limit, _ := strconv.Atoi(vars["limit"])
	offset, _ := strconv.Atoi(vars["offset"])

	// call controller
	payload, err := ControllerOrder.GetOrderListByCustomerId(customerId, limit, offset)
	Restify(w, payload, err)
}

func getOrderById(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()

	vars := mux.Vars(r)
	orderId, _ := uuid.Parse(vars["order_id"])

	// call controller
	payload, err := ControllerOrder.GetOrderById(orderId)
	Restify(w, payload, err)
}

func getProductListByOrderId(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()

	vars := mux.Vars(r)
	orderId, _ := uuid.Parse(vars["order_id"])

	// call controller
	payload, err := ControllerProduct.GetProductListByOrderId(orderId)
	Restify(w, payload, err)
}

func createNewOrder(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()

	var orderDTO DTO.Order

	err := json.NewDecoder(r.Body).Decode(&orderDTO)
	if err != nil {
		Restify(w, nil, err)
	}

	// call controller
	payload, err := ControllerOrder.CreateNewOrder(orderDTO)
	Restify(w, payload, err)
}

func handler404(w http.ResponseWriter, r *http.Request) {

	Restify(w, nil, errors.New(http.StatusText(http.StatusNotFound)))
}
