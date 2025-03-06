package port

import "github.com/tamiresviegas/desafio_pismo_transacoes/internal/core/entity"

type AccountRepository interface {
	CreateAccount(account entity.Account) (entity.Account, error)
	GetAccountByID(accountID int) (entity.Account, error)
	GetAllAccount() ([]entity.Account, error)
	UpdateAccount(account entity.Account) (entity.Account, error)
	DeleteAccount(accountID int) error
}
