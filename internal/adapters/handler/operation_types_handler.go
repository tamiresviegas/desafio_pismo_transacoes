package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tamiresviegas/desafio_pismo_transacoes/internal/core/entity"
	"github.com/tamiresviegas/desafio_pismo_transacoes/internal/core/service"
)

type OperationTypesHandler struct {
	service *service.OperationTypesService
}

func NewOperationTypesHandler(service *service.OperationTypesService) *OperationTypesHandler {
	return &OperationTypesHandler{service: service}
}

func (h *OperationTypesHandler) CreateOperationTypes(c *gin.Context) {

	var opTypes entity.OperationsType

	if err := c.ShouldBindJSON(&opTypes); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	opTypesCreated, err := h.service.CreateOperationTypes(opTypes)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, opTypesCreated)
}

func (h *OperationTypesHandler) GetOperationTypesByID(c *gin.Context) {
	opTypesId, err := strconv.Atoi(c.Param("operationtypeId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid opTypes ID"})
		return
	}

	opTypes, err := h.service.GetOperationTypesByID(opTypesId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "opTypes not found"})
		return
	}

	c.JSON(http.StatusOK, opTypes)
}

func (h *OperationTypesHandler) GetAllOperationTypes(c *gin.Context) {

	opTypes, err := h.service.GetAllOperationTypes()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "opTypes not found"})
		return
	}

	c.JSON(http.StatusOK, opTypes)
}

func (h *OperationTypesHandler) UpdateOperationTypes(c *gin.Context) {

	opTypesId, err := strconv.Atoi(c.Param("operationtypeId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid opTypes ID"})
		return
	}

	opTypes, err := h.service.GetOperationTypesByID(opTypesId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "opTypes not found"})
		return
	}
	if err := c.ShouldBindJSON(&opTypes); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	opTypesupdated, err := h.service.UpdateOperationTypes(opTypes)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, opTypesupdated)
}

func (h *OperationTypesHandler) DeleteOperationTypes(c *gin.Context) {
	opTypesId, err := strconv.Atoi(c.Param("operationtypeId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid opTypes ID"})
		return
	}

	err = h.service.DeleteOperationTypes(opTypesId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "opTypes not found"})
		return
	}

	c.JSON(http.StatusOK, "opTypes was deleted")
}
