package service

import (
	"fmt"

	"github.com/tamiresviegas/desafio_pismo_transacoes/internal/core/entity"
	"github.com/tamiresviegas/desafio_pismo_transacoes/internal/core/port"
)

type TransactionService struct {
	repo        port.TransactionRepository
	accountRepo port.AccountRepository
	opRepo      port.OperationTypesRepository
}

func NewTransactionService(repo port.TransactionRepository, accountRepo port.AccountRepository, opRepo port.OperationTypesRepository) *TransactionService {
	return &TransactionService{repo: repo, accountRepo: accountRepo, opRepo: opRepo}
}

func (s *TransactionService) CreateTransaction(transactiontInput entity.Transaction) (entity.Transaction, error) {

	_, err := s.accountRepo.GetAccountByID(transactiontInput.AccountId)
	if err != nil {
		return entity.Transaction{}, fmt.Errorf("invalid account_id: %d", transactiontInput.AccountId)
	}

	_, err = s.opRepo.GetOperationTypesByID(transactiontInput.OperationTypeId)
	if err != nil {
		return entity.Transaction{}, fmt.Errorf("invalid operation_type: %d", transactiontInput.AccountId)
	}

	transaction := entity.Transaction{
		TransactionId:   transactiontInput.TransactionId,
		AccountId:       transactiontInput.AccountId,
		OperationTypeId: transactiontInput.OperationTypeId,
		Amount:          transactiontInput.Amount,
		EventDate:       transactiontInput.EventDate,
	}

	return s.repo.CreateTransaction(transaction)
}

func (s *TransactionService) GetTransactionByID(transactionID int) (entity.Transaction, error) {
	return s.repo.GetTransactionByID(transactionID)
}

func (s *TransactionService) GetAllTransaction() ([]entity.Transaction, error) {
	return s.repo.GetAllTransaction()
}

func (s *TransactionService) UpdateTransaction(transactiontInput entity.Transaction) (entity.Transaction, error) {

	transaction := entity.Transaction{
		TransactionId:   transactiontInput.TransactionId,
		AccountId:       transactiontInput.AccountId,
		OperationTypeId: transactiontInput.OperationTypeId,
		Amount:          transactiontInput.Amount,
		EventDate:       transactiontInput.EventDate,
	}

	return s.repo.UpdateTransaction(transaction)
}

func (s *TransactionService) DeleteTransaction(transactionID int) error {
	return s.repo.DeleteTransaction(transactionID)
}
