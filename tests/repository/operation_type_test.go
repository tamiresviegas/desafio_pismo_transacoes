package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tamiresviegas/desafio_pismo_transacoes/config"
	"github.com/tamiresviegas/desafio_pismo_transacoes/internal/adapters/repository"
	"github.com/tamiresviegas/desafio_pismo_transacoes/internal/core/entity"
)

func TestCreateOperationType_Repo(t *testing.T) {
	db, err := config.ConnectBD()
	if err != nil {
		return
	}
	repo := repository.NewOperationTypesRepository(db.DB)

	opType := entity.OperationsType{OperationTypeId: 1, Description0: "Normal Purchase"}
	createdOpType, err := repo.CreateOperationTypes(opType)
	defer repo.DeleteOperationTypes(createdOpType.OperationTypeId)

	assert.NoError(t, err)
	assert.NotZero(t, createdOpType.OperationTypeId)
	assert.Equal(t, 1, createdOpType.OperationTypeId)
	err = repo.DeleteOperationTypes(createdOpType.OperationTypeId)
	if err != nil {
		return
	}
}

func TestGetOperationTypeByID_Repo(t *testing.T) {
	db, err := config.ConnectBD()
	if err != nil {
		return
	}
	repo := repository.NewOperationTypesRepository(db.DB)

	opType := entity.OperationsType{OperationTypeId: 1, Description0: "Normal Purchase test"}
	createdOpType, _ := repo.CreateOperationTypes(opType)

	foundAccount, err := repo.GetOperationTypesByID(createdOpType.OperationTypeId)

	assert.NoError(t, err)
	assert.Equal(t, createdOpType.OperationTypeId, foundAccount.OperationTypeId)
	assert.Equal(t, createdOpType.Description0, foundAccount.Description0)
	err = repo.DeleteOperationTypes(createdOpType.OperationTypeId)
	if err != nil {
		return
	}
}

func TestGetAllOperationType_Repo(t *testing.T) {
	db, err := config.ConnectBD()
	if err != nil {
		t.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}
	defer db.Close()

	repo := repository.NewOperationTypesRepository(db.DB)

	opType1 := entity.OperationsType{OperationTypeId: 10000, Description0: "Normal Purchase test"}
	opType2 := entity.OperationsType{OperationTypeId: 10001, Description0: "Purchase with installments test"}

	result1, err := repo.CreateOperationTypes(opType1)
	if err != nil {
		t.Fatalf("Erro ao criar account1: %v", err)
	}

	result2, err := repo.CreateOperationTypes(opType2)
	if err != nil {
		t.Fatalf("Erro ao criar account2: %v", err)
	}

	optype, err := repo.GetAllOperationTypes()
	if err != nil {
		t.Fatalf("Erro ao obter todas as contas: %v", err)
	}

	var foundAccount1, foundAccount2 bool
	for _, op := range optype {
		if op.Description0 == "Normal Purchase test" {
			foundAccount1 = true
		}
		if op.Description0 == "Purchase with installments test" {
			foundAccount2 = true
		}
	}

	assert.True(t, foundAccount1)
	assert.True(t, foundAccount2)
	err = repo.DeleteOperationTypes(result1.OperationTypeId)
	if err != nil {
		return
	}
	err = repo.DeleteOperationTypes(result2.OperationTypeId)
	if err != nil {
		return
	}
}

func TestUpdateOperationType_Repo(t *testing.T) {
	db, err := config.ConnectBD()
	if err != nil {
		t.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}
	defer db.Close()

	repo := repository.NewOperationTypesRepository(db.DB)

	opType := entity.OperationsType{OperationTypeId: 100002, Description0: "Normal Purchase test"}
	createdOpType, err := repo.CreateOperationTypes(opType)
	if err != nil {
		t.Fatalf("Erro ao criar a conta: %v", err)
	}

	createdOpType.Description0 = "Normal Purchase test"
	updatedopt, err := repo.UpdateOperationTypes(createdOpType)

	assert.NoError(t, err)
	assert.Equal(t, "Normal Purchase test", updatedopt.Description0)
	assert.Equal(t, createdOpType.OperationTypeId, updatedopt.OperationTypeId)
	err = repo.DeleteOperationTypes(createdOpType.OperationTypeId)
	if err != nil {
		return
	}
}

func TestDeleteOperationType_Repo(t *testing.T) {
	db, err := config.ConnectBD()
	if err != nil {
		t.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}
	defer db.Close()

	repo := repository.NewOperationTypesRepository(db.DB)

	opType := entity.OperationsType{OperationTypeId: 100002, Description0: "Normal Purchase test"}
	createdOpType, err := repo.CreateOperationTypes(opType)
	if err != nil {
		t.Fatalf("Erro ao criar a conta: %v", err)
	}

	err = repo.DeleteOperationTypes(createdOpType.OperationTypeId)
	if err != nil {
		return
	}
	assert.NoError(t, err)

}
