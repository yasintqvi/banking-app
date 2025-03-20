package dto

import (
	"github.com/yasintaqvi/banking-app-with-hexagonal-architecture/errs"
	"strings"
)

type CreateAccountRequestDto struct {
	CustomerId  string  `json:"customer_id"`
	Amount      float64 `json:"amount"`
	AccountType string  `json:"account_type"`
}

func (createAccountRequestDto CreateAccountRequestDto) Validate() *errs.AppError {

	if createAccountRequestDto.CustomerId == "" {
		return errs.NewValidationError("customer id cannot be empty")
	}

	if createAccountRequestDto.Amount <= 100 {
		return errs.NewValidationError("amount cannot be less than 100")
	}

	if strings.ToLower(createAccountRequestDto.AccountType) != "checking" && strings.ToLower(createAccountRequestDto.AccountType) != "saving" {
		return errs.NewValidationError("invalid account type")
	}

	return nil
}
