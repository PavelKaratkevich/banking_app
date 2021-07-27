package app

import (
	"banking/domain"
	"banking/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Start() {

	router := mux.NewRouter()

	ch := CustomerHandlers{
		service.NewCustomerService(domain.NewCustomerRepositoryStub()),
	}

	// define routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)

	// start server
	log.Fatal(http.ListenAndServe(":8000", router))
}
