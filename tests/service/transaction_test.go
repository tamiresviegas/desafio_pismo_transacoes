package service

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/tamiresviegas/desafio_pismo_transacoes/internal/core/entity"
	"github.com/tamiresviegas/desafio_pismo_transacoes/internal/core/service"
)

type MockTransactionRepo struct {
	mock.Mock
}

func (m *MockTransactionRepo) DeleteTransaction(transactions int) error {
	args := m.Called(transactions)
	return args.Error(0)
}

func (m *MockTransactionRepo) GetAllTransaction() ([]entity.Transaction, error) {
	args := m.Called()
	return args.Get(0).([]entity.Transaction), args.Error(1)
}

func (m *MockTransactionRepo) UpdateTransaction(transaction entity.Transaction) (entity.Transaction, error) {
	args := m.Called(transaction)
	return args.Get(0).(entity.Transaction), args.Error(1)
}

func (m *MockTransactionRepo) CreateTransaction(transaction entity.Transaction) (entity.Transaction, error) {
	args := m.Called(transaction)
	return args.Get(0).(entity.Transaction), args.Error(1)
}

func (m *MockTransactionRepo) GetTransactionByID(transactions int) (entity.Transaction, error) {
	args := m.Called(transactions)
	return args.Get(0).(entity.Transaction), args.Error(1)
}

func TestCreateTransaction(t *testing.T) {
	mockRepo := new(MockTransactionRepo)
	mockAccountRepo := new(MockAccountRepo)
	mockOperationTypesRepo := new(MockOperationTypesRepo)

	service := service.NewTransactionService(mockRepo, mockAccountRepo, mockOperationTypesRepo)

	eventDateStr := "2020-01-05T09:34:18.5893223"
	eventDate, err := time.Parse("2006-01-02T15:04:05.9999999", eventDateStr)
	if err != nil {
		t.Fatalf("Erro ao converter string para time: %v", err)
	}
	customEventDate := entity.CustomTime{Time: eventDate}
	transaction := entity.Transaction{TransactionId: 10007, AccountId: 1, OperationTypeId: 1, Amount: 10.0, EventDate: customEventDate}
	mockAccountRepo.On("GetAccountByID", 1).Return(entity.Account{AccountId: 1, DocumentNumber: "12345678900"}, nil)
	mockOperationTypesRepo.On("GetOperationTypesByID", 1).Return(entity.OperationsType{OperationTypeId: 1, Description0: "Credit Voucher"}, nil)

	mockRepo.On("CreateTransaction", mock.Anything).Return(transaction, nil)

	createdTransaction, err := service.CreateTransaction(transaction)

	assert.NoError(t, err)
	assert.Equal(t, 10007, createdTransaction.TransactionId)
	mockRepo.AssertExpectations(t)
}

func TestGetTransactionByID(t *testing.T) {
	mockRepo := new(MockTransactionRepo)
	mockAccountRepo := new(MockAccountRepo)
	mockOperationTypesRepo := new(MockOperationTypesRepo)

	service := service.NewTransactionService(mockRepo, mockAccountRepo, mockOperationTypesRepo)

	eventDateStr := "2020-01-05T09:34:18.5893223"
	eventDate, err := time.Parse("2006-01-02T15:04:05.9999999", eventDateStr)
	if err != nil {
		t.Fatalf("Erro ao converter string para time: %v", err)
	}
	customEventDate := entity.CustomTime{Time: eventDate}
	transaction := entity.Transaction{TransactionId: 10008, AccountId: 1, OperationTypeId: 1, Amount: 10.0, EventDate: customEventDate}
	expected := time.Date(2020, time.January, 5, 9, 34, 18, 589322300, time.UTC).Truncate(time.Microsecond)
	actual := transaction.EventDate.Time.Truncate(time.Microsecond)
	mockRepo.On("GetTransactionByID", 1).Return(transaction, nil)

	foundTransaction, err := service.GetTransactionByID(1)

	assert.NoError(t, err)
	assert.Equal(t, 10008, foundTransaction.TransactionId)
	assert.Equal(t, 1, foundTransaction.AccountId)
	assert.Equal(t, 1, foundTransaction.OperationTypeId)
	assert.Equal(t, 10.0, foundTransaction.Amount)
	assert.Equal(t, expected, actual)
	mockRepo.AssertExpectations(t)
}

func TestGetAllTransaction(t *testing.T) {
	mockRepo := new(MockTransactionRepo)
	mockAccountRepo := new(MockAccountRepo)
	mockOperationTypesRepo := new(MockOperationTypesRepo)

	service := service.NewTransactionService(mockRepo, mockAccountRepo, mockOperationTypesRepo)

	eventDateStr := "2020-01-05T09:34:18.5893223"
	eventDate, err := time.Parse("2006-01-02T15:04:05.9999999", eventDateStr)
	if err != nil {
		t.Fatalf("Erro ao converter string para time: %v", err)
	}
	customEventDate := entity.CustomTime{Time: eventDate}
	transaction := []entity.Transaction{
		{TransactionId: 10009, AccountId: 1, OperationTypeId: 1, Amount: 10.0, EventDate: customEventDate},
		{TransactionId: 10010, AccountId: 1, OperationTypeId: 1, Amount: 10.0, EventDate: customEventDate},
	}

	mockRepo.On("GetAllTransaction").Return(transaction, nil)

	foundTransactions, err := service.GetAllTransaction()

	assert.NoError(t, err)
	assert.Len(t, foundTransactions, 2)
	assert.Equal(t, 10009, foundTransactions[0].TransactionId)
	assert.Equal(t, 1, foundTransactions[0].AccountId)
	assert.Equal(t, 10010, foundTransactions[1].TransactionId)
	assert.Equal(t, 1, foundTransactions[1].AccountId)
	mockRepo.AssertExpectations(t)
}

func TestUpdateTransaction(t *testing.T) {
	mockRepo := new(MockTransactionRepo)
	mockAccountRepo := new(MockAccountRepo)
	mockOperationTypesRepo := new(MockOperationTypesRepo)

	service := service.NewTransactionService(mockRepo, mockAccountRepo, mockOperationTypesRepo)

	eventDateStr := "2020-01-05T09:34:18.5893223"
	eventDate, err := time.Parse("2006-01-02T15:04:05.9999999", eventDateStr)
	if err != nil {
		t.Fatalf("Erro ao converter string para time: %v", err)
	}
	customEventDate := entity.CustomTime{Time: eventDate}

	transInput := entity.Transaction{TransactionId: 10011, AccountId: 1, OperationTypeId: 1, Amount: 10.0, EventDate: customEventDate}
	updatedTrans := entity.Transaction{TransactionId: 10012, AccountId: 1, OperationTypeId: 1, Amount: 10.0, EventDate: customEventDate}

	mockRepo.On("UpdateTransaction", transInput).Return(updatedTrans, nil)

	resultTrans, err := service.UpdateTransaction(transInput)

	assert.NoError(t, err)
	assert.Equal(t, 10012, resultTrans.TransactionId)
	assert.Equal(t, 1, resultTrans.AccountId)
	mockRepo.AssertExpectations(t)
}

func TestDeleteTransaction(t *testing.T) {
	mockRepo := new(MockTransactionRepo)
	mockAccountRepo := new(MockAccountRepo)
	mockOperationTypesRepo := new(MockOperationTypesRepo)

	service := service.NewTransactionService(mockRepo, mockAccountRepo, mockOperationTypesRepo)

	transID := 10012

	mockRepo.On("DeleteTransaction", transID).Return(nil)

	err := service.DeleteTransaction(transID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
