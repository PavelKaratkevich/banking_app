package service

import (
	"banking/domain"
	"banking/errs"
)

// Primary port for interaction with users
type CustomerService interface {
	GetAllCustomer(string) ([]domain.Customer, *errs.AppErr)
	GetCustomer(string) (*domain.Customer, *errs.AppErr)
}

// port implementation which has a dependency of the Repository
type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

// применяю предусмотренный интерфейсом CustomerService метод к DefaultCustomerService
func (s DefaultCustomerService) GetAllCustomer(status string) ([]domain.Customer, *errs.AppErr) {
	if status == "active" {
		status = "1"
	} else if status == "inactive" {
		status = "0"
	} else {
		status = ""
	}
	// применяю метод FindAll() к CustomerRepository, тем самым получая данные от сервера (FindAll() ) и передавая их пользователю (GetAllCustomers() )
	return s.repo.FindAll(status)
}

func (s DefaultCustomerService) GetCustomer(id string) (*domain.Customer, *errs.AppErr) {
	return s.repo.ById(id)
}

// вспомогательная функция, которая внедряет зависимость репо от DefaultCustomerService
func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{
		repository,
	}
}
