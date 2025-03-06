package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/tamiresviegas/desafio_pismo_transacoes/internal/core/entity"
	"github.com/tamiresviegas/desafio_pismo_transacoes/internal/core/service"
)

type MockOperationTypesRepo struct {
	mock.Mock
}

func (m *MockOperationTypesRepo) DeleteOperationTypes(operationTypesID int) error {
	args := m.Called(operationTypesID)
	return args.Error(0)
}

func (m *MockOperationTypesRepo) GetAllOperationTypes() ([]entity.OperationsType, error) {
	args := m.Called()
	return args.Get(0).([]entity.OperationsType), args.Error(1)
}

func (m *MockOperationTypesRepo) UpdateOperationTypes(operationType entity.OperationsType) (entity.OperationsType, error) {
	args := m.Called(operationType)
	return args.Get(0).(entity.OperationsType), args.Error(1)
}

func (m *MockOperationTypesRepo) CreateOperationTypes(operationType entity.OperationsType) (entity.OperationsType, error) {
	args := m.Called(operationType)
	return args.Get(0).(entity.OperationsType), args.Error(1)
}

func (m *MockOperationTypesRepo) GetOperationTypesByID(operationTypesID int) (entity.OperationsType, error) {
	args := m.Called(operationTypesID)
	return args.Get(0).(entity.OperationsType), args.Error(1)
}

func TestCreateOperationType(t *testing.T) {
	mockRepo := new(MockOperationTypesRepo)

	service := service.NewOperationTypesService(mockRepo)

	opType := entity.OperationsType{OperationTypeId: 100002, Description0: "Normal Purchase test"}

	mockRepo.On("CreateOperationTypes", mock.Anything).Return(opType, nil)

	createdopt, err := service.CreateOperationTypes(opType)

	assert.NoError(t, err)
	assert.Equal(t, "Normal Purchase test", createdopt.Description0)
	mockRepo.AssertExpectations(t)
}

func TestGetOperationTypeByID(t *testing.T) {
	mockRepo := new(MockOperationTypesRepo)
	service := service.NewOperationTypesService(mockRepo)

	opType := entity.OperationsType{OperationTypeId: 100002, Description0: "Normal Purchase test"}

	mockRepo.On("GetOperationTypesByID", 1).Return(opType, nil)

	foundopt, err := service.GetOperationTypesByID(1)

	assert.NoError(t, err)
	assert.Equal(t, 100002, foundopt.OperationTypeId)
	assert.Equal(t, "Normal Purchase test", foundopt.Description0)
	mockRepo.AssertExpectations(t)
}

func TestGetAllOperationType(t *testing.T) {
	mockRepo := new(MockOperationTypesRepo)
	service := service.NewOperationTypesService(mockRepo)

	opts := []entity.OperationsType{
		{OperationTypeId: 10002, Description0: "Normal Purchase test"},
		{OperationTypeId: 20002, Description0: "Normal Purchase test 2"},
	}

	mockRepo.On("GetAllOperationTypes").Return(opts, nil)

	foundAccounts, err := service.GetAllOperationTypes()

	assert.NoError(t, err)
	assert.Len(t, foundAccounts, 2)
	assert.Equal(t, 10002, foundAccounts[0].OperationTypeId)
	assert.Equal(t, "Normal Purchase test", foundAccounts[0].Description0)
	assert.Equal(t, 20002, foundAccounts[1].OperationTypeId)
	assert.Equal(t, "Normal Purchase test 2", foundAccounts[1].Description0)
	mockRepo.AssertExpectations(t)
}

func TestUpdateOperationType(t *testing.T) {
	mockRepo := new(MockOperationTypesRepo)
	service := service.NewOperationTypesService(mockRepo)

	optInput := entity.OperationsType{OperationTypeId: 10002, Description0: "Normal Purchase test"}
	updatedOpt := entity.OperationsType{OperationTypeId: 10003, Description0: "Normal Purchase test 2"}

	mockRepo.On("UpdateOperationTypes", optInput).Return(updatedOpt, nil)

	resultOpt, err := service.UpdateOperationTypes(optInput)

	assert.NoError(t, err)
	assert.Equal(t, 10003, resultOpt.OperationTypeId)
	assert.Equal(t, "Normal Purchase test 2", resultOpt.Description0)
	mockRepo.AssertExpectations(t)
}

func TestDeleteOperationType(t *testing.T) {
	mockRepo := new(MockOperationTypesRepo)
	service := service.NewOperationTypesService(mockRepo)

	operationTypesID := 1003

	mockRepo.On("DeleteOperationTypes", operationTypesID).Return(nil)

	err := service.DeleteOperationTypes(operationTypesID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
