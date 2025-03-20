package app

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/yasintaqvi/banking-app-with-hexagonal-architecture/dto"
	"github.com/yasintaqvi/banking-app-with-hexagonal-architecture/service"
	"net/http"
	"time"
)

type AccountHandler struct {
	service service.AccountService
}

func (accountHandler AccountHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	customerId, _ := vars["customer_id"]

	var createAccountRequestDto dto.CreateAccountRequestDto

	err := json.NewDecoder(r.Body).Decode(&createAccountRequestDto)

	if err != nil {
		writeHttpResponse(w, http.StatusBadRequest, err.Error())
	} else {
		createAccountRequestDto.CustomerId = customerId
		account, err := accountHandler.service.CreateAccount(createAccountRequestDto)

		if err != nil {
			writeHttpResponse(w, err.Code, err.AsMessage())
		} else {
			writeHttpResponse(w, http.StatusCreated, account)
		}
	}
}

func (accountHandler AccountHandler) CreateTransaction(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	customerId, _ := vars["customer_id"]
	accountId, _ := vars["account_id"]

	// decode incoming request
	var createTransactionDto dto.CreateTransactionRequestDto
	if err := json.NewDecoder(request.Body).Decode(&createTransactionDto); err != nil {
		writeHttpResponse(writer, http.StatusBadRequest, err.Error())
	} else {

		createTransactionDto.AccountId = accountId
		createTransactionDto.CustomerId = customerId
		createTransactionDto.TransactionDate = time.Now().Format("2006-01-02 15:04:05")

		account, appError := accountHandler.service.CreateTransaction(&createTransactionDto)

		if appError != nil {
			writeHttpResponse(writer, appError.Code, appError.AsMessage())
		} else {
			writeHttpResponse(writer, http.StatusOK, account)
		}
	}
}
