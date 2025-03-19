package app

import (
	"github.com/gorilla/mux"
	"github.com/yasintaqvi/banking-app-with-hexagonal-architecture/domain"
	"github.com/yasintaqvi/banking-app-with-hexagonal-architecture/service"
	"log"
	"net/http"
)

func Start() {
	router := mux.NewRouter()

	//customerHandler := CustomerHandler{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	customerHandler := CustomerHandler{service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	router.HandleFunc("/api/customers", customerHandler.GetAllCustomers).Methods("GET")
	router.HandleFunc("/api/customers/{customer_id:[0-9]+}", customerHandler.GetCustomer).Methods("GET")

	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
