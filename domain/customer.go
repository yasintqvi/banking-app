package domain

import (
	"github.com/yasintaqvi/banking-app-with-hexagonal-architecture/dto"
	"github.com/yasintaqvi/banking-app-with-hexagonal-architecture/errs"
)

type Customer struct {
	ID          string `json:"id" db:"customer_id"`
	Name        string `json:"name"`
	City        string `json:"city"`
	ZipCode     string `json:"zipCode"`
	DateOfBirth string `json:"dateOfBirth" db:"date_of_birth"`
	Status      bool   `json:"status"`
}

func (customer Customer) statusAsText() string {
	if customer.Status {
		return "Active"
	}

	return "Inactive"
}

func (customer Customer) ToDto() dto.CustomerResponse {
	return dto.CustomerResponse{
		Id:          customer.ID,
		Name:        customer.Name,
		City:        customer.City,
		ZipCode:     customer.ZipCode,
		DateOfBirth: customer.DateOfBirth,
		Status:      customer.statusAsText(),
	}
}

type CustomerRepository interface {
	FindAll() ([]Customer, *errs.AppError)
	ById(id string) (*Customer, *errs.AppError)
}
