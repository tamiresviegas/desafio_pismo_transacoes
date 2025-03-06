package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"

	"github.com/stretchr/testify/assert"
	"github.com/tamiresviegas/desafio_pismo_transacoes/config"
	"github.com/tamiresviegas/desafio_pismo_transacoes/internal/adapters/handler"
	"github.com/tamiresviegas/desafio_pismo_transacoes/internal/adapters/repository"
	"github.com/tamiresviegas/desafio_pismo_transacoes/internal/core/entity"
	"github.com/tamiresviegas/desafio_pismo_transacoes/internal/core/service"
)

func TestCreateAccountHandler(t *testing.T) {
	db, err := config.ConnectBD()
	if err != nil {
		return
	}
	repo := repository.NewAccountRepository(db.DB)
	accountService := service.NewAccountService(repo)
	handler := handler.NewAccountHandler(accountService)

	account := entity.Account{DocumentNumber: "12345678900"}

	body, _ := json.Marshal(account)
	req, _ := http.NewRequest("POST", "/accounts", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r := gin.Default()
	r.POST("/accounts", handler.CreateAccount)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var resp entity.Account
	json.Unmarshal(w.Body.Bytes(), &resp)
	assert.Equal(t, "12345678900", resp.DocumentNumber)
	assert.NotZero(t, resp.AccountId)
}

func TestGetAccountByIDHandler(t *testing.T) {
	db, err := config.ConnectBD()
	if err != nil {
		return
	}
	repo := repository.NewAccountRepository(db.DB)
	accountService := service.NewAccountService(repo)
	handler := handler.NewAccountHandler(accountService)

	account := entity.Account{AccountId: 1, DocumentNumber: "12345678900"}
	db.DB.Create(&account)

	req, _ := http.NewRequest("GET", "/accounts/1", nil)
	w := httptest.NewRecorder()

	r := gin.Default()
	r.GET("/accounts/:accountId", handler.GetAccountByID)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp entity.Account
	json.Unmarshal(w.Body.Bytes(), &resp)
	assert.Equal(t, 1, resp.AccountId)
	assert.Equal(t, "12345678900", resp.DocumentNumber)
}

func TestGetAllAccountHandler(t *testing.T) {
	db, err := config.ConnectBD()
	if err != nil {
		return
	}
	repo := repository.NewAccountRepository(db.DB)
	accountService := service.NewAccountService(repo)
	handler := handler.NewAccountHandler(accountService)

	// Criar contas para teste
	accounts := []entity.Account{
		{DocumentNumber: "12345678900"},
		{DocumentNumber: "98765432100"},
	}
	db.DB.Create(&accounts)

	req, _ := http.NewRequest("GET", "/accounts", nil)
	w := httptest.NewRecorder()

	r := gin.Default()
	r.GET("/accounts", handler.GetAllAccount)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp []entity.Account
	json.Unmarshal(w.Body.Bytes(), &resp)
}

func TestUpdateAccountHandler(t *testing.T) {
	db, err := config.ConnectBD()
	if err != nil {
		return
	}
	repo := repository.NewAccountRepository(db.DB)
	accountService := service.NewAccountService(repo)
	handler := handler.NewAccountHandler(accountService)

	// Criar uma conta para teste
	account := entity.Account{DocumentNumber: "12345678900"}
	db.DB.Create(&account)

	// Modificar a conta
	updatedAccount := entity.Account{AccountId: account.AccountId, DocumentNumber: "11122233344"}
	body, _ := json.Marshal(updatedAccount)

	req, _ := http.NewRequest("PATCH", "/accounts/"+strconv.Itoa(account.AccountId), bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	r := gin.Default()
	r.PATCH("/accounts/:accountId", handler.UpdateAccount)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var resp entity.Account
	json.Unmarshal(w.Body.Bytes(), &resp)
	assert.Equal(t, updatedAccount.DocumentNumber, resp.DocumentNumber)
}

func TestDeleteAccountHandler(t *testing.T) {
	db, err := config.ConnectBD()
	if err != nil {
		return
	}
	repo := repository.NewAccountRepository(db.DB)
	accountService := service.NewAccountService(repo)
	handler := handler.NewAccountHandler(accountService)

	// Criar uma conta para deletar
	account := entity.Account{DocumentNumber: "12345678900"}
	db.DB.Create(&account)

	req, _ := http.NewRequest("DELETE", "/accounts/"+strconv.Itoa(account.AccountId), nil)
	w := httptest.NewRecorder()

	r := gin.Default()
	r.DELETE("/accounts/:accountId", handler.DeleteAccount)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp string
	json.Unmarshal(w.Body.Bytes(), &resp)
	assert.Equal(t, "Account was deleted", resp)
}
