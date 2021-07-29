package app

import (
	"banking/errs"
	//"banking/domain"
	//"banking/domain"
	"banking/service"
	"encoding/json"
	"github.com/gorilla/mux"
	//"fmt"
	//"github.com/gorilla/mux"
	"net/http"
)

// имплементирует зависимость от CustomerService
type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers, err := ch.service.GetAllCustomer()

	if err != nil {
		writeResponse(w, err.Code, errs.NewUnexpectedError("Server error"))
	} else {
		writeResponse(w, http.StatusOK, customers)
	}
}

func (ch *CustomerHandlers) getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["customer_id"]

	customer, err := ch.service.GetCustomer(id)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, customer)
	}
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(data)
	//////// отличается от кода в видеоуроке
	if err != nil {
		panic(err)
	}
	/////////
}
