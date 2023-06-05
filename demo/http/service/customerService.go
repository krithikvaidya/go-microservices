package service

import "learning-http/domain"

type CustomerService struct {
	repo domain.CustomerRepositoryDb
}

func (c *CustomerService) GetAllCustomers() ([]domain.Customer, error) {
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

func (c *CustomerService) GetCustomer(id string) (*domain.Customer, error) {
	return c.repo.ById(id)
}

func NewCustomerService(repo domain.CustomerRepositoryDb) CustomerService {
	// repo := domain.NewCustomerRepositoryDb()
	return CustomerService{repo}
}
