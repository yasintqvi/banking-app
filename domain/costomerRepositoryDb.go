package domain

import (
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/yasintaqvi/banking-app-with-hexagonal-architecture/errs"
	"github.com/yasintaqvi/banking-app-with-hexagonal-architecture/logger"
	"go.uber.org/zap"
	"time"
)

type CustomerRepositoryDb struct {
	client *sqlx.DB
}

func (customerRepo CustomerRepositoryDb) FindAll() ([]Customer, *errs.AppError) {
	var customers []Customer

	err := customerRepo.client.Select(&customers, "SELECT customer_id, name, city, zipcode, date_of_birth, status FROM customers")

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.NewHttpNotFoundError("No customers found")
		}

		logger.Error("Error while fetching customers from database", zap.Error(err))
		return nil, errs.NewInternalServerError(err.Error())
	}

	return customers, nil
}

func (customerRepo CustomerRepositoryDb) ById(id string) (*Customer, *errs.AppError) {
	var customer Customer

	err := customerRepo.client.Get(&customer, "SELECT customer_id, name, city, zipcode, date_of_birth, status FROM customers WHERE customer_id=?", id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.NewHttpNotFoundError("No customer found")
		}

		logger.Error(err.Error(), zap.Error(err))
		return nil, errs.NewInternalServerError(err.Error())
	}

	return &customer, nil
}

func NewCustomerRepositoryDb() *CustomerRepositoryDb {
	client, err := sqlx.Connect("mysql", "admin:password123@tcp(localhost:3306)/banking")

	if err != nil {
		logger.Error(err.Error(), zap.Error(err))
		panic(err)
	}

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return &CustomerRepositoryDb{client}
}
