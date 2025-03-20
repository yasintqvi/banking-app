package domain

import (
	"github.com/jmoiron/sqlx"
	"github.com/yasintaqvi/banking-app-with-hexagonal-architecture/errs"
	"github.com/yasintaqvi/banking-app-with-hexagonal-architecture/logger"
	"go.uber.org/zap"
	"strconv"
)

type AccountRepositoryDb struct {
	client *sqlx.DB
}

func (repo *AccountRepositoryDb) Save(account *Account) (*Account, *errs.AppError) {
	insertAccountQuery := "INSERT INTO accounts (customer_id, opening_date, amount, account_type, status) VALUES (?, ?, ?, ?, ?)"

	result, err := repo.client.Exec(insertAccountQuery, account.CustomerId, account.OpeningDate, account.Amount, account.AccountType, account.Status)

	if err != nil {
		logger.Error(err.Error(), zap.Error(err))
		return nil, errs.NewInternalServerError(err.Error())
	}

	id, err := result.LastInsertId()

	if err != nil {
		logger.Error(err.Error(), zap.Error(err))
		return nil, errs.NewInternalServerError(err.Error())
	}

	account.AccountId = strconv.FormatInt(id, 10)

	return account, nil
}

func NewAccountRepositoryDb(client *sqlx.DB) *AccountRepositoryDb {
	return &AccountRepositoryDb{client: client}
}
