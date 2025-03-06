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

func TestCreateOperationTypesHandler(t *testing.T) {
	db, err := config.ConnectBD()
	if err != nil {
		return
	}
	repo := repository.NewOperationTypesRepository(db.DB)
	optService := service.NewOperationTypesService(repo)
	handler := handler.NewOperationTypesHandler(optService)

	opType := entity.OperationsType{OperationTypeId: 100002, Description0: "Normal Purchase test"}

	body, _ := json.Marshal(opType)
	req, _ := http.NewRequest("POST", "/operationtypes", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r := gin.Default()
	r.POST("/operationtypes", handler.CreateOperationTypes)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var resp entity.OperationsType
	json.Unmarshal(w.Body.Bytes(), &resp)
	assert.Equal(t, "Normal Purchase test", resp.Description0)
	assert.NotZero(t, resp.OperationTypeId)
}

func TestGetOperationTypeByIDHandler(t *testing.T) {
	db, err := config.ConnectBD()
	if err != nil {
		return
	}
	repo := repository.NewOperationTypesRepository(db.DB)
	optService := service.NewOperationTypesService(repo)
	handler := handler.NewOperationTypesHandler(optService)

	opType := entity.OperationsType{OperationTypeId: 100002, Description0: "Normal Purchase test"}
	db.DB.Create(&opType)

	req, _ := http.NewRequest("GET", "/operationtypes/100002", nil)
	w := httptest.NewRecorder()

	r := gin.Default()
	r.GET("/operationtypes/:operationtypeId", handler.GetOperationTypesByID)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp entity.OperationsType
	json.Unmarshal(w.Body.Bytes(), &resp)
	assert.Equal(t, 100002, resp.OperationTypeId)
	assert.Equal(t, "Normal Purchase test", resp.Description0)
}

func TestGetAllOperationTypeHandler(t *testing.T) {
	db, err := config.ConnectBD()
	if err != nil {
		return
	}
	repo := repository.NewOperationTypesRepository(db.DB)
	optService := service.NewOperationTypesService(repo)
	handler := handler.NewOperationTypesHandler(optService)

	opts := []entity.OperationsType{
		{OperationTypeId: 10002, Description0: "Normal Purchase test"},
		{OperationTypeId: 20002, Description0: "Normal Purchase test 2"},
	}
	db.DB.Create(&opts)

	req, _ := http.NewRequest("GET", "/operationtypes", nil)
	w := httptest.NewRecorder()

	r := gin.Default()
	r.GET("/operationtypes", handler.GetAllOperationTypes)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp []entity.OperationsType
	json.Unmarshal(w.Body.Bytes(), &resp)
}

func TestUpdateOperationTypeHandler(t *testing.T) {
	db, err := config.ConnectBD()
	if err != nil {
		return
	}
	repo := repository.NewOperationTypesRepository(db.DB)
	optService := service.NewOperationTypesService(repo)
	handler := handler.NewOperationTypesHandler(optService)

	opType := entity.OperationsType{OperationTypeId: 100002, Description0: "Normal Purchase test"}
	db.DB.Create(&opType)

	updatedOpt := entity.OperationsType{OperationTypeId: 10003, Description0: "Normal Purchase test 2"}
	body, _ := json.Marshal(updatedOpt)

	req, _ := http.NewRequest("PATCH", "/operationtypes/"+strconv.Itoa(opType.OperationTypeId), bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	r := gin.Default()
	r.PATCH("/operationtypes/:operationtypeId", handler.UpdateOperationTypes)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var resp entity.OperationsType
	json.Unmarshal(w.Body.Bytes(), &resp)
	assert.Equal(t, updatedOpt.Description0, resp.Description0)
}

func TestDeleteOperationTypeHandler(t *testing.T) {
	db, err := config.ConnectBD()
	if err != nil {
		return
	}
	repo := repository.NewOperationTypesRepository(db.DB)
	optService := service.NewOperationTypesService(repo)
	handler := handler.NewOperationTypesHandler(optService)

	opType := entity.OperationsType{OperationTypeId: 100002, Description0: "Normal Purchase test"}
	db.DB.Create(&opType)

	req, _ := http.NewRequest("DELETE", "/operationtypes/"+strconv.Itoa(opType.OperationTypeId), nil)
	w := httptest.NewRecorder()

	r := gin.Default()
	r.DELETE("/operationtypes/:operationtypeId", handler.DeleteOperationTypes)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp string
	json.Unmarshal(w.Body.Bytes(), &resp)
	assert.Equal(t, "opTypes was deleted", resp)
}
