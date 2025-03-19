package service

import (
	"github.com/yasintaqvi/banking-app-with-hexagonal-architecture/domain"
	"github.com/yasintaqvi/banking-app-with-hexagonal-architecture/dto"
	"github.com/yasintaqvi/banking-app-with-hexagonal-architecture/errs"
)

type CustomerService interface {
	GetAllCustomer() ([]dto.CustomerResponse, *errs.AppError)
	GetCustomer(customerID string) (*dto.CustomerResponse, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (customerService DefaultCustomerService) GetAllCustomer() ([]dto.CustomerResponse, *errs.AppError) {
	customers, err := customerService.repo.FindAll()

	if err != nil {
		return nil, err
	}

	response := make([]dto.CustomerResponse, 0)

	for _, customer := range customers {
		response = append(response, customer.ToDto())
	}

	return response, nil
}

func (customerService DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError) {
	customer, err := customerService.repo.ById(id)

	if err != nil {
		return nil, err
	}

	response := customer.ToDto()

	return &response, nil
}

func NewCustomerService(repo domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo}
}
