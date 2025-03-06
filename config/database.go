package config

import (
	"log"

	"github.com/tamiresviegas/desafio_pismo_transacoes/internal/core/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
}

func (d *Database) Close() {
	sqlDB, err := d.DB.DB()
	if err != nil {
		log.Panic("Error getting SQL DB instance: ", err)
		return
	}
	err = sqlDB.Close()
	if err != nil {
		log.Panic("Error closing the DB connection: ", err)
	}
}

func ConnectBD() (*Database, error) {
	conn := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	DB, err := gorm.Open(postgres.Open(conn), &gorm.Config{})
	if err != nil {
		log.Panic("Error connect to DB")
	}
	DB.AutoMigrate(&entity.Account{})
	DB.AutoMigrate(&entity.OperationsType{})
	DB.AutoMigrate(&entity.Transaction{})

	//DB.Exec("ALTER TABLE transactions ADD CONSTRAINT fk_accounts_transactions FOREIGN KEY (account_id) REFERENCES accounts(account_id)")
	//DB.Exec("ALTER TABLE transactions ADD CONSTRAINT fk_operation_type_id_transactions FOREIGN KEY (operation_type_id) REFERENCES operations_types(operation_type_id)")

	return &Database{DB: DB}, nil
}
