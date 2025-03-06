package entity

import (
	"gopkg.in/validator.v2"
)

type Account struct {
	AccountId      int    `json:"account_id" gorm:"primaryKey"`
	DocumentNumber string `json:"document_number" validate:"nonzero, regexp=^[0-9]*$"`
}

func ValidAccount(account *Account) error {

	if err := validator.Validate(account); err != nil {
		return err
	}

	return nil
}
