package service

import (
	"learning-http/domain"
	"learning-http/errs"
)

type CustomerService struct {
	repo domain.CustomerRepository
}

func (c *CustomerService) GetAllCustomers() ([]domain.Customer, *errs.AppError) {
	customers, err := c.repo.FindAll()
	if err != nil {
		return nil, err
	}
	for idx, _ := range customers {
		if customers[idx].Status == "1" {
			customers[idx].Status = "active"
		} else {
			customers[idx].Status = "Inactive"
		}
	}
	return customers, err
}

func (c *CustomerService) GetCustomer(id string) (*domain.Customer, *errs.AppError) {
	return c.repo.ById(id)
}

func NewCustomerService(repo domain.CustomerRepository) CustomerService {
	// repo := domain.NewCustomerRepositoryDb()
	return CustomerService{repo}
}
