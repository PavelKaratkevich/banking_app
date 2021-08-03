package service

import (
	"banking/domain"
	"banking/dto"
	"banking/errs"
	"time"
)

// TransactionService это Primary port
type TransactionService interface {
	MakeTransaction(dto.NewTransactionRequest) (*dto.NewTransactionResponse, *errs.AppErr)
}

// привязываем Primary Port к domain, т.е. это связь домена с внешним миром
type DefaultTransactionService struct {
	repo domain.TransactionRepository
}

func (s DefaultTransactionService) MakeTransaction(req dto.NewTransactionRequest) (*dto.NewTransactionResponse, *errs.AppErr) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}
	a := domain.Transaction{
		TransactionId: 0,
		AccountId:     req.AccountId,
		Amount:        req.Amount,
		Type:          req.Type,
		Date:          time.Now().Format("2006-01-02 15:04:05"),
	}
	newTransaction, err := s.repo.Update(a)
	if err != nil {
		return nil, err
	}
	response := newTransaction.ToNewTransactionResponseDto()
	return &response, nil
}

// привязываем Transaction Repository к домену
func NewTransactionService(repo domain.TransactionRepository) DefaultTransactionService {
	return DefaultTransactionService{repo}
}
