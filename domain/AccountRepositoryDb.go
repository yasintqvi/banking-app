package domain

import (
	"database/sql"
	"errors"
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

func (repo *AccountRepositoryDb) SaveTransaction(transaction *Transaction) (*Transaction, *errs.AppError) {
	beginTx, err := repo.client.Begin()

	if err != nil {
		logger.Error("Error while starting a new transaction for bank account transaction", zap.Error(err))
		return nil, errs.NewInternalServerError(err.Error())
	}

	result, err := beginTx.Exec("INSERT INTO transactions (account_id, amount, transaction_type, transaction_date) VALUES(?, ?, ?, ?)", transaction.AccountId, transaction.Amount, transaction.TransactionType, transaction.TransactionDate)

	if err != nil {
		logger.Error("Error while inserting new transaction", zap.Error(err))
		return nil, errs.NewInternalServerError(err.Error())
	}

	if transaction.IsWithdrawal() {
		_, err := beginTx.Exec("UPDATE accounts SET amount = amount - ? WHERE account_id = ?", transaction.Amount, transaction.AccountId)
		if err != nil {
			err := beginTx.Rollback()
			if err != nil {
				return nil, nil
			}
			logger.Error("Error while updating balance", zap.Error(err))
			return nil, errs.NewInternalServerError(err.Error())
		}
	} else {
		_, err := beginTx.Exec("UPDATE accounts SET amount = amount + ? WHERE account_id = ?", transaction.Amount, transaction.AccountId)
		if err != nil {
			err := beginTx.Rollback()
			if err != nil {
				return nil, nil
			}
			logger.Error("Error while updating balance", zap.Error(err))
			return nil, errs.NewInternalServerError(err.Error())
		}
	}

	err = beginTx.Commit()

	if err != nil {
		logger.Error("Error while commiting transaction", zap.Error(err))
		return nil, errs.NewInternalServerError(err.Error())
	}

	transactionId, err := result.LastInsertId()

	if err != nil {
		logger.Error("Error while inserting new transaction", zap.Error(err))
		return nil, errs.NewInternalServerError(err.Error())
	}

	transaction.TransactionId = strconv.FormatInt(transactionId, 10)

	account, _ := repo.FindById(transaction.AccountId)

	transaction.Amount = account.Amount

	return transaction, nil
}

func (repo *AccountRepositoryDb) FindById(id string) (*Account, *errs.AppError) {
	var account Account
	err := repo.client.Get(&account, "SELECT  account_id, customer_id, amount, opening_date, account_type FROM accounts WHERE account_id = ?", id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			logger.Error("bank account not defined", zap.Error(err))
			return nil, errs.NewHttpNotFoundError(err.Error())
		}

		logger.Error(err.Error(), zap.Error(err))
		return nil, errs.NewInternalServerError(err.Error())
	}

	return &account, nil
}

func NewAccountRepositoryDb(client *sqlx.DB) *AccountRepositoryDb {
	return &AccountRepositoryDb{client: client}
}
