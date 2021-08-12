package app

import (
	"banking/dto"
	"banking/errs"
	"banking/mocks/service"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"testing"
)

var router *mux.Router
var ch CustomerHandlers
var mockService *service.MockCustomerService

func setup(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockService = service.NewMockCustomerService(ctrl)
	ch = CustomerHandlers{mockService}
	router = mux.NewRouter()
}

// Test_should_return_customers_with_status_code_200 оставил без рефакторинга для наглядности комментариев
func Test_should_return_customers_with_status_code_200(t *testing.T) {
	// Arrange

	// создаю контроллер
	ctrl := gomock.NewController(t)
	mockService := service.NewMockCustomerService(ctrl)
	ch := CustomerHandlers{mockService}
	/*
		ch - это CustomerHandlers (имплементация PrimaryPort) в который вместо обычного CustomerService
		внедрен MockCustomerService который, в свою очередь, инициализируется функцией NewMockCustomerService(), которая
		возвращает MockCustomerService
	*/

	// создаю выдуманных пользователей
	dummyCustomers := []dto.CustomerResponse{
		{"1001", "Pavel", "Minsk", "220025", "1991-12-04", "1"},
		{"2001", "Helen", "Lida", "220025", "1993-10-31", "1"},
	}

	// "" - это про отсутствие статуса (valid/invalid)
	mockService.EXPECT().GetAllCustomer("").Return(dummyCustomers, nil)

	// создаю роутер и прописываю endpoint как и в обычном случае
	router := mux.NewRouter()
	router.HandleFunc("/customers", ch.getAllCustomers)

	// создаем http-запрос
	request, _ := http.NewRequest(http.MethodGet, "/customers", nil)

	// Act

	// NewRecorder is the implementation of http.ResponseWriter
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	//Assert

	if recorder.Code != http.StatusOK {
		t.Error("Failed while testing the status code")
	}
}

func Test_should_return_status_code_500_with_error_message(t *testing.T) {

	setup(t)
	router.HandleFunc("/customers", ch.getAllCustomers)
	mockService.EXPECT().GetAllCustomer("").Return(nil, errs.NewUnexpectedError("some database error"))
	request, _ := http.NewRequest(http.MethodGet, "/customers", nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	if recorder.Code != http.StatusInternalServerError {
		t.Error("")
	}
}

func Test_should_return_customer_with_status_code_200(t *testing.T) {
	setup(t)

	dummyCustomer := dto.CustomerResponse{"1001", "Pavel", "Minsk", "220025", "1991-12-04", "1"}

	router.HandleFunc("/customers/1001", ch.getCustomer)

	request, _ := http.NewRequest(http.MethodGet, "/customers/1001", nil)
	mockService.EXPECT().GetCustomer("").Return(&dummyCustomer, nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Error("Failed while testing the status code")
	}
}

func Test_should_return_customer_with_status_code_500(t *testing.T) {
	setup(t)

	router.HandleFunc("/customers/1001", ch.getCustomer)

	request, _ := http.NewRequest(http.MethodGet, "/customers/1001", nil)
	mockService.EXPECT().GetCustomer("").Return(nil, errs.NewUnexpectedError("Some database error"))
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusInternalServerError {
		t.Error("Failed while testing the status code")
	}
}
