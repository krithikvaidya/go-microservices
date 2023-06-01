package domain

type Customer struct {
	Id      string `json:"id"`
	Name    string `json:"full_name"`
	City    string `json:"city,omitempty"`
	Zipcode string `json:"-"`
}