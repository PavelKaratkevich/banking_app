package service

import (
	"banking/domain"
	"banking/dto"
	"banking/errs"
	"strings"
	"time"
)

// AccountService это Primary port
type AccountService interface {
	NewAccount(dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppErr)
	MakeTransaction(req dto.NewTransactionRequest) (*dto.NewTransactionResponse, *errs.AppErr)
}

// привязываем Primary Port к domain, т.е. это связь домена с внешним миром
type DefaultAccountService struct {
	repo domain.AccountRepository
}

func (s DefaultAccountService) NewAccount(req dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppErr) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	a := domain.Account{
		AccountID:   "",
		CustomerId:  req.CusomerId,
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
		AccountType: req.AccountType,
		Amount:      req.Amount,
		Status:      "1",
	}
	if newAccount, err := s.repo.Save(a); err != nil {
		return nil, err
	} else {
		response := newAccount.ToNewAccountResponseDto()
		return &response, nil
	}
}

func (s DefaultAccountService) MakeTransaction(req dto.NewTransactionRequest) (*dto.NewTransactionResponse, *errs.AppErr) {
	err := req.Validate()

	if req.IsTransactionTypeWithdrawal() {
		account, err := s.repo.FindBy(req.AccountId)
		if err != nil {
			return nil, err
		}
		if !account.CanWithdraw(req.Amount) {
			return nil, errs.NewValidationError("Insufficient balance in the account")
		}
	}

	if err != nil {
		return nil, err
	}
	a := domain.Transaction{
		TransactionId: 0,
		AccountId:     req.AccountId,
		Amount:        req.Amount,
		Type:          strings.ToLower(req.Type),
		Date:          time.Now().Format("2006-01-02 15:04:05"),
	}
	newTransaction, err := s.repo.Update(a)
	if err != nil {
		return nil, err
	}
	response := newTransaction.ToNewTransactionResponseDto()
	return &response, nil
}

// привязываем Account Repository к домену
func NewAccountService(repo domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{repo}
}
