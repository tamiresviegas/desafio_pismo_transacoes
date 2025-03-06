package port

import "github.com/tamiresviegas/desafio_pismo_transacoes/internal/core/entity"

type TransactionRepository interface {
	CreateTransaction(transaction entity.Transaction) (entity.Transaction, error)
	GetTransactionByID(transactionID int) (entity.Transaction, error)
	GetAllTransaction() ([]entity.Transaction, error)
	UpdateTransaction(transaction entity.Transaction) (entity.Transaction, error)
	DeleteTransaction(transactionID int) error
}
