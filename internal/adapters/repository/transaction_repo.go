package repository

import (
	"github.com/tamiresviegas/desafio_pismo_transacoes/internal/core/entity"
	"github.com/tamiresviegas/desafio_pismo_transacoes/internal/core/port"
	"gorm.io/gorm"
)

type TransactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) port.TransactionRepository {
	return &TransactionRepository{db: db}
}

func (r *TransactionRepository) CreateTransaction(transaction entity.Transaction) (entity.Transaction, error) {

	if err := r.db.Create(&transaction).Error; err != nil {
		return entity.Transaction{}, err
	}
	return transaction, nil
}

func (r *TransactionRepository) GetTransactionByID(transactionId int) (entity.Transaction, error) {
	var transaction entity.Transaction
	if err := r.db.Where("account_id = ?", transactionId).First(&transaction).Error; err != nil {
		return entity.Transaction{}, err
	}
	return transaction, nil
}

func (r *TransactionRepository) GetAllTransaction() ([]entity.Transaction, error) {
	var transaction []entity.Transaction
	if err := r.db.Find(&transaction).Error; err != nil {
		return []entity.Transaction{}, err
	}
	return transaction, nil
}

func (r *TransactionRepository) UpdateTransaction(transaction entity.Transaction) (entity.Transaction, error) {
	if err := r.db.Save(&transaction).Error; err != nil {
		return entity.Transaction{}, err
	}
	return transaction, nil
}

func (r *TransactionRepository) DeleteTransaction(transactionId int) error {
	var transaction entity.Transaction
	if err := r.db.Delete(&transaction, transactionId).Error; err != nil {
		return err
	}
	return nil
}
