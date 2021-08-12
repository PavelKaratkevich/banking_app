package service

import (
	domain2 "banking/domain"
	"banking/dto"
	"banking/errs"
	"banking/mocks/domain"
	"github.com/golang/mock/gomock"
	"strings"
	"time"

	"testing"
)

var mockRepo *domain.MockAccountRepository
var ctrl *gomock.Controller
var service AccountService

func setup(t *testing.T) {
	ctrl = gomock.NewController(t)
	mockRepo = domain.NewMockAccountRepository(ctrl)
	service = NewAccountService(mockRepo)
}

// Тестируем функцию func (s DefaultAccountService) NewAccount(req dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppErr)
func Test_should_return_a_validation_error_response_when_the_request_is_not_valid(t *testing.T) {
	// Arrange
	request := dto.NewAccountRequest{
		CusomerId:   "100",
		AccountType: "saving",
		Amount:      -7000,
	}
	service := NewAccountService(nil)
	_, appErr := service.NewAccount(request)
	if appErr == nil {
		t.Error("failed while testing the new account validation")
	}
}

func Test_should_return_an_error_from_the_server_side_if_the_new_account_cannot_be_created(t *testing.T) {
	// Arrange
	setup(t)
	request := dto.NewAccountRequest{
		CusomerId:   "2000",
		AccountType: "saving",
		Amount:      6000,
	}
	account := domain2.Account{
		CustomerId:  request.CusomerId,
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
		AccountType: request.AccountType,
		Amount:      request.Amount,
		Status:      "1",
	}
	mockRepo.EXPECT().Save(account).Return(nil, errs.NewUnexpectedError("Error from a database"))
	// Act
	_, err := service.NewAccount(request)
	//Assert
	if err == nil {
		t.Error("Error while creating a new account")
	}
}

func Test_should_return_new_account_response_when_a_new_account_is_saved_successfully(t *testing.T) {
	// Arrange
	setup(t)
	request := dto.NewAccountRequest{
		CusomerId:   "2000",
		AccountType: "saving",
		Amount:      6000,
	}
	account := domain2.Account{
		CustomerId:  request.CusomerId,
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
		AccountType: request.AccountType,
		Amount:      request.Amount,
		Status:      "1",
	}

	accountWithId := account
	accountWithId.AccountID = "201"

	mockRepo.EXPECT().Save(account).Return(&accountWithId, nil)

	// Act
	newAccount, err := service.NewAccount(request)

	//Assert
	if err != nil {
		t.Error("Test failed while creating a new account")
	}
	if newAccount.AccountId != accountWithId.AccountID {
		t.Error("Test failed while matching new account ID")
	}
}

// Тестируем func(s DefaultAccountService) MakeTransaction(req dto.NewTransactionRequest) (*dto.NewTransactionResponse, *errs.AppErr)
func Test_if_validate_returns_an_error_while_passing_a_wrong_transaction_type(t *testing.T) {
	transaction := dto.NewTransactionRequest{
		Type:   "invalid_type",
		Amount: -2000,
	}
	err := transaction.Validate()
	if err.Message != "Transaction type should be 'withdrawal' or 'deposit'" && err.Message != "To make a transaction you need to deposit an amount equal to at least 1.0" {
		t.Error("Error while checking transaction type")
	}
}

func Test_should_return_error_when_FindBy_is_called(t *testing.T) {
	// Arragne
	setup(t)
	request := dto.NewTransactionRequest{
		AccountId:  2001,
		CustomerId: 95470,
		Type:       "withdrawal",
		Amount:     6000,
	}
	mockRepo.EXPECT().FindBy(request.AccountId).Return(nil, errs.NewUnexpectedError("Error while finding an account in a database"))

	// Act
	_, err := service.MakeTransaction(request)

	// Assert
	if err == nil {
		t.Error("Error while making a transaction")
	}
}

func Test_should_return_new_transaction_after_Update_is_called(t *testing.T) {
	// Arrange
	setup(t)
	req := dto.NewTransactionRequest{
		AccountId:  2001,
		CustomerId: 95470,
		Type:       "deposit",
		Amount:     6000,
	}
	transaction := domain2.Transaction{
		TransactionId: 0,
		AccountId:     req.AccountId,
		Amount:        req.Amount,
		Type:          strings.ToLower(req.Type),
		Date:          time.Now().Format("2006-01-02 15:04:05"),
	}
	transactionWithTransactionId := transaction
	transactionWithTransactionId.TransactionId = 201
	mockRepo.EXPECT().Update(transaction).Return(&transactionWithTransactionId, nil)

	// Act
	newTransaction, err := service.MakeTransaction(req)

	// Assert
	if err != nil {
		t.Error("Error while making transaction")
	}
	if newTransaction.TransactionId != transactionWithTransactionId.TransactionId {
		t.Error("Error with transaction ID")
	}
}

func Test_should_return_error_when_Update_is_called(t *testing.T) {
	// Arrange
	setup(t)
	req := dto.NewTransactionRequest{
		AccountId:  2001,
		CustomerId: 95470,
		Type:       "deposit",
		Amount:     6000,
	}
	transaction := domain2.Transaction{
		TransactionId: 0,
		AccountId:     req.AccountId,
		Amount:        req.Amount,
		Type:          strings.ToLower(req.Type),
		Date:          time.Now().Format("2006-01-02 15:04:05"),
	}
	transactionWithTransactionId := transaction
	transactionWithTransactionId.TransactionId = 201
	mockRepo.EXPECT().Update(transaction).Return(nil, errs.NewUnexpectedError("Error while updating transaction"))

	// Act
	_, err := service.MakeTransaction(req)

	// Assert
	if err == nil {
		t.Error("Test failed while updating transaction")
	}
}
