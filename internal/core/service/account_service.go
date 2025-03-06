package service

import (
	"github.com/tamiresviegas/desafio_pismo_transacoes/internal/core/entity"
	"github.com/tamiresviegas/desafio_pismo_transacoes/internal/core/port"
)

type AccountService struct {
	repo port.AccountRepository
}

func NewAccountService(repo port.AccountRepository) *AccountService {
	return &AccountService{repo: repo}
}

func (s *AccountService) CreateAccount(accountInput entity.Account) (entity.Account, error) {

	account := entity.Account{
		AccountId:      accountInput.AccountId,
		DocumentNumber: accountInput.DocumentNumber,
	}

	return s.repo.CreateAccount(account)
}

func (s *AccountService) GetAccountByID(accountID int) (entity.Account, error) {
	return s.repo.GetAccountByID(accountID)
}

func (s *AccountService) GetAllAccount() ([]entity.Account, error) {
	return s.repo.GetAllAccount()
}

func (s *AccountService) UpdateAccount(accountInput entity.Account) (entity.Account, error) {

	account := entity.Account{
		AccountId:      accountInput.AccountId,
		DocumentNumber: accountInput.DocumentNumber,
	}

	return s.repo.UpdateAccount(account)
}

func (s *AccountService) DeleteAccount(accountID int) error {
	return s.repo.DeleteAccount(accountID)
}
