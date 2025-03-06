package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/tamiresviegas/desafio_pismo_transacoes/internal/core/entity"
	"github.com/tamiresviegas/desafio_pismo_transacoes/internal/core/service"
)

type MockAccountRepo struct {
	mock.Mock
}

func (m *MockAccountRepo) DeleteAccount(accountID int) error {
	args := m.Called(accountID)
	return args.Error(0)
}

func (m *MockAccountRepo) GetAllAccount() ([]entity.Account, error) {
	args := m.Called()
	return args.Get(0).([]entity.Account), args.Error(1)
}

func (m *MockAccountRepo) UpdateAccount(account entity.Account) (entity.Account, error) {
	args := m.Called(account)
	return args.Get(0).(entity.Account), args.Error(1)
}

func (m *MockAccountRepo) CreateAccount(account entity.Account) (entity.Account, error) {
	args := m.Called(account)
	return args.Get(0).(entity.Account), args.Error(1)
}

func (m *MockAccountRepo) GetAccountByID(accountID int) (entity.Account, error) {
	args := m.Called(accountID)
	return args.Get(0).(entity.Account), args.Error(1)
}

func TestCreateAccount(t *testing.T) {
	mockRepo := new(MockAccountRepo)

	service := service.NewAccountService(mockRepo)

	account := entity.Account{AccountId: 1, DocumentNumber: "12345678900"}

	mockRepo.On("CreateAccount", mock.Anything).Return(account, nil)

	createdAccount, err := service.CreateAccount(account)

	assert.NoError(t, err)
	assert.Equal(t, "12345678900", createdAccount.DocumentNumber)
	mockRepo.AssertExpectations(t)
}

func TestGetAccountByID(t *testing.T) {
	mockRepo := new(MockAccountRepo)
	service := service.NewAccountService(mockRepo)

	account := entity.Account{AccountId: 1, DocumentNumber: "12345678900"}

	mockRepo.On("GetAccountByID", 1).Return(account, nil)

	foundAccount, err := service.GetAccountByID(1)

	assert.NoError(t, err)
	assert.Equal(t, 1, foundAccount.AccountId)
	assert.Equal(t, "12345678900", foundAccount.DocumentNumber)
	mockRepo.AssertExpectations(t)
}

func TestGetAllAccount(t *testing.T) {
	mockRepo := new(MockAccountRepo)
	service := service.NewAccountService(mockRepo)

	accounts := []entity.Account{
		{AccountId: 1, DocumentNumber: "12345678900"},
		{AccountId: 2, DocumentNumber: "98765432100"},
	}

	mockRepo.On("GetAllAccount").Return(accounts, nil)

	foundAccounts, err := service.GetAllAccount()

	assert.NoError(t, err)
	assert.Len(t, foundAccounts, 2)
	assert.Equal(t, 1, foundAccounts[0].AccountId)
	assert.Equal(t, "12345678900", foundAccounts[0].DocumentNumber)
	assert.Equal(t, 2, foundAccounts[1].AccountId)
	assert.Equal(t, "98765432100", foundAccounts[1].DocumentNumber)
	mockRepo.AssertExpectations(t)
}

func TestUpdateAccount(t *testing.T) {
	mockRepo := new(MockAccountRepo)
	service := service.NewAccountService(mockRepo)

	accountInput := entity.Account{AccountId: 1, DocumentNumber: "12345678900"}
	updatedAccount := entity.Account{AccountId: 1, DocumentNumber: "98765432100"}

	mockRepo.On("UpdateAccount", accountInput).Return(updatedAccount, nil)

	resultAccount, err := service.UpdateAccount(accountInput)

	assert.NoError(t, err)
	assert.Equal(t, 1, resultAccount.AccountId)
	assert.Equal(t, "98765432100", resultAccount.DocumentNumber)
	mockRepo.AssertExpectations(t)
}

func TestDeleteAccount(t *testing.T) {
	mockRepo := new(MockAccountRepo)
	service := service.NewAccountService(mockRepo)

	accountID := 1

	mockRepo.On("DeleteAccount", accountID).Return(nil)

	err := service.DeleteAccount(accountID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}


