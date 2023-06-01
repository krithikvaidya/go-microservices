package service

import "learning-http/domain"

type CustomerService struct{}

func (c *CustomerService) GetAllCustomers() ([]domain.Customer, error) {
	// repo := domain.NewStubCustomerRepository()
	repo := domain.NewCustomerRepositoryDb()
	customers, err := repo.FindAll()
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

func NewCustomerService() CustomerService {
	return CustomerService{}
}
