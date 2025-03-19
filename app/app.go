package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/yasintaqvi/banking-app-with-hexagonal-architecture/domain"
	"github.com/yasintaqvi/banking-app-with-hexagonal-architecture/logger"
	"github.com/yasintaqvi/banking-app-with-hexagonal-architecture/service"
	"log"
	"net/http"
	"os"
)

func sanityCheck() {
	envProps := []string{
		"SERVER_ADDRESS",
		"SERVER_PORT",
		"DB_HOST",
		"DB_PORT",
		"DB_USER",
		"DB_PASSWORD",
		"DB_NAME",
	}

	for _, prop := range envProps {
		if os.Getenv(prop) == "" {
			log.Fatalf("Environment variable %s is not set", prop)
		}
	}
}

func Start() {

	sanityCheck()

	router := mux.NewRouter()

	//customerHandler := CustomerHandler{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	customerHandler := CustomerHandler{service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	router.HandleFunc("/api/customers", customerHandler.GetAllCustomers).Methods("GET")
	router.HandleFunc("/api/customers/{customer_id:[0-9]+}", customerHandler.GetCustomer).Methods("GET")

	serverAddress := fmt.Sprintf("%s:%s", os.Getenv("SERVER_ADDRESS"), os.Getenv("SERVER_PORT"))

	logger.Info(fmt.Sprintf("Starting server on %s", serverAddress))

	log.Fatal(http.ListenAndServe(serverAddress, router))
}
