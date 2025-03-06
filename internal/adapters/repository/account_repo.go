package repository

import (
	"github.com/tamiresviegas/desafio_pismo_transacoes/internal/core/entity"
	"github.com/tamiresviegas/desafio_pismo_transacoes/internal/core/port"
	"gorm.io/gorm"
)

type AccountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) port.AccountRepository {
	return &AccountRepository{db: db}
}

func (r *AccountRepository) CreateAccount(account entity.Account) (entity.Account, error) {
	if err := r.db.Create(&account).Error; err != nil {
		return entity.Account{}, err
	}
	return account, nil
}

func (r *AccountRepository) GetAccountByID(accountID int) (entity.Account, error) {
	var account entity.Account
	if err := r.db.Where("account_id = ?", accountID).First(&account).Error; err != nil {
		return entity.Account{}, err
	}
	return account, nil
}

func (r *AccountRepository) GetAllAccount() ([]entity.Account, error) {
	var account []entity.Account
	if err := r.db.Find(&account).Error; err != nil {
		return []entity.Account{}, err
	}
	return account, nil
}

func (r *AccountRepository) UpdateAccount(account entity.Account) (entity.Account, error) {
	if err := r.db.Save(&account).Error; err != nil {
		return entity.Account{}, err
	}
	return account, nil
}

func (r *AccountRepository) DeleteAccount(accountID int) error {
	var account entity.Account
	if err := r.db.Delete(&account, accountID).Error; err != nil {
		return err
	}
	return nil
}
