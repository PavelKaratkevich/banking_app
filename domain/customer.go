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
	FindAll() ([]Customer, *errs.AppErr)
	ById(string) (*Customer, *errs.AppErr)
}
