package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	return CustomerRepositoryStub{
		[]Customer{
			{"1", "yasin", "lahijan", "10001", "2000-14-20", true},
			{"2", "alireza", "lahijan", "10001", "2001-01-20", true},
		},
	}
}
