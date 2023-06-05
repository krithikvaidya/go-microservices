package domain

type StubCustomerRepository struct{}

func (cr StubCustomerRepository) FindAll() ([]Customer, error) {
	return []Customer{
		{Id: "1", Name: "Abhay", City: "Blr", Zipcode: "123456", DateOfBirth: "10/12/1970", Status: "1"},
		{Id: "2", Name: "Sumit", City: "Mum", Zipcode: "220022", DateOfBirth: "10/04/1970", Status: "1"},
		{Id: "3", Name: "Ashish", City: "Del", Zipcode: "110011", DateOfBirth: "10/02/1970", Status: "0"},
	}, nil
}

func (cr StubCustomerRepository) ById(id string) (*Customer, error) {
	return &Customer{Id: "3", Name: "Ashish", City: "Del", Zipcode: "110011", DateOfBirth: "10/02/1970", Status: "0"}, nil
}

func NewStubCustomerRepository() StubCustomerRepository {
	return StubCustomerRepository{}
}
