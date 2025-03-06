package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tamiresviegas/desafio_pismo_transacoes/config"
	"github.com/tamiresviegas/desafio_pismo_transacoes/internal/adapters/repository"
	"github.com/tamiresviegas/desafio_pismo_transacoes/internal/core/entity"
)

func TestCreateAccount_Repo(t *testing.T) {
	db, err := config.ConnectBD()
	if err != nil {
		return
	}
	repo := repository.NewAccountRepository(db.DB)

	account := entity.Account{DocumentNumber: "12345678900"}
	createdAccount, err := repo.CreateAccount(account)
	defer repo.DeleteAccount(createdAccount.AccountId)

	assert.NoError(t, err)
	assert.NotZero(t, createdAccount.AccountId)
	assert.Equal(t, "12345678900", createdAccount.DocumentNumber)
	err = repo.DeleteAccount(createdAccount.AccountId)
	if err != nil {
		return
	}
}

func TestGetAccountByID_Repo(t *testing.T) {
	db, err := config.ConnectBD()
	if err != nil {
		return
	}
	repo := repository.NewAccountRepository(db.DB)

	account := entity.Account{DocumentNumber: "12345678900"}
	createdAccount, _ := repo.CreateAccount(account)

	foundAccount, err := repo.GetAccountByID(createdAccount.AccountId)

	assert.NoError(t, err)
	assert.Equal(t, createdAccount.AccountId, foundAccount.AccountId)
	assert.Equal(t, createdAccount.DocumentNumber, foundAccount.DocumentNumber)
	err = repo.DeleteAccount(createdAccount.AccountId)
	if err != nil {
		return
	}
}

func TestGetAllAccount_Repo(t *testing.T) {
	db, err := config.ConnectBD()
	if err != nil {
		t.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}
	defer db.Close()

	repo := repository.NewAccountRepository(db.DB)

	account1 := entity.Account{DocumentNumber: "12345678900"}
	account2 := entity.Account{DocumentNumber: "98765432100"}

	result1, err := repo.CreateAccount(account1)
	if err != nil {
		t.Fatalf("Erro ao criar account1: %v", err)
	}

	result2, err := repo.CreateAccount(account2)
	if err != nil {
		t.Fatalf("Erro ao criar account2: %v", err)
	}

	accounts, err := repo.GetAllAccount()
	if err != nil {
		t.Fatalf("Erro ao obter todas as contas: %v", err)
	}

	var foundAccount1, foundAccount2 bool
	for _, account := range accounts {
		if account.DocumentNumber == "12345678900" {
			foundAccount1 = true
		}
		if account.DocumentNumber == "98765432100" {
			foundAccount2 = true
		}
	}

	assert.True(t, foundAccount1)
	assert.True(t, foundAccount2)
	err = repo.DeleteAccount(result1.AccountId)
	if err != nil {
		return
	}
	err = repo.DeleteAccount(result2.AccountId)
	if err != nil {
		return
	}
}

func TestUpdateAccount_Repo(t *testing.T) {
	db, err := config.ConnectBD()
	if err != nil {
		t.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}
	defer db.Close()

	repo := repository.NewAccountRepository(db.DB)

	account := entity.Account{DocumentNumber: "12345678900"}
	createdAccount, err := repo.CreateAccount(account)
	if err != nil {
		t.Fatalf("Erro ao criar a conta: %v", err)
	}

	createdAccount.DocumentNumber = "11122334455"
	updatedAccount, err := repo.UpdateAccount(createdAccount)

	assert.NoError(t, err)
	assert.Equal(t, "11122334455", updatedAccount.DocumentNumber)
	assert.Equal(t, createdAccount.AccountId, updatedAccount.AccountId)
	err = repo.DeleteAccount(createdAccount.AccountId)
	if err != nil {
		return
	}
}

func TestDeleteAccount_Repo(t *testing.T) {
	db, err := config.ConnectBD()
	if err != nil {
		t.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}
	defer db.Close()

	repo := repository.NewAccountRepository(db.DB)
	account := entity.Account{DocumentNumber: "12345678900"}
	createdAccount, err := repo.CreateAccount(account)
	if err != nil {
		t.Fatalf("Erro ao criar a conta: %v", err)
	}

	err = repo.DeleteAccount(createdAccount.AccountId)
	if err != nil {
		return
	}
	assert.NoError(t, err)

}
