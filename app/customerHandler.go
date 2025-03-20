package app

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/yasintaqvi/banking-app-with-hexagonal-architecture/service"
	"net/http"
)

type CustomerHandler struct {
	service service.CustomerService
}

func (handler CustomerHandler) GetAllCustomers(writer http.ResponseWriter, request *http.Request) {
	customers, err := handler.service.GetAllCustomer()

	if err != nil {
		writeHttpResponse(writer, err.Code, err.AsMessage())
	} else {
		writeHttpResponse(writer, http.StatusOK, customers)
	}
}

func (handler CustomerHandler) GetCustomer(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)

	id := vars["customer_id"]

	customer, err := handler.service.GetCustomer(id)

	request.Header.Set("Content-Type", "application/json")

	if err != nil {
		writeHttpResponse(writer, err.Code, err.AsMessage())
	} else {
		writeHttpResponse(writer, http.StatusOK, customer)
	}
}

func writeHttpResponse(writer http.ResponseWriter, code int, data interface{}) {

	writer.Header().Set("Content-Type", "application/json")

	writer.WriteHeader(code)

	err := json.NewEncoder(writer).Encode(data)

	if err != nil {
		panic(err)
	}
}
