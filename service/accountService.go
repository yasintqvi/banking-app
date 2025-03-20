package service

import (
	"github.com/yasintaqvi/banking-app-with-hexagonal-architecture/domain"
	"github.com/yasintaqvi/banking-app-with-hexagonal-architecture/dto"
	"github.com/yasintaqvi/banking-app-with-hexagonal-architecture/errs"
	"github.com/yasintaqvi/banking-app-with-hexagonal-architecture/logger"
	"time"
)

type AccountService interface {
	CreateAccount(createAccountDto dto.CreateAccountRequestDto) (*dto.CreateAccountResponseDto, *errs.AppError)
	CreateTransaction(*dto.CreateTransactionRequestDto) (*dto.CreateTransactionResponse, *errs.AppError)
}

type DefaultAccountService struct {
	repo domain.AccountRepository
}

func (accountService DefaultAccountService) CreateAccount(createAccountDto dto.CreateAccountRequestDto) (*dto.CreateAccountResponseDto, *errs.AppError) {

	err := createAccountDto.Validate()

	if err != nil {
		return nil, err
	}

	newAccount := &domain.Account{
		AccountId:   "",
		CustomerId:  createAccountDto.CustomerId,
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
		Amount:      createAccountDto.Amount,
		AccountType: createAccountDto.AccountType,
		Status:      "1",
	}

	if account, err := accountService.repo.Save(newAccount); err != nil {
		return nil, err
	} else {
		response := account.ToCreateAccountResponseDto()
		return response, nil
	}
}

func (accountService DefaultAccountService) CreateTransaction(createTransactionDto *dto.CreateTransactionRequestDto) (*dto.CreateTransactionResponse, *errs.AppError) {

	createTransactionDto.Validate()

	account, _ := accountService.repo.FindById(createTransactionDto.AccountId)

	if !account.CanWithdraw(createTransactionDto.Amount) {
		logger.Error("Insufficient balance in the account")
		return nil, errs.NewValidationError("Insufficient balance in the account")
	}

	newTransaction := &domain.Transaction{
		TransactionId:   "",
		AccountId:       createTransactionDto.AccountId,
		Amount:          createTransactionDto.Amount,
		TransactionType: createTransactionDto.TransactionType,
		TransactionDate: createTransactionDto.TransactionDate,
	}

	transaction, appError := accountService.repo.SaveTransaction(newTransaction)
	if appError != nil {
		return nil, appError
	}
	response := transaction.ToDto()

	return response, nil
}

func NewAccountService(repo domain.AccountRepository) *DefaultAccountService {
	return &DefaultAccountService{repo: repo}
}
