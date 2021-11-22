package mainTesting

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"

	app "orders/gateway"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", app.HandlerHome).Methods("GET")
	return router
}

func TestRest(t *testing.T) {
	request, _ := http.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)

	assert.Equal(t, 200, response.Code, "")
}
