package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

type Customer struct {
	Customer_ID string `json:"customer_id"`
	Name string `json:"full_name" xml:"name"`
	City string `json:"city" xml:"city"`
	Zipcode string `json:"zip_code" xml:"zipcode"`
}

func greet (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}

func getAllCustomers (w http.ResponseWriter, r *http.Request) {

	customers := []Customer{
		{"1","Pavel", "Minsk", "220025"},
		{"2", "Helen", "Lida", "123456"},
	}

	// формат json или xml в зависимости от запроса
	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}