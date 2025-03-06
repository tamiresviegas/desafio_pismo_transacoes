package http

import (
	"github.com/gin-gonic/gin"
	"github.com/tamiresviegas/desafio_pismo_transacoes/internal/adapters/handler"
)

func SetupRoutes(accountHandler *handler.AccountHandler, opTypesRepoHandler *handler.OperationTypesHandler, transactionHandler *handler.TransactionHandler) *gin.Engine {
	r := gin.Default()

	//Account
	r.POST("/accounts", accountHandler.CreateAccount)
	r.GET("/accounts", accountHandler.GetAllAccount)
	r.GET("/accounts/:accountId", accountHandler.GetAccountByID)
	r.DELETE("/account/:accountId", accountHandler.DeleteAccount)
	r.PATCH("/account/:accountId", accountHandler.UpdateAccount)

	//OperationType
	r.GET("/operationtypes", opTypesRepoHandler.GetAllOperationTypes)
	r.GET("/operationtypes/:operationtypeId", opTypesRepoHandler.GetOperationTypesByID)
	r.POST("/newoperationtypes", opTypesRepoHandler.CreateOperationTypes)
	r.DELETE("/operationtypes/:operationtypeId", opTypesRepoHandler.DeleteOperationTypes)
	r.PATCH("/operationtypes/:operationtypeId", opTypesRepoHandler.UpdateOperationTypes)

	//Transaction
	r.GET("/transactions", transactionHandler.GetAllTransaction)
	r.GET("/transaction/:transactions", transactionHandler.GetTransactionByID)
	r.POST("/transactions", transactionHandler.CreateTransaction)
	r.DELETE("/transaction/:transactions", transactionHandler.DeleteTransaction)
	r.PATCH("/transaction/:transactions", transactionHandler.UpdateTransaction)

	return r
}
