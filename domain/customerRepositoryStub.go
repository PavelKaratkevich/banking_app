package domain

// создание адаптера на стороне сервера
type CustomerRepositoryStub struct {
	customers []Customer
}

//присвоение адаптера методу, обозначенному портом
func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

// создание вспомогательной функции для mock data
func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"1001", "Pavel", "Minsk", "220025", "1991-12-04", "1"},
		{"2001", "Helen", "Lida", "220025", "1993-10-31", "1"},
	}
	return CustomerRepositoryStub{
		customers,
	}
}
