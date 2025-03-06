package entity

import (
	"database/sql/driver"
	"strings"
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	TransactionId   int            `json:"transaction_id" gorm:"primaryKey"`
	AccountId       int            `json:"account_id" gorm:"index"`
	OperationTypeId int            `json:"operation_type_id" gorm:"index"`
	Amount          float64        `json:"amount"`
	EventDate       CustomTime     `json:"event_date" gorm:"type:timestamptz"`
	Account         Account        `gorm:"foreignKey:AccountId;references:AccountId" json:"-"`
	OperationsType  OperationsType `gorm:"foreignKey:OperationTypeId;references:OperationTypeId" json:"-"`
}

func (t *Transaction) BeforeSave(tx *gorm.DB) (err error) {
	t.EventDate.Time = t.EventDate.UTC().Truncate(time.Microsecond)
	return nil
}

type CustomTime struct {
	time.Time
}

// UnmarshalJSON: Converte string do JSON para CustomTime
func (ct *CustomTime) UnmarshalJSON(b []byte) error {
	str := strings.Trim(string(b), "\"")
	if str == "" {
		return nil
	}

	// Primeiro tenta o formato vindo do JSON
	t, err := time.Parse("2006-01-02T15:04:05.9999999", str)
	if err != nil {
		// Tenta um formato alternativo com timezone
		t, err = time.Parse("2006-01-02T15:04:05.9999999Z07:00", str)
		if err != nil {
			return err
		}
	}

	ct.Time = t.UTC().Truncate(time.Microsecond) // Garante que está em UTC e com precisão correta
	return nil
}

// Value: Converte CustomTime para um formato aceito pelo PostgreSQL
func (ct CustomTime) Value() (driver.Value, error) {
	return ct.UTC().Format("2006-01-02 15:04:05.999999"), nil
}
