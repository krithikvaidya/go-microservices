package domain

type StubCustomerRepository struct{}

func (cr *StubCustomerRepository) FindAll() []Customer {
	return []Customer{
		{"1", "Abhay", "Blr", "123456"},
		{"2", "Ashish", "", "110011"},
		{"3", "Rob", "Mum", "220011"},
	}
}

func NewStubCustomerRepository() StubCustomerRepository {
	return StubCustomerRepository{}
}
