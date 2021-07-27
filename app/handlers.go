package app

import (
	//"banking/domain"
	//"banking/domain"
	"banking/service"
	"encoding/json"
	"encoding/xml"
	//"fmt"
	//"github.com/gorilla/mux"
	"net/http"
)

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {

	//customers := []domain.Customer{
	//	{"1","Pavel", "Minsk", "220025"},
	//	{"2", "Helen", "Lida", "123456"},
	//}

	customers, _ := ch.service.GetAllCustomer()

	// формат json или xml в зависимости от запроса
	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}
