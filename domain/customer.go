package domain

import (
	"banking/dto"
	"banking/errs"
)

type Customer struct {
	ID          string `db:"customer_id"` // для помощи sqlx.StructScan()
	Name        string
	City        string
	Zipcode     string
	DateofBirth string `db:"date_of_birth"` // для помощи sqlx.StructScan()
	Status      string
}

func (c Customer) statusAsText() string {
	statusAsText := "active"
	if c.Status == "0" {
		statusAsText = "inactive"
	}
	return statusAsText
}

func (c Customer) ToDto() dto.CustomerResponse {
	return dto.CustomerResponse{
		ID:          c.ID,
		Name:        c.Name,
		City:        c.City,
		Zipcode:     c.Zipcode,
		DateofBirth: c.DateofBirth,
		Status:      c.statusAsText(),
	}
}

// создание порта
type CustomerRepository interface {
	// status == 1; status == 0; status == ""
	FindAll(status string) ([]Customer, *errs.AppErr)
	ById(string) (*Customer, *errs.AppErr)
}
