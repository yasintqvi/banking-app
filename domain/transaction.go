package domain

import (
	"github.com/yasintaqvi/banking-app-with-hexagonal-architecture/dto"
	"github.com/yasintaqvi/banking-app-with-hexagonal-architecture/errs"
)

type Transaction struct {
	TransactionId   string  `db:"transaction_id"`
	AccountId       string  `db:"account_id"`
	Amount          float64 `db:"amount"`
	TransactionType string  `db:"transaction_type"`
	TransactionDate string  `db:"transaction_date"`
}

const WITHDRAWAL = "withdrawal"

func (transaction Transaction) IsWithdrawal() bool {
	return transaction.TransactionType == WITHDRAWAL
}

func (transaction Transaction) ToDto() *dto.CreateTransactionResponse {
	return &dto.CreateTransactionResponse{
		TransactionId:   transaction.TransactionId,
		AccountId:       transaction.AccountId,
		Amount:          transaction.Amount,
		TransactionType: transaction.TransactionType,
		TransactionDate: transaction.TransactionDate,
	}
}

type TransactionRepository interface {
	Save(Transaction) (*Transaction, *errs.AppError)
}
