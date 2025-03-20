package app

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/yasintaqvi/banking-app-with-hexagonal-architecture/dto"
	"github.com/yasintaqvi/banking-app-with-hexagonal-architecture/service"
	"net/http"
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
