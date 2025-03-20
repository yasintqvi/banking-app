package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/yasintaqvi/banking-app-with-hexagonal-architecture/domain"
	"github.com/yasintaqvi/banking-app-with-hexagonal-architecture/logger"
	"github.com/yasintaqvi/banking-app-with-hexagonal-architecture/service"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"time"
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

	client := getClientDb()

	customerRepo := domain.NewCustomerRepositoryDb(client)
	accountRepo := domain.NewAccountRepositoryDb(client)

	//customerHandler := CustomerHandler{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	customerHandler := CustomerHandler{service.NewCustomerService(customerRepo)}
	accountHandler := AccountHandler{service.NewAccountService(accountRepo)}

	router.HandleFunc("/api/customers", customerHandler.GetAllCustomers).Methods("GET")
	router.HandleFunc("/api/customers/{customer_id:[0-9]+}", customerHandler.GetCustomer).Methods("GET")
	router.HandleFunc("/api/customers/{customer_id:[0-9]+}/account", accountHandler.CreateAccount).Methods("POST")

	serverAddress := fmt.Sprintf("%s:%s", os.Getenv("SERVER_ADDRESS"), os.Getenv("SERVER_PORT"))

	logger.Info(fmt.Sprintf("Starting server on %s", serverAddress))

	log.Fatal(http.ListenAndServe(serverAddress, router))
}

func getClientDb() *sqlx.DB {
	dbInfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))

	client, err := sqlx.Connect("mysql", dbInfo)

	if err != nil {
		logger.Error(err.Error(), zap.Error(err))
		panic(err)
	}

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return client
}
