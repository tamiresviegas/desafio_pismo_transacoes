package repository

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/tamiresviegas/desafio_pismo_transacoes/config"
	"github.com/tamiresviegas/desafio_pismo_transacoes/internal/adapters/repository"
	"github.com/tamiresviegas/desafio_pismo_transacoes/internal/core/entity"
)

func TestCreateTransaction_Repo(t *testing.T) {
	db, err := config.ConnectBD()
	if err != nil {
		return
	}
	repo := repository.NewTransactionRepository(db.DB)

	eventDateStr := "2020-01-05T09:34:18.5893223"
	eventDate, err := time.Parse("2006-01-02T15:04:05.9999999", eventDateStr)
	if err != nil {
		t.Fatalf("Erro ao converter string para time: %v", err)
	}
	customEventDate := entity.CustomTime{Time: eventDate}
	transaction := entity.Transaction{TransactionId: 10013, AccountId: 1, OperationTypeId: 1, Amount: 10.0, EventDate: customEventDate}
	createdTransaction, err := repo.CreateTransaction(transaction)
	defer repo.DeleteTransaction(createdTransaction.TransactionId)

	expected := time.Date(2020, time.January, 5, 9, 34, 18, 589322300, time.UTC).Truncate(time.Microsecond)
	actual := createdTransaction.EventDate.Time.Truncate(time.Microsecond)

	assert.NoError(t, err)
	assert.NotZero(t, createdTransaction.TransactionId)
	assert.Equal(t, 10013, createdTransaction.TransactionId)
	assert.Equal(t, 1, createdTransaction.AccountId)
	assert.Equal(t, 1, createdTransaction.OperationTypeId)
	assert.Equal(t, 10.0, createdTransaction.Amount)
	assert.Equal(t, expected, actual)
	err = repo.DeleteTransaction(createdTransaction.TransactionId)
	if err != nil {
		return
	}
}

func TestGetTransactionByID_Repo(t *testing.T) {
	db, err := config.ConnectBD()
	if err != nil {
		return
	}
	repo := repository.NewTransactionRepository(db.DB)

	eventDateStr := "2020-01-05T09:34:18.5893223"
	eventDate, err := time.Parse("2006-01-02T15:04:05.9999999", eventDateStr)
	if err != nil {
		t.Fatalf("Erro ao converter string para time: %v", err)
	}
	customEventDate := entity.CustomTime{Time: eventDate}
	transaction := entity.Transaction{TransactionId: 1001, AccountId: 1, OperationTypeId: 1, Amount: 10.0, EventDate: customEventDate}
	createdTransaction, err := repo.CreateTransaction(transaction)
	if err != nil {
		return
	}

	foundTransaction, err := repo.GetTransactionByID(createdTransaction.TransactionId)

	assert.NoError(t, err)
	assert.Equal(t, createdTransaction.TransactionId, foundTransaction.TransactionId)
	err = repo.DeleteTransaction(createdTransaction.TransactionId)
	if err != nil {
		return
	}
}

func TestGetAllTransaction_Repo(t *testing.T) {
	db, err := config.ConnectBD()
	if err != nil {
		t.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}
	defer db.Close()

	repo := repository.NewTransactionRepository(db.DB)

	eventDateStr := "2020-01-05T09:34:18.5893223"
	eventDate, err := time.Parse("2006-01-02T15:04:05.9999999", eventDateStr)
	if err != nil {
		t.Fatalf("Erro ao converter string para time: %v", err)
	}
	customEventDate := entity.CustomTime{Time: eventDate}
	transaction1 := entity.Transaction{TransactionId: 1003, AccountId: 1, OperationTypeId: 1, Amount: 10.0, EventDate: customEventDate}
	transaction2 := entity.Transaction{TransactionId: 1004, AccountId: 1, OperationTypeId: 1, Amount: 10.0, EventDate: customEventDate}

	result1, err := repo.CreateTransaction(transaction1)
	if err != nil {
		t.Fatalf("Erro ao criar transaction1: %v", err)
	}

	result2, err := repo.CreateTransaction(transaction2)
	if err != nil {
		t.Fatalf("Erro ao criar transaction2: %v", err)
	}

	transactions, err := repo.GetAllTransaction()
	if err != nil {
		t.Fatalf("Erro ao obter todas as contas: %v", err)
	}

	var foundAccount1, foundAccount2 bool
	for _, transaction := range transactions {
		if transaction.TransactionId == 1003 {
			foundAccount1 = true
		}
		if transaction.TransactionId == 1004 {
			foundAccount2 = true
		}
	}

	assert.True(t, foundAccount1)
	assert.True(t, foundAccount2)
	err = repo.DeleteTransaction(result1.TransactionId)
	if err != nil {
		return
	}
	err = repo.DeleteTransaction(result2.TransactionId)
	if err != nil {
		return
	}
}

func TestUpdateTransaction_Repo(t *testing.T) {
	db, err := config.ConnectBD()
	if err != nil {
		t.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}
	defer db.Close()

	repo := repository.NewTransactionRepository(db.DB)

	eventDateStr := "2020-01-05T09:34:18.5893223"
	eventDate, err := time.Parse("2006-01-02T15:04:05.9999999", eventDateStr)
	if err != nil {
		t.Fatalf("Erro ao converter string para time: %v", err)
	}
	customEventDate := entity.CustomTime{Time: eventDate}
	transaction := entity.Transaction{TransactionId: 1005, AccountId: 1, OperationTypeId: 1, Amount: 10.0, EventDate: customEventDate}
	createdTransaction, err := repo.CreateTransaction(transaction)
	if err != nil {
		return
	}

	createdTransaction.TransactionId = 1005
	updatedTRans, err := repo.UpdateTransaction(createdTransaction)

	assert.NoError(t, err)
	assert.Equal(t, 1005, updatedTRans.TransactionId)
	assert.Equal(t, createdTransaction.TransactionId, updatedTRans.TransactionId)
	err = repo.DeleteTransaction(createdTransaction.AccountId)
	if err != nil {
		return
	}
}

func TestDeleteTransaction_Repo(t *testing.T) {
	db, err := config.ConnectBD()
	if err != nil {
		t.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}
	defer db.Close()

	repo := repository.NewTransactionRepository(db.DB)

	eventDateStr := "2020-01-05T09:34:18.5893223"
	eventDate, err := time.Parse("2006-01-02T15:04:05.9999999", eventDateStr)
	if err != nil {
		t.Fatalf("Erro ao converter string para time: %v", err)
	}
	customEventDate := entity.CustomTime{Time: eventDate}
	transaction := entity.Transaction{TransactionId: 1006, AccountId: 1, OperationTypeId: 1, Amount: 10.0, EventDate: customEventDate}
	createdTransaction, err := repo.CreateTransaction(transaction)
	if err != nil {
		return
	}

	err = repo.DeleteTransaction(createdTransaction.TransactionId)
	if err != nil {
		return
	}
	assert.NoError(t, err)

}
