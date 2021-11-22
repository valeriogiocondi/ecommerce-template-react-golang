package gateway

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/gorilla/mux"

	ControllerProduct "products/controller/product"
	REST "products/utils/rest"
)

var Restify = REST.Restify

func Init() {

	r := mux.NewRouter()
	r.
		Methods(http.MethodGet).
		Path("/").
		HandlerFunc(handlerHome)

	r.
		Methods(http.MethodGet).
		Path("/product").
		Queries("offset", "{offset:[0-9]+}", "limit", "{limit:[0-9]+}").
		HandlerFunc(getProductList)

	r.
		Methods(http.MethodGet).
		Path("/product/{product_id}").
		HandlerFunc(getProductById)

	r.NotFoundHandler = http.HandlerFunc(handler404)

	http.Handle("/", r)
}

func handlerHome(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "300 Erbe - Product List")
}

func getProductList(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()

	vars := mux.Vars(r)
	limit, _ := strconv.Atoi(vars["limit"])
	offset, _ := strconv.Atoi(vars["offset"])

	// call controller
	payload, err := ControllerProduct.GetProductList(limit, offset)
	Restify(w, payload, err)
}

func getProductById(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()

	vars := mux.Vars(r)
	productId, _ := uuid.Parse(vars["product_id"])

	// call controller
	payload, err := ControllerProduct.GetProductById(productId)
	Restify(w, payload, err)
}

func handler404(w http.ResponseWriter, r *http.Request) {

	Restify(w, nil, errors.New(http.StatusText(http.StatusNotFound)))
}
