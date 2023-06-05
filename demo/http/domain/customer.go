package domain

type Customer struct {
	Id          string `json:"id" db:"customer_id"`
	Name        string `json:"full_name" db:"name"`
	City        string `json:"city,omitempty" db:"city"`
	Zipcode     string `json:"zipcode" db:"zipcode"`
	DateOfBirth string `json:"dob" db:"date_of_birth"`
	Status      string `json:"status" db:"status"`
}
