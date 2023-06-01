package service

import "learning-http/domain"

type CustomerService struct{}

func (c *CustomerService) GetAllCustomers() []domain.Customer {
	// make a call to repo and FindAll customers
	repo := domain.NewStubCustomerRepository()
	return repo.FindAll()
}

func NewCustomerService() CustomerService {
	return CustomerService{}
}
