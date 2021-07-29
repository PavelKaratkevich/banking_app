package domain

import "banking/errs"

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
	// status == 1; status == 0; status == ""
	FindAll(status string) ([]Customer, *errs.AppErr)
	ById(string) (*Customer, *errs.AppErr)
}
