package dto

import "github.com/yasintaqvi/banking-app-with-hexagonal-architecture/errs"

const WITHDRAWAL = "withdrawal"
const DEPOSIT = "deposit"

type CreateTransactionRequestDto struct {
	CustomerId      string  `json:"customer_id"`
	AccountId       string  `json:"account_id"`
	Amount          float64 `json:"amount"`
	TransactionType string  `json:"transaction_type"`
	TransactionDate string  `json:"transaction_date"`
}

func (r CreateTransactionRequestDto) IsTransactionTypeWithdrawal() bool {
	return r.TransactionType == WITHDRAWAL
}

func (r CreateTransactionRequestDto) IsTransactionTypeDeposit() bool {
	return r.TransactionType == DEPOSIT
}

func (r CreateTransactionRequestDto) Validate() *errs.AppError {
	if !r.IsTransactionTypeWithdrawal() && !r.IsTransactionTypeDeposit() {
		return errs.NewValidationError("Transaction type can only be deposit or withdrawal")
	}
	if r.Amount < 0 {
		return errs.NewValidationError("Amount cannot be less than zero")
	}
	return nil
}
