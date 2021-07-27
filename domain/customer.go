package domain

type Customer struct {
	ID          string
	Name        string
	City        string
	Zipcode     string
	DateofBirth string
	Status      string
}

// создание порта
type CustomerRepository interface {
	FindAll() ([]Customer, error)
}
