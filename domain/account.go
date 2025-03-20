package domain

import (
	"github.com/yasintaqvi/banking-app-with-hexagonal-architecture/dto"
	"github.com/yasintaqvi/banking-app-with-hexagonal-architecture/errs"
)

type Account struct {
	AccountId   string  `db:"account_id"`
	CustomerId  string  `db:"customer_id"`
	OpeningDate string  `db:"opening_date"`
	AccountType string  `db:"account_type"`
	Amount      float64 `db:"amount"`
	Status      string  `db:"status"`
}

type AccountRepository interface {
	Save(account *Account) (*Account, *errs.AppError)
}

func (account Account) ToCreateAccountResponseDto() *dto.CreateAccountResponseDto {
	return &dto.CreateAccountResponseDto{
		AccountId: account.AccountId,
	}
}
