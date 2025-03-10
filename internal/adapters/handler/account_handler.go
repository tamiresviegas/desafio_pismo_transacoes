package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/swaggo/swag/example/celler/httputil"
	"github.com/tamiresviegas/desafio_pismo_transacoes/internal/core/entity"
	"github.com/tamiresviegas/desafio_pismo_transacoes/internal/core/service"
)

type AccountHandler struct {
	service *service.AccountService
}

func NewAccountHandler(service *service.AccountService) *AccountHandler {
	return &AccountHandler{service: service}
}

// Ping godoc
// @Summary      Health check
// @Description  Retorna um status 200 se a API estiver online
// @Tags         health
// @Success      200  {string}  string  "pong"
// @Router       /ping [get]
func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

// CreateAccount godoc
// @Summary      Cria uma conta
// @Description  Cria uma nova conta com um número de documento único
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        account  body  entity.Account  true  "Informações da Conta"
// @Success      200  {object}  entity.Account
// @Failure      400  {object}  httputil.HTTPError
// @Router       /accounts [post]
func (h *AccountHandler) CreateAccount(c *gin.Context) {

	var account entity.Account

	if err := c.ShouldBindJSON(&account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	if err := entity.ValidAccount(&account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	accountCreated, err := h.service.CreateAccount(account)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, accountCreated)
}

func (h *AccountHandler) GetAccountByID(c *gin.Context) {
	accountId, err := strconv.Atoi(c.Param("accountId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid account ID"})
		return
	}

	account, err := h.service.GetAccountByID(accountId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "account not found"})
		return
	}

	c.JSON(http.StatusOK, account)
}

func (h *AccountHandler) GetAllAccount(c *gin.Context) {

	account, err := h.service.GetAllAccount()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "account not found"})
		return
	}

	c.JSON(http.StatusOK, account)
}

func (h *AccountHandler) UpdateAccount(c *gin.Context) {

	accountId, err := strconv.Atoi(c.Param("accountId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid account ID"})
		return
	}

	account, err := h.service.GetAccountByID(accountId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "account not found"})
		return
	}
	if err := c.ShouldBindJSON(&account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	if accountId != account.AccountId {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You Can not update the accountID"})
		return
	}

	accountupdated, err := h.service.UpdateAccount(account)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, accountupdated)
}

func (h *AccountHandler) DeleteAccount(c *gin.Context) {
	accountId, err := strconv.Atoi(c.Param("accountId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid account ID"})
		return
	}

	err = h.service.DeleteAccount(accountId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "account not found"})
		return
	}

	c.JSON(http.StatusOK, "Account was deleted")
}
