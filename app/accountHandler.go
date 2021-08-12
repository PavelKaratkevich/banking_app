package app

import (
	"banking/dto"
	"banking/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// AccountHandlers это адаптер к порту AccountService
type AccountHandler struct {
	service service.AccountService
}

func (h *AccountHandler) NewAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId := vars["customer_id"]

	// создание инстанса запроса DTO в адрес домена
	var request dto.NewAccountRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		request.CusomerId = customerId
		account, appErr := h.service.NewAccount(request)
		if appErr != nil {
			writeResponse(w, appErr.Code, appErr.Message)
		} else {
			writeResponse(w, http.StatusCreated, account)
		}
	}

}

func (h *AccountHandler) MakeTransaction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	accountId := vars["account_id"]
	customerId := vars["customer_id"]

	var request dto.NewTransactionRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		request.AccountId, _ = strconv.Atoi(accountId)
		request.CustomerId, _ = strconv.Atoi(customerId)

		transaction, appErr := h.service.MakeTransaction(request)
		if appErr != nil {
			writeResponse(w, appErr.Code, appErr.Message)
		} else {
			writeResponse(w, http.StatusCreated, transaction)
		}
	}
}
