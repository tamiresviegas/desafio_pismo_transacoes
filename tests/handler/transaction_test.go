package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/stretchr/testify/assert"
	"github.com/tamiresviegas/desafio_pismo_transacoes/config"
	"github.com/tamiresviegas/desafio_pismo_transacoes/internal/adapters/handler"
	"github.com/tamiresviegas/desafio_pismo_transacoes/internal/adapters/repository"
	"github.com/tamiresviegas/desafio_pismo_transacoes/internal/core/entity"
	"github.com/tamiresviegas/desafio_pismo_transacoes/internal/core/service"
)

func TestCreateTransactionHandler(t *testing.T) {
	db, err := config.ConnectBD()
	if err != nil {
		return
	}
	repo := repository.NewTransactionRepository(db.DB)
	accountrepo := repository.NewAccountRepository(db.DB)
	oprepo := repository.NewOperationTypesRepository(db.DB)
	accountService := service.NewTransactionService(repo, accountrepo, oprepo)
	handler := handler.NewTransactionHandler(accountService)

	eventDateStr := "2020-01-05T09:34:18.5893223"
	eventDate, err := time.Parse("2006-01-02T15:04:05.9999999", eventDateStr)
	if err != nil {
		t.Fatalf("Erro ao converter string para time: %v", err)
	}
	customEventDate := entity.CustomTime{Time: eventDate}
	transaction := entity.Transaction{TransactionId: 10014, AccountId: 1, OperationTypeId: 1, Amount: 10.0, EventDate: customEventDate}

	body, _ := json.Marshal(transaction)
	req, _ := http.NewRequest("POST", "/transactions", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r := gin.Default()
	r.POST("/transactions", handler.CreateTransaction)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var resp entity.Transaction
	json.Unmarshal(w.Body.Bytes(), &resp)
	assert.Equal(t, 10014, resp.TransactionId)
	assert.NotZero(t, resp.AccountId)
}

func TestGetTransactionByIDHandler(t *testing.T) {
	db, err := config.ConnectBD()
	if err != nil {
		return
	}
	repo := repository.NewTransactionRepository(db.DB)
	accountrepo := repository.NewAccountRepository(db.DB)
	oprepo := repository.NewOperationTypesRepository(db.DB)
	accountService := service.NewTransactionService(repo, accountrepo, oprepo)
	handler := handler.NewTransactionHandler(accountService)

	eventDateStr := "2020-01-05T09:34:18.5893223"
	eventDate, err := time.Parse("2006-01-02T15:04:05.9999999", eventDateStr)
	if err != nil {
		t.Fatalf("Erro ao converter string para time: %v", err)
	}
	customEventDate := entity.CustomTime{Time: eventDate}
	transaction := entity.Transaction{TransactionId: 10015, AccountId: 1, OperationTypeId: 1, Amount: 10.0, EventDate: customEventDate}
	db.DB.Create(&transaction)

	req, _ := http.NewRequest("GET", "/transaction/10015", nil)
	w := httptest.NewRecorder()

	r := gin.Default()
	r.GET("/transaction/:transactions", handler.GetTransactionByID)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp entity.Transaction
	json.Unmarshal(w.Body.Bytes(), &resp)
	assert.Equal(t, 10015, resp.TransactionId)
}

func TestGetAllTransactionHandler(t *testing.T) {
	db, err := config.ConnectBD()
	if err != nil {
		return
	}
	repo := repository.NewTransactionRepository(db.DB)
	accountrepo := repository.NewAccountRepository(db.DB)
	oprepo := repository.NewOperationTypesRepository(db.DB)
	accountService := service.NewTransactionService(repo, accountrepo, oprepo)
	handler := handler.NewTransactionHandler(accountService)

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
	db.DB.Create(&transaction)

	req, _ := http.NewRequest("GET", "/transactions", nil)
	w := httptest.NewRecorder()

	r := gin.Default()
	r.GET("/transactions", handler.GetAllTransaction)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp []entity.Transaction
	json.Unmarshal(w.Body.Bytes(), &resp)
}

func TestUpdateTransactionHandler(t *testing.T) {
	db, err := config.ConnectBD()
	if err != nil {
		return
	}
	repo := repository.NewTransactionRepository(db.DB)
	accountrepo := repository.NewAccountRepository(db.DB)
	oprepo := repository.NewOperationTypesRepository(db.DB)
	accountService := service.NewTransactionService(repo, accountrepo, oprepo)
	handler := handler.NewTransactionHandler(accountService)

	eventDateStr := "2020-01-05T09:34:18.5893223"
	eventDate, err := time.Parse("2006-01-02T15:04:05.9999999", eventDateStr)
	if err != nil {
		t.Fatalf("Erro ao converter string para time: %v", err)
	}
	customEventDate := entity.CustomTime{Time: eventDate}
	transaction := entity.Transaction{TransactionId: 10016, AccountId: 1, OperationTypeId: 1, Amount: 10.0, EventDate: customEventDate}
	db.DB.Create(&transaction)

	updatedTransaction := entity.Transaction{TransactionId: 10017, AccountId: 1, OperationTypeId: 1, Amount: 10.0, EventDate: customEventDate}
	body, _ := json.Marshal(updatedTransaction)

	req, _ := http.NewRequest("PATCH", "/transaction/"+strconv.Itoa(transaction.TransactionId), bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	r := gin.Default()
	r.PATCH("/transaction/:transactions", handler.UpdateTransaction)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var resp entity.Transaction
	json.Unmarshal(w.Body.Bytes(), &resp)
	assert.Equal(t, updatedTransaction.TransactionId, resp.TransactionId)
}

func TestDeleteTransactionHandler(t *testing.T) {
	db, err := config.ConnectBD()
	if err != nil {
		return
	}
	repo := repository.NewTransactionRepository(db.DB)
	accountrepo := repository.NewAccountRepository(db.DB)
	oprepo := repository.NewOperationTypesRepository(db.DB)
	accountService := service.NewTransactionService(repo, accountrepo, oprepo)
	handler := handler.NewTransactionHandler(accountService)

	eventDateStr := "2020-01-05T09:34:18.5893223"
	eventDate, err := time.Parse("2006-01-02T15:04:05.9999999", eventDateStr)
	if err != nil {
		t.Fatalf("Erro ao converter string para time: %v", err)
	}
	customEventDate := entity.CustomTime{Time: eventDate}

	transaction := entity.Transaction{TransactionId: 10016, AccountId: 1, OperationTypeId: 1, Amount: 10.0, EventDate: customEventDate}
	db.DB.Create(&transaction)

	req, _ := http.NewRequest("DELETE", "/transaction/"+strconv.Itoa(transaction.TransactionId), nil)
	w := httptest.NewRecorder()

	r := gin.Default()
	r.DELETE("/transaction/:transactions", handler.DeleteTransaction)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp string
	json.Unmarshal(w.Body.Bytes(), &resp)
	assert.Equal(t, "transaction was deleted", resp)
}
